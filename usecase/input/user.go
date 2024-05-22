package input

type LoginUserDTO struct {
	ShopID   string `json:"shopID" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type GetUserDTO struct {
	UserID string `json:"-" param:"userId" validate:"required"`
}

type CreateUserDTO struct {
	ShopID string `json:"shopID" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
}
