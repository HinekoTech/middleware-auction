package orm

import (
	"time"
)

type Maintenance struct {
	ID                  string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	MaintenanceTypeName string    `json:"maintenance_type_name" query:"maintenance_type_name" gorm:"type:varchar(200)"`
	LocationName        string    `json:"location_name" query:"location_name" gorm:"type:varchar(500)"`
	ChassisNumber       string    `json:"chassis_number" query:"chassis_number" gorm:"type:varchar(20)"`
	Mile                int       `json:"mile" query:"mile" gorm:"type:int"`
	MaintenanceDate     time.Time `json:"maintenance_date" query:"maintenance_date"`
	Detail              string    `json:"detail" query:"detail" gorm:"type:varchar(500)"`
}
