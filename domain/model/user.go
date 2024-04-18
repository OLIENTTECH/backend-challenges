package model

import (
	"time"

	"github.com/OLIENTTECH/backend-challenges/pkg/ctime"
	"github.com/OLIENTTECH/backend-challenges/pkg/uuid"
	"github.com/OLIENTTECH/backend-challenges/usecase/output"
	"github.com/uptrace/bun"
)

const (
	// role types
	RoleAdmin = iota + 1
	RoleGeneral
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`
	ID            string       `bun:"id,pk"`
	ShopID        string       `bun:"shop_id,notnull"`
	Name          string       `bun:"name,notnull"`
	Email         string       `bun:"email,notnull"`
	Password      string       `bun:"password,notnull"`
	IsShopManager bool         `bun:"is_shop_manager,notnull"`
	LastLoginedAt bun.NullTime `bun:"last_logined_at,nullzero"`
	CreatedAt     time.Time    `bun:"created_at,notnull"`
	UpdatedAt     time.Time    `bun:"updated_at,notnull"`
	DeletedAt     bun.NullTime `bun:"deleted_at,nullzero"`
}

func NewUser(
	loginID string,
	password string,
	familyName string,
	givenName string,
	isShopManager bool,
) *User {
	return &User{
		ID:            uuid.NewUUID(),
		Password:      password,
		IsShopManager: isShopManager,
	}
}

func (u *User) Role() string {
	switch u.IsShopManager {
	case true:
		return "admin"
	case false:
		return "general"
	default:
		return ""
	}
}

func (u *User) Set (roleType string) {
	switch roleType {
	case "admin":
		u.IsShopManager = true
	case "general":
		u.IsShopManager = false
	default:
		u.IsShopManager = false
	}
}

func (u *User) ToDTO() *output.UserDTO {
	return &output.UserDTO{
		ID:            u.ID,
		ShopID:        u.ShopID,
		Name:          u.Name,
		Email:         u.Email,
		Password:      u.Password,
		Role:          u.Role(),
		LastLoginedAt: ctime.NullTimeToPtrJST(u.LastLoginedAt),
		CreatedAt:     &u.CreatedAt,
		UpdatedAt:     &u.UpdatedAt,
	}
}
