package firebaseauth

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewClient ... クライアントを作成
func NewClient(credentialsPath string) *auth.Client {
	ctx := context.Background()
	cOpt := option.WithCredentialsFile(credentialsPath)
	gOpt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Millisecond,
		Timeout:             20 * time.Millisecond,
		PermitWithoutStream: true,
	}))
	app, err := firebase.NewApp(ctx, nil, cOpt, gOpt)
	if err != nil {
		panic(err)
	}
	cli, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	return cli
}
