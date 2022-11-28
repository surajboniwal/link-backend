package models

type Phone struct {
	CountryCode string `json:"country_code" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
