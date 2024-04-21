package input

type GetUserDTO struct {
	UserID string `json:"-" param:"userId" validate:"required"`
}

type CreateUserDTO struct {
	LoginID       string `json:"loginID" validate:"required"`
	ShopID        string `json:"shop_id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Password      string `json:"password" validate:"required"`
	IsShopManager bool   `json:"is_shop_manager" validate:"required"`
}
