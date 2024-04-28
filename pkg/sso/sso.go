package sso

import (
	"context"
	"fmt"
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type Client struct {
	log   *slog.Logger
	gRPC  ssov1.AuthClient
	appID int
}

func (c *Client) NewClient(log *slog.Logger, host string, port, appID int) *Client {
	address := fmt.Sprintf("%s:%v", host, port)
	conn := mustListen(address)
	client := ssov1.NewAuthClient(conn)
	return &Client{
		log:   log,
		gRPC:  client,
		appID: appID,
	}
}

func mustListen(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("SSO gRPC client couldn't listen on " + address)
	}
	return conn
}

func (c *Client) Login(ctx context.Context, email, password string) (string, error) {
	request := ssov1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    int32(c.appID),
	}
	response, err := c.gRPC.Login(ctx, &request)
	if err != nil {
		c.log.With(slog.String("email", email)).Info("error logging user in")
		return "", err
	}
	return response.GetToken(), nil
}

func (c *Client) Register(ctx context.Context, email, password string) (int64, error) {
	request := ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	}
	response, err := c.gRPC.Register(ctx, &request)
	if err != nil {
		c.log.With(slog.String("email", email)).Info("error registering user")
		return 0, err
	}
	return response.GetUserId(), nil
}

func (c *Client) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	request := ssov1.IsAdminRequest{
		UserId: userID,
	}
	response, err := c.gRPC.IsAdmin(ctx, &request)
	if err != nil {
		c.log.With(slog.Int("id", int(userID))).Info("error authorizing user")
		return false, err
	}
	return response.GetIsAdmin(), nil
}
