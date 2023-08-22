package orm

import "github.com/FourWD/middleware/orm"

type Branch struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Label      string  `json:"label" query:"label" gorm:"type:varchar(100)"`
	Name       string  `json:"name" query:"name" gorm:"type:varchar(100)"`
	Address    string  `json:"address" query:"address" gorm:"type:text"`
	ProvinceID string  `json:"province_id" query:"province_id" gorm:"type:varchar(36)"`
	Phone1     string  `json:"phone_1" query:"phone_1" gorm:"type:varchar(20)"`
	Phone2     string  `json:"phone_2" query:"phone_2" gorm:"type:varchar(20)"`
	Line       string  `json:"line" query:"line" gorm:"type:varchar(20)"`
	Facebook   string  `json:"facebook" query:"facebook" gorm:"type:varchar(20)"`
	Tiktok     string  `json:"tiktok" query:"tiktok" gorm:"type:varchar(20)"`
	Latitude   float64 `json:"latitude" query:"latitude" gorm:"type:float"`
	Longitude  float64 `json:"longitude" query:"longitude" gorm:"type:float"`
	orm.RowOrder
}
