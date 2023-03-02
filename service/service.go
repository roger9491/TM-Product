package service

import (
	"TM-Product/dao"
	"TM-Product/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var (
	OrderRepo OrderRepoInterface = &orderSQL{}
)

type OrderRepoInterface interface {
	Initialize(*gorm.DB)
	GetAllProducts() ([]model.Product, error)
	InsertProductData(model.ProductDate) (error)
}

type orderSQL struct {
	db *gorm.DB
}

func (od *orderSQL) Initialize(db *gorm.DB) {
	od.db = db
}

func (od *orderSQL) GetAllProducts() (products []model.Product, err error) {
	tx := od.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()

	products, err = dao.SelectAllProducts(tx)
	if err != nil {
		panic(err)
	}

	tx.Commit()

	return
}

func (od *orderSQL) InsertProductData(productdata model.ProductDate) (err error) {
	tx := od.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()

	resp, err := http.Get(productdata.URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var products []model.Product
	if err = json.Unmarshal(body, &products); err != nil {
		panic(err)
	}

	fmt.Println(products)
	err = dao.InsertProducts(products, tx)
	if err != nil {
		panic(err)
	}
	tx.Commit()

	return

}
