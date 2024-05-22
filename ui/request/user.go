package request

import (
	"github.com/OLIENTTECH/backend-challenges/usecase/input"
)

type (
	LoginUserRequest = input.LoginUserDTO
	GetUserRequest  = input.GetUserDTO
	PostUserRequest = input.CreateUserDTO
)
