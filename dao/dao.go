package dao

import (
	"TM-Product/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SelectAllProducts(tx *gorm.DB) (products []model.Product, err error) {
	sql := "select * from product"

	if err = tx.Raw(sql).Scan(&products).Error; err != nil {
		log.Println("err ", err)
	}

	return
}

func InsertProducts(product []model.Product, tx *gorm.DB) (err error) {
	result := tx.Debug().Create(&product)
	if result.Error != nil {
		log.Println("err ", result.Error)
		err = result.Error
	}
	fmt.Println("asd", err)
	return
}
