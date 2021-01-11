package authorfx

import (
	authorv1 "github.com/haunt98/bloguru/gen/author/v1"
	"github.com/haunt98/bloguru/internal/author"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	provideAuthorServer,
)

func provideAuthorServer() authorv1.ServiceServer {
	return author.NewServer()
}
