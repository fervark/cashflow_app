package categoryModel

type Category struct {
	Id   uint   `gorm:"primaryKey;autoIncrement;default:uuid_generate_v3()" json:"id"`
	Name string `gorm:"size:255" json:"name" validate:"required"`
}
