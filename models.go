package user_sdk

type User struct {
	UserID int    `json:"userId" db:"user_id"`
	Name   string `json:"name"   db:"name"`
	CityID int    `json:"cityId" db:"city_id"`
}
type UserWithCity struct {
	UserID    int     `json:"userId" db:"user_id"`
	Name      string  `json:"name" db:"name"`
	CityID    int     `json:"cityId" db:"city_id"`
	CityName  string  `json:"cityName" db:"city_name"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}
