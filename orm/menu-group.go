package orm

import "github.com/FourWD/middleware/orm"

type MenuGroup struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code string `json:"code" query:"code" db:"code" gorm:"unique;index;type:varchar(45)"`
	Name string `json:"name" query:"name" db:"name" gorm:"not null;type:varchar(50)"`
}
