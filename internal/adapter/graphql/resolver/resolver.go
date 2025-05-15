package resolver

import (
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/gateway"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
)

type Resolver struct {
	auth *auth.Auth
}

func NewResolver(auth *auth.Auth) *Resolver {
	return &Resolver{
		auth: auth,
	}
}

func (r *Resolver) Mutation() gateway.MutationResolver { //nolint:ireturn // TODO: fix the linter if possible
	return &mutationResolver{r}
}

func (r *Resolver) Query() gateway.QueryResolver { //nolint:ireturn // TODO: fix the linter if possible
	return &queryResolver{r}
}

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
