package data

import (
	"database/sql"
	"technology/day15/officer/models"

	_ "github.com/go-sql-driver/mysql"
)

type OfficerRepository struct {
	DB *sql.DB
}

func GetAll(db *OfficerRepository) []models.Officer {
	result, err := db.DB.Query("Select TblOfficerID, OfficerCode, OfficerName, OfficerPassword, OfficerStatus from tblOfficer")
	if err != nil {
		return nil
	}
	defer result.Close()
	officer := []models.Officer{}
	for result.Next() {
		var o models.Officer
		if err := result.Scan(&o.TblOfficerID, &o.OfficerCode, &o.OfficerName, &o.OfficerPassword, &o.OfficerStatus); err != nil {
			return nil
		}
		officer = append(officer, o)
	}
	return officer
}
