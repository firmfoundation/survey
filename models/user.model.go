package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string    `gorm:"size:255;not null;unique" json:"full_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Role      int       `json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (q *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &User{}, err
	}
	return q, nil
}

func Exe(db *gorm.DB, survey_id string) []map[string]interface{} {

	var r []map[string]interface{}
	sql := `select i.name as indicator,sum(i.weight) as total_weight, sum(a.answer_point) as total_indicator
	 from survey_journals as a
	 inner join questions as q on a.question_id=q.id
	 inner join indicators as i on i.id=q.indicator_id
	 where a.survey_id=?
	 group by i.name`
	db.Debug().Raw(sql, survey_id).Limit(100).Find(&r)
	return r
}
