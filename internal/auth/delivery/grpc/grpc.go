package delivery

import (
	context "context"

	app "github.com/Hareshutit/ShopEase/internal/auth/usecase"
)

type GrpcServer struct {
	UnimplementedAuthServer

	command app.Commands
	query   app.Queries
}

func (g *GrpcServer) GenerateToken(ctx context.Context, in *Id) (*UuidAuth, error) {

	resultByte, _, err := g.command.CreateAccessToken.Create(in.Id)
	if err != nil {
		return nil, err
	}
	result := UuidAuth{Value: string(resultByte)}
	return &result, nil
}
