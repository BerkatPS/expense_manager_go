package model

type Categorys struct {
	CategoryID   int    `json:"category_id" gorm:"primaryKey"`
	CategoryName string `json:"category_name"`
	//	Transactions []Transaction `gorm:"foreignKey:CategoryID;references:CategoryID"` // One-to-Many relationship with Transaction
}
