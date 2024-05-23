package output

import "time"

type ListUsers struct {
	Users []*UserDTO `json:"users"`
}

type UserDTO struct {
	User User `json:"user"`
	Shop Shop `json:"shop"`
}

type User struct {
	ID            string     `json:"id"`
	ShopID        string     `json:"shopID"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	IsShopManager bool       `json:"isShopManager"`
	LastLoginedAt *time.Time `json:"lastLoginedAt"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

type Shop struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
