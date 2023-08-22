package orm

import "github.com/FourWD/middleware/orm"

type SurveyAnwser struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SurveyID       string `json:"survey_group_id" query:"survey_group_id" gorm:"type:varchar(36);"`
	SurveyOptionID string `json:"survey_group_id" query:"survey_group_id" gorm:"type:varchar(36);"`
	UserID         string `json:"user_id" query:"user_id" gorm:"type:varchar(36);"`
}
