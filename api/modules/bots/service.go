package bots

import (
	"bls/db"
	"context"
	"time"

	"github.com/elisiei/gache"
)

type Service struct {
	q *db.Queries
	c *gache.Cache[string, db.Bot]
}

func NewService(q *db.Queries) *Service {
	c := gache.New[string, db.Bot](time.Hour * 12)
	return &Service{q, c}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int64) ([]db.Bot, error) {
	return s.q.ListBotsByStatus(ctx, db.ListBotsByStatusParams{Limit: limit, Offset: offset, Status: "approved"})
}

func (s *Service) GetOne(ctx context.Context, id string) (db.Bot, error) {
	if bot, cached := s.c.Get(id); cached {
		return bot, nil
	}

	bot, err := s.q.GetBot(ctx, id)
	if err == nil {
		s.c.Set(id, bot, time.Hour)
	}
	return bot, err
}

func (s *Service) Create(ctx context.Context, dto db.CreateBotParams) (db.Bot, error) {
	return s.q.CreateBot(ctx, dto)
}
