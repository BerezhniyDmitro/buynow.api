package auth

import "context"

type AuthStorage interface {
	Register(ctx context.Context, username, password string) error
}
