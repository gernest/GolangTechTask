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
	err = talkToService(context.Background(), log, vo)
	if err != nil {
		log.Error("Failed to talkToService", zap.Error(err))
		os.Exit(1)
	}
}

func talkToService(ctx context.Context, log *zap.Logger, x api.VotingServiceClient) error {
	log.Info("CreateVoteable")
	total := 3
	// we create and collect uuid for 3 voteables
	var voteables []string
	for i := 0; i < total; i++ {
		res, err := x.CreateVoteable(ctx, &api.CreateVoteableRequest{
			Question: fmt.Sprintf("%d - Truth or Dare?", i),
			Answers:  []string{"Truth", "Dare"},
		})
		if err != nil {
			return err
		}
		voteables = append(voteables, res.Uuid)
	}

	log.Info("ListVoteables - list all voteables")
	_, err := x.ListVoteables(ctx, &api.ListVoteableRequest{})
	if err != nil {
		return err
	}

	log.Info("ListVoteables - list with limit")
	limit := 2
	res, err := x.ListVoteables(ctx, &api.ListVoteableRequest{Limit: int32(limit)})
	if err != nil {
		return err
	}

	log.Info("ListVoteables - list with pagination")
	_, err = x.ListVoteables(ctx, &api.ListVoteableRequest{LastIndex: res.LastIndex})
	if err != nil {
		return err
	}
	log.Info("CastVote")
	_, err = x.CastVote(ctx, &api.CastVoteRequest{Uuid: voteables[0], AnswerIndex: 0})
	if err != nil {
		return err
	}
	return nil
}
