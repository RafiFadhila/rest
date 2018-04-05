package models

type Officer struct {
	TblOfficerID    int    `json:"TblOfficeID"`
	OfficerCode     string `json:"OfficerCode"`
	OfficerName     string `json:"OfficerName"`
	OfficerPassword string `json:"OfficerPassword"`
	OfficerStatus   string `json:"OfficerStatus"`
}
