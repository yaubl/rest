package users

import (
	"bls/db"
	"context"
	"time"

	"github.com/elisiei/gache"
)

type Service struct {
	q *db.Queries
	c *gache.Cache[string, db.User]
}

func NewService(q *db.Queries) *Service {
	c := gache.New[string, db.User](time.Hour * 12)
	return &Service{q, c}
}

func (s *Service) GetAll(ctx context.Context, limit, offset int64) ([]db.User, error) {
	return s.q.ListUsers(ctx, db.ListUsersParams{Limit: limit, Offset: offset})
}

func (s *Service) GetOne(ctx context.Context, id string) (db.User, error) {
	if user, cached := s.c.Get(id); cached {
		return user, nil
	} else {
		user, err := s.q.GetUser(ctx, id)
		if err == nil {
			s.c.Set(id, user, time.Hour*2)
		}
		return user, err
	}
}
