package output

import "time"

type ListUsers struct {
	Users []*UserDTO `json:"users"`
}

type UserDTO struct {
	ID            string     `json:"id"`
	ShopID        string     `json:"shopID"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	LastLoginedAt *time.Time `json:"lastLoginedAt"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}
