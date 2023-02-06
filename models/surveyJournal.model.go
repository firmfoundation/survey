package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SurveyJournal struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	QuestionID  uuid.UUID `json:"question_id" gorm:"type:varchar(255);not null"`
	Question    Question  `json:"question"`
	SurveyID    uuid.UUID `json:"survey_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Survey      Survey    `json:"survey"`
	AnswerPoint int       `json:"answer_point"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	User        User      `json:"user"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SurveyResult struct {
	SurveyID uuid.UUID `json:"survey_id"`
	UserID   uuid.UUID `json:"user_id"`
	Result   []Result  `json:"result"`
}

type Result struct {
	QuestionID uuid.UUID `json:"question_id"`
	Answer     int       `json:"answer"`
}

func (q *SurveyJournal) CreateSurveyJournal(db *gorm.DB) (*SurveyJournal, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &SurveyJournal{}, err
	}
	return q, nil
}

func (q *SurveyJournal) CreateSurveyResult(db *gorm.DB, batch []SurveyJournal) (*SurveyJournal, error) {
	var err error
	err = db.Debug().Create(&batch).Error
	if err != nil {
		return &SurveyJournal{}, err
	}
	return q, nil
}
