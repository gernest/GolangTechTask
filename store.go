package GolangTechTask

import (
	"context"
	"errors"
	"sort"
	"strconv"
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
	List(ctx context.Context, lastResultIndex string, limit int) (result []*api.Voteable, lastIndex string, err error)
	Cast(ctx context.Context, id string, answer int) error
	// Clear is needed for testing purpose only. This deletes everything and resets
	// the storage state
	Clear() error
}

var _ Store = (*MemStore)(nil)

type MemStore struct {
	va  *sync.Map
	idx atomic.Int64
}

func (m *MemStore) Clear() error {
	var keys []interface{}
	m.va.Range(func(key, value interface{}) bool {
		keys = append(keys, key)
		return true
	})
	for _, v := range keys {
		m.va.Delete(v)
	}
	m.idx.Store(0)
	return nil
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

func (m *MemStore) List(ctx context.Context, lastResultIndexStr string, limit int) (result []*api.Voteable, lastIndex string, err error) {
	lastResultIndex, _ := strconv.ParseInt(lastResultIndexStr, 10, 64)
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
			lastIndex = strconv.FormatInt(v.index, 10)
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
	if c.Memory {
		return NewMemStore(), nil
	}
	return NewDynamo(c)
}

func NewMemStore() *MemStore {
	return &MemStore{va: &sync.Map{}}
}
