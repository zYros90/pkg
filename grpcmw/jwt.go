package grpcmw

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type jwtConfig struct {
	Key    []byte
	logger *zap.Logger
}

const BearerKey = "Bearer"

func NewJwtMW(
	key string,
	logger *zap.Logger,
) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	m := &jwtConfig{
		Key:    []byte(key),
		logger: logger,
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, exist := metadata.FromIncomingContext(ctx)
		if !exist {
			return nil, errors.New("no metadata")
		}

		tokenArr := md.Get(BearerKey)
		if len(tokenArr) != 1 {
			return nil, errors.New("no token")
		}

		token := tokenArr[0]

		claims, err := m.validateJWT(token)
		if err != nil {
			return nil, err
		}
		for key, value := range claims {
			md.Set(key, fmt.Sprintf("%v", value))
		}

		ctx = metadata.NewIncomingContext(ctx, md)
		return handler(ctx, req)
	}
}

func (m *jwtConfig) validateJWT(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return m.Key, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid jwt")
	}
	claimMap := map[string]interface{}(claims)
	return claimMap, nil
}
