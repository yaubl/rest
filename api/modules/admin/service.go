package admin

import (
	"bls/db"
	"context"
	"database/sql"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q}
}

// i know i could just make an "updatebot" function and go on from there
// but no.

func (s *Service) UpdateBotStatus(ctx context.Context, id, status string) (db.Bot, error) {
	return s.q.UpdateBot(ctx, db.UpdateBotParams{ID: id, Status: sql.NullString{String: status, Valid: true}})
}

func (s *Service) GetBots(ctx context.Context, limit, offset int64) ([]db.Bot, error) {
	return s.q.ListBots(ctx, db.ListBotsParams{Limit: limit, Offset: offset})
}

func (s *Service) GetBotsByStatus(ctx context.Context, limit, offset int64) ([]db.Bot, error) {
	return s.q.ListBotsByStatus(ctx, db.ListBotsByStatusParams{Limit: limit, Offset: offset})
}
