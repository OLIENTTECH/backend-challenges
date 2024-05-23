package model

import (
	"time"

	"github.com/OLIENTTECH/backend-challenges/usecase/output"
	"github.com/uptrace/bun"
)

type Shop struct {
    bun.BaseModel `bun:"shops,alias:s"`
    ID            string    `bun:"id,pk"`
    Name          string    `bun:"name,notnull"`
    CreatedAt     time.Time `bun:"created_at,notnull"`
    UpdatedAt     time.Time `bun:"updated_at,notnull"`
    DeletedAt     time.Time `bun:"deleted_at,nullzero"`
}

func (s *Shop) ToDTO() *output.Shop {
	return &output.Shop{
		ID:				s.ID,
		Name:			s.Name,
		CreatedAt:     &s.CreatedAt,
        UpdatedAt:     &s.UpdatedAt,
	}
}