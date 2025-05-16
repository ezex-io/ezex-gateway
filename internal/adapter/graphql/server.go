package graphql

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	ext "github.com/ezex-io/ezex-gateway/internal/adapter/graphql/extension"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
	"github.com/ezex-io/gopkg/logger"
	mdl "github.com/ezex-io/gopkg/middleware/http-mdl"
	"github.com/vektah/gqlparser/v2/ast"
)

type Server struct {
	server *http.Server
	logger logger.Logger
}

func New(cfg *Config, logger logger.Logger, resolver gateway.ResolverRoot,
	middlewares ...mdl.Middleware,
) *Server {
	mux := http.NewServeMux()

	graphSrv := handler.New(gateway.NewExecutableSchema(gateway.Config{
		Resolvers: resolver,
	}))

	graphSrv.AddTransport(transport.Options{})
	graphSrv.AddTransport(transport.GET{})
	graphSrv.AddTransport(transport.POST{})

	graphSrv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	graphSrv.SetErrorPresenter(ext.FormatGQLError)

	graphSrv.Use(extension.Introspection{})
	graphSrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	graphSrv.Use(ext.LoggingExt(logger))

	if cfg.Playground {
		mux.Handle("/playground", playground.Handler("ezeX playground", "/query"))
	}

	queryPath := "/query"

	if cfg.QueryPath != "" {
		queryPath = cfg.QueryPath
	}

	mux.Handle(queryPath, graphSrv)

	defaultCors := mdl.DefaultCORSConfig()

	if cfg.CORS.AllowedOrigins != nil {
		defaultCors.AllowedOrigins = cfg.CORS.AllowedOrigins
		defaultCors.AllowedHeaders = cfg.CORS.AllowedHeaders
		defaultCors.AllowedMethods = cfg.CORS.AllowedMethods
		defaultCors.AllowCredentials = cfg.CORS.AllowCredentials
	}

	middlewares = append(middlewares, mdl.CORS(defaultCors))

	var finalHandler http.Handler = mux
	if len(middlewares) != 0 {
		finalHandler = mdl.Chain(middlewares...)(mux)
	}

	srv := &http.Server{
		Addr:           cfg.Address,
		Handler:        finalHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		server: srv,
		logger: logger,
	}
}

func (s *Server) Start() {
	go func() {
		s.logger.Info("GraphQL server start listening", "address", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Debug("error on GraphQL server", "error", err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	if s.server != nil {
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		_ = s.server.Shutdown(shutdownCtx)
	}
}
