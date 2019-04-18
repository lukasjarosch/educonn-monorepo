package server

import (
	"errors"

	"../../pers/educonn-monorepo/services/account"
	"context"
)

func (s AccountServer) Get(ctx context.Context, r *account.GetRequest) (*account.GetResponse, error) {
	return nil, errors.New("not yet implemented")
}
