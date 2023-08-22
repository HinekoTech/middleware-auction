package orm

import "github.com/FourWD/middleware/orm"

type Config struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Key      string `json:"key" query:"key" gorm:"type:varchar(50)"`
	KeyValue string `json:"key_value" query:"key_value" gorm:"type:varchar(200)"`
}
