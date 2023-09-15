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

func (g *GrpcServer) GenerateToken(ctx context.Context, in *Id) (*Token, error) {

	accessToken, _, err := g.command.CreateAccessToken.Create(in.Id)
	if err != nil {
		return nil, err
	}

	ctxn := context.TODO()
	refreshToken, _, err := g.command.CreateRefreshToken.Create(ctxn, in.Id)
	if err != nil {
		return nil, err
	}
	result := Token{Access: string(accessToken), Refresh: string(refreshToken)}
	return &result, nil
}
