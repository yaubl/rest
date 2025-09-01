package users

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

func (s *Service) GetAll(ctx context.Context, limit, offset int64) ([]db.User, error) {
	return s.q.ListUsers(ctx, db.ListUsersParams{Limit: limit, Offset: offset})
}

func (s *Service) GetOne(ctx context.Context, id string) (db.User, error) {
	return s.q.GetUser(ctx, id)
}
