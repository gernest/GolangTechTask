package GolangTechTask

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/buffup/GolangTechTask/api"
)

var _ api.VotingServiceServer = (*Server)(nil)

type Server struct {
	store Store
}

func (s *Server) CreateVoteable(ctx context.Context, a *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	uuid, err := s.store.Create(ctx, a.Question, a.Answers)
	if err != nil {
		return nil, err
	}
	return &api.CreateVoteableResponse{Uuid: uuid}, nil
}

func (s *Server) ListVoteables(ctx context.Context, a *api.ListVoteableRequest) (*api.ListVoteableResponse, error) {
	limit := math.MaxInt32
	if a.Limit != 0 {
		limit = int(a.Limit)
	}
	ls, idx, err := s.store.List(ctx, a.LastIndex, limit)
	if err != nil {
		return nil, err
	}
	return &api.ListVoteableResponse{
		Votables:  ls,
		LastIndex: idx,
	}, nil
}

func (s *Server) CastVote(ctx context.Context, a *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	err := s.store.Cast(ctx, a.Uuid, int(a.AnswerIndex))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return &api.CastVoteResponse{
				Status: fmt.Sprintf("Votable %q was not found", a.Uuid),
			}, err
		}
		if errors.Is(err, ErrAlreadyVoted) {
			return &api.CastVoteResponse{
				Status: "Forbidden tou cant cast vote more than once",
			}, err
		}
		if errors.Is(err, ErrWrongIndex) {
			return &api.CastVoteResponse{
				Status: "Wrong answer index",
			}, err
		}
		return nil, err

	}
	return &api.CastVoteResponse{Status: "ok"}, nil
}
