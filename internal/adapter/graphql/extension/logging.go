package extensions

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ezex-io/gopkg/logger"
)

type LoggingExtension struct {
	logging logger.Logger
}

func LoggingExt(logging logger.Logger) *LoggingExtension {
	return &LoggingExtension{logging}
}

func (*LoggingExtension) ExtensionName() string {
	return "LoggingExtension"
}

func (*LoggingExtension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (l *LoggingExtension) InterceptOperation(ctx context.Context,
	next graphql.OperationHandler,
) graphql.ResponseHandler {
	start := time.Now()
	opCtx := graphql.GetOperationContext(ctx)

	respHandler := next(ctx)

	return func(ctx context.Context) *graphql.Response {
		res := respHandler(ctx)
		duration := time.Since(start)

		l.logging.Debug("[GraphQL] new operation called",
			"operation", opCtx.Operation.Operation,
			"operation", opCtx.OperationName,
			"name", duration,
			"errors", res.Errors,
		)

		return res
	}
}
