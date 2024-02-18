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
	LoginID       string       `bun:"login_id,notnull"`
	Password      string       `bun:"password,notnull"`
	FamilyName    string       `bun:"family_name,notnull"`
	GivenName     string       `bun:"given_name,notnull"`
	RoleID        int          `bun:"role_id,notnull"`
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
	roleID int,
) *User {
	return &User{
		ID:         uuid.NewUUID(),
		LoginID:    loginID,
		Password:   password,
		FamilyName: familyName,
		GivenName:  givenName,
		RoleID:     roleID,
	}
}

func (u *User) Role() string {
	switch u.RoleID {
	case RoleAdmin:
		return "admin"
	case RoleGeneral:
		return "general"
	default:
		return ""
	}
}

func (u *User) SetRoleID(roleType string) {
	switch roleType {
	case "admin":
		u.RoleID = RoleAdmin
	case "general":
		u.RoleID = RoleGeneral
	default:
		u.RoleID = 0
	}
}

func (u *User) ToDTO() *output.UserDTO {
	return &output.UserDTO{
		ID:            u.ID,
		LoginID:       u.LoginID,
		Password:      u.Password,
		FamilyName:    u.FamilyName,
		GivenName:     u.GivenName,
		Role:          u.Role(),
		LastLoginedAt: ctime.NullTimeToPtrJST(u.LastLoginedAt),
	}
}
