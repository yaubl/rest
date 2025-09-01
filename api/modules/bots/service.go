package bots

import (
	"bls/db"
	"context"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int64) ([]db.Bot, error) {
	return s.q.ListBots(ctx, db.ListBotsParams{Limit: limit, Offset: offset})
}

func (s *Service) GetOne(ctx context.Context, id string) (db.Bot, error) {
	return s.q.GetBot(ctx, id)
}

func (s *Service) Create(ctx context.Context, dto db.CreateBotParams) (db.Bot, error) {
	return s.q.CreateBot(ctx, dto)
}
