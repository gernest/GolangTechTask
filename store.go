package GolangTechTask

import (
	"context"
	"errors"
	"sort"
	"sync"

	"github.com/buffup/GolangTechTask/api"
	"github.com/google/uuid"
	"go.uber.org/atomic"
)

//go:generate protoc -I api/ --go_out=plugins=grpc:./api api/service.proto

var ErrNotFound = errors.New("404")
var ErrAlreadyVoted = errors.New("AlreadyVoted")
var ErrWrongIndex = errors.New("ErrWrongIndex")

type Store interface {
	Create(ctx context.Context, question string, answers []string) (uuid string, ee error)
	List(ctx context.Context, lastResultIndex int64, limit int) (result []*api.Voteable, lastIndex int64, err error)
	Cast(ctx context.Context, id string, answer int) error
}

var _ Store = (*MemStore)(nil)

type MemStore struct {
	va  *sync.Map
	idx atomic.Int64
}

type memVoteable struct {
	index  int64
	va     *api.Voteable
	answer *int
}

func (m *MemStore) Create(ctx context.Context, question string, answers []string) (string, error) {
	u := uuid.New().String()
	v := &api.Voteable{
		Uuid:     u,
		Question: question,
		Answers:  answers,
	}
	m.va.Store(u, &memVoteable{
		index: m.idx.Inc(),
		va:    v,
	})
	return u, nil
}

func (m *MemStore) List(ctx context.Context, lastResultIndex int64, limit int) (result []*api.Voteable, lastIndex int64, err error) {
	var all []*memVoteable
	m.va.Range(func(key, value interface{}) bool {
		all = append(all, value.(*memVoteable))
		return true
	})
	sort.Slice(all, func(i, j int) bool {
		return all[i].index < all[j].index
	})
	for _, v := range all {
		if v.index > lastResultIndex && len(result) < limit {
			result = append(result, v.va)
			lastIndex = v.index
		}
	}
	return
}

func (m *MemStore) Cast(ctx context.Context, id string, answer int) error {
	if v, ok := m.va.Load(id); ok {
		x := v.(*memVoteable)
		if x.answer != nil {
			return ErrAlreadyVoted
		}
		if answer < 0 || answer > len(x.va.Answers) {
			return ErrWrongIndex
		}
		x.answer = &answer
		return nil
	}
	return ErrNotFound
}

func NewStore(c *Config) (Store, error) {
	return NewMemStore(), nil
}

func NewMemStore() *MemStore {
	return &MemStore{va: &sync.Map{}}
}
