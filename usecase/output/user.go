package output

import "time"

type ListUsers struct {
	Users []*UserDTO `json:"users"`
}

type UserDTO struct {
	ID            string     `json:"id"`
	LoginID       string     `json:"loginID"`
	Password      string     `json:"password"`
	FamilyName    string     `json:"familyName"`
	GivenName     string     `json:"givenName"`
	Role          string     `json:"role"`
	LastLoginedAt *time.Time `json:"lastLoginedAt"`
}
