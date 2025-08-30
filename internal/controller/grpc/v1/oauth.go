package v1

import (
	"context"
	"fmt"

	v1 "github.com/kr0106686/oauth2/v2/docs/proto/v1"
)

func (r *V1) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	u, err := r.t.TokenParser(req.Token)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &v1.GetUserResponse{User: u}, nil
}
