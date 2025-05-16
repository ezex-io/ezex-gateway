package extensions

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// FormatGQLError converts internal app errors into a structured GraphQL error.
// If the error is not of type *appErr.Error, it falls back to the default presenter.
func FormatGQLError(ctx context.Context, err error) *gqlerror.Error {
	return graphql.DefaultErrorPresenter(ctx, err)
}
