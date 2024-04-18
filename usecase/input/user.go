package input

type GetUserDTO struct {
	UserID string `json:"-" param:"userId" validate:"required"`
}

type CreateUserDTO struct {
	LoginID       string `json:"loginID" validate:"required"`
	Password      string `json:"password" validate:"required"`
	FamilyName    string `json:"familyName" validate:"required"`
	GivenName     string `json:"givenName" validate:"required"`
	IsShopManager bool   `json:"is_shop_manager" validate:"required"`
}
