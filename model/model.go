package model

type Product struct {
	ID          int64  `gorm:"index" json:"id"`
	Title		string  `gorm:"column:title;" json:"title"`
	Price		float64	`gorm:"column:price;" json:"price"`
	Description string	`gorm:"column:description;" json:"description"`
	Category	string	`gorm:"column:category;" json:"category"`
	Image 		string	`gorm:"column:image;" json:"image"`
	Count       int64  `gorm:"column:count;" json:"count"`
	IsDelete 	int		`gorm:"column:isdelete" json:"isdelete"`
}

func (u *Product) TableName() string {
	return "product"
}


type ProductDate struct {
	URL		string `json:"url"`
}