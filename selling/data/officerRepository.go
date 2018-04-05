package data

import (
	"database/sql"
	"technology/day15/selling/models"

	_ "github.com/go-sql-driver/mysql"
)

type SellingRepository struct {
	DB *sql.DB
}

func GetAll(db *SellingRepository) []models.Selling {
	result, err := db.DB.Query("Select TblSellingID, Invoice, InvoiceDate, Item, Total, Paid, Kembali, OfficerCode from tblSelling")
	if err != nil {
		return nil
	}
	defer result.Close()
	selling := []models.Selling{}
	for result.Next() {
		var s models.Selling
		if err := result.Scan(&s.TblSellingID, &s.Invoice, &s.InvoiceDate, &s.Item, &s.Total, &s.Paid, &s.Kembali, &s.OfficerCode); err != nil {
			return nil
		}
		selling = append(selling, s)
	}
	return selling
}
