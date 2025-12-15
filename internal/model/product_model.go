package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductModel struct {
	ID            uuid.UUID
	SKU           string
	Name          string
	Stock         int
	ProductTypeID *int    // nullable (LEFT JOIN)
	TypeName      *string // PT.NAME
	ImageURL      *string
	CreatedBy     uuid.UUID
	CreatedAt     time.Time
	UpdatedBy     *uuid.UUID
	UpdatedAt     *time.Time
}
