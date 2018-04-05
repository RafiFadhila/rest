package data

import (
	"database/sql"
	"technology/day15/item/models"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepository struct {
	DB *sql.DB
}

func GetAll(db *ItemRepository) []models.Item {
	result, err := db.DB.Query("Select ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount, Pieces from tblItem")
	if err != nil {
		return nil
	}
	defer result.Close()
	item := []models.Item{}
	for result.Next() {
		var i models.Item
		if err := result.Scan(&i.ItemCode, &i.ItemName, &i.BuyingPrice, &i.SellingPrice, &i.ItemAmount, &i.Pieces); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}
