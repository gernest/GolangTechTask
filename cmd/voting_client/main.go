package main

import (
	"context"
	"fmt"
	"os"

	"github.com/buffup/GolangTechTask"
	"github.com/buffup/GolangTechTask/api"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	log := GolangTechTask.Logger.Named("main.Client")
	log.Info("Connecting to the service", zap.Int("port", 80080))
	decider := func(ctx context.Context, fullMethodName string) bool { return true }

	x, err := grpc.Dial(":8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_zap.PayloadUnaryClientInterceptor(log, decider)),
	)
	if err != nil {
		log.Error("Failed to connect", zap.Error(err))
		os.Exit(1)
	}
	vo := api.NewVotingServiceClient(x)
	talkToService(context.Background(), log, vo)
}

func talkToService(ctx context.Context, log *zap.Logger, x api.VotingServiceClient) error {
	log.Info("CreateVoteable")
	total := 10
	for i := 0; i < total; i++ {
		_, err := x.CreateVoteable(ctx, &api.CreateVoteableRequest{
			Question: fmt.Sprintf("%d - Truth or Dare?", i),
			Answers:  []string{"Truth", "Dare"},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
