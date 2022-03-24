package types

type User struct {
	ID   int    `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}
