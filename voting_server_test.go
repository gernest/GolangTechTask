package GolangTechTask

import (
	"context"
	"fmt"
	"testing"

	"github.com/buffup/GolangTechTask/api"
	"github.com/google/uuid"
)

func TestServer_memStore(t *testing.T) {
	VotingServerTest(t, NewMemStore())
}

func VotingServerTest(t *testing.T, store Store) {
	t.Run("CreateVoteable", func(t *testing.T) {
		if err := store.Clear(); err != nil {
			t.Fatal(err)
		}
		svr := &Server{store: store}
		vo := &api.CreateVoteableRequest{
			Question: "To be or not to be?",
			Answers:  []string{"yes", "no"},
		}
		res, err := svr.CreateVoteable(context.TODO(), vo)
		if err != nil {
			t.Fatal(err)
		}
		t.Run("must create new uuid uuid set", func(t *testing.T) {
			_, err := uuid.Parse(res.Uuid)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("ListVoteables", func(t *testing.T) {
		if err := store.Clear(); err != nil {
			t.Fatal(err)
		}
		svr := &Server{store: store}
		ctx := context.TODO()
		total := 20
		for i := 0; i < total; i++ {
			_, err := svr.CreateVoteable(ctx, &api.CreateVoteableRequest{
				Question: fmt.Sprintf("Truth or Dare? -%d", i),
				Answers:  []string{"Truth", "Dare"},
			})
			if err != nil {
				t.Fatal(err)
			}
		}

		t.Run("list all", func(t *testing.T) {
			res, err := svr.ListVoteables(ctx, &api.ListVoteableRequest{})
			if err != nil {
				t.Fatal(err)
			}
			if len(res.Votables) != total {
				t.Errorf("expected %d got %d", total, len(res.Votables))
			}
		})
		t.Run("list with limit", func(t *testing.T) {
			limit := 5
			res, err := svr.ListVoteables(ctx, &api.ListVoteableRequest{
				Limit: int32(limit),
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(res.Votables) != limit {
				t.Errorf("expected %d got %d", limit, len(res.Votables))
			}
		})
		t.Run("list with pagination", func(t *testing.T) {
			limit := 5
			res, err := svr.ListVoteables(ctx, &api.ListVoteableRequest{
				Limit: int32(limit),
			})
			if err != nil {
				t.Fatal(err)
			}
			if res.LastIndex == 0 {
				t.Fatal("expected last index to be non zero")
			}
			res, err = svr.ListVoteables(ctx, &api.ListVoteableRequest{
				LastIndex: res.LastIndex,
			})
			if err != nil {
				t.Fatal(err)
			}
			if got, expect := len(res.Votables), total-limit; got != expect {
				t.Errorf("expected %d got %d", expect, got)
			}
		})

		t.Run("list with pagination and limit", func(t *testing.T) {
			limit := 5
			res, err := svr.ListVoteables(ctx, &api.ListVoteableRequest{
				Limit: int32(limit),
			})
			if err != nil {
				t.Fatal(err)
			}
			if res.LastIndex == 0 {
				t.Fatal("expected last index to be non zero")
			}
			res, err = svr.ListVoteables(ctx, &api.ListVoteableRequest{
				LastIndex: res.LastIndex,
				Limit:     int32(limit),
			})
			if err != nil {
				t.Fatal(err)
			}
			if got, expect := len(res.Votables), limit; got != expect {
				t.Errorf("expected %d got %d", expect, got)
			}
		})
	})
}
