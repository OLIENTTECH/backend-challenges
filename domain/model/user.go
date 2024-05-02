package model

import (
	"time"

	"github.com/OLIENTTECH/backend-challenges/pkg/ctime"
	"github.com/OLIENTTECH/backend-challenges/pkg/ulid"
	"github.com/OLIENTTECH/backend-challenges/usecase/output"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`
	ID            string       `bun:"id,pk"`
	ShopID        string       `bun:"shop_id,notnull"`
	Shop          *Shop        `bun:"rel:belongs-to,join:shop_id=id"`
	Name          string       `bun:"name,notnull"`
	Email         string       `bun:"email,notnull"`
	Password      string       `bun:"password,notnull"`
	IsShopManager bool         `bun:"is_shop_manager,notnull"`
	LastLoginedAt bun.NullTime `bun:"last_logined_at,nullzero"`
	CreatedAt     time.Time    `bun:"created_at,notnull"`
	UpdatedAt     time.Time    `bun:"updated_at,notnull"`
	DeletedAt     bun.NullTime `bun:"deleted_at,nullzero"`
}

type Shop struct {
    bun.BaseModel `bun:"shops,alias:s"`
    ID            string    `bun:"id,pk"`
    Name          string    `bun:"name,notnull"`
    CreatedAt     time.Time `bun:"created_at,notnull"`
    UpdatedAt     time.Time `bun:"updated_at,notnull"`
    DeletedAt     time.Time `bun:"deleted_at,nullzero"`
}

func NewUser(
	loginID string,
	shopID string,
    name string,
	email string,
	password string,
	isShopManager bool,
) *User {
	return &User{
		ID:            ulid.NewULID(),
		ShopID:        shopID,
		Name:          name,
		Email:         email,
		Password:      password,
		IsShopManager: isShopManager,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (u *User) ToDTO() *output.UserDTO {
	if u == nil {
        return nil
    }
    userDTO := &output.UserDTO{
        User: output.User{
			ID:            u.ID,
            ShopID:        u.ShopID,
            Name:          u.Name,
            Email:         u.Email,
            Password:      u.Password,
            LastLoginedAt: ctime.NullTimeToPtrJST(u.LastLoginedAt),
            CreatedAt:     &u.CreatedAt,
            UpdatedAt:     &u.UpdatedAt,
        },
    }
	if u.Shop != nil {
        userDTO.Shop = output.Shop{
            ID:            u.Shop.ID,
            Name:          u.Shop.Name,
            CreatedAt:     &u.Shop.CreatedAt,
            UpdatedAt:     &u.Shop.UpdatedAt,
        }
    } else {
        userDTO.Shop = output.Shop{}
    }
    return userDTO
}
