package models

type User struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
	Email   string  `gorm:"unique" json:"email"`
	Profile Profile `gorm:"foreignKey:UserId" json:"profile"`
}

type Profile struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Bio    string `json:"bio"`
	UserId uint   `json:"user_id"`
}
