package model

import "time"

type Users struct {
	UserID    int       `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//Accounts  []Accounts `gorm:"foreignKey:UserID;references:UserID"` // One-to-Many relationship with Account
}
