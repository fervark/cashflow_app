package userModel

type User struct {
	FirstName  string `gorm:"size:255" json:"first_name" validate:"required"`
	LastName   string `gorm:"size:255" json:"last_name" validate:"required"`
	FamilyName string `gorm:"size:255" json:"family_name" validate:"required"`
}
