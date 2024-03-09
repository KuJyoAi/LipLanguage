package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

const (
	QuestionsTypeElementary = "element"
	QuestionsTypeCharacter  = "char"
	QuestionsTypeWord       = "word"
)

type Question struct {
	gorm.Model
	Type   string          `gorm:"type:varchar(20)" json:"type"`
	Data   json.RawMessage `gorm:"type:json" json:"data"`
	Answer string          `json:"answer"` // 答案
}

func (Question) TableName() string {
	return "questions"
}

type QuestionElement struct {
	Name              string `json:"name"`               // element name
	Desc              string `json:"desc"`               // element desc
	SpellPicPath      string `json:"spell_pic_path"`     // spell pic
	LipPicPath        string `json:"lip_pic_path"`       // lip pic
	LongSentencePath  string `json:"long_sentence_path"` // one long sentence
	ShortSentencePath string `json:"short_sentence"`     // one short sentence
}

type QuestionCharacter struct {
	Name         string `json:"name"`           // character name 你
	Spell        string `json:"spell"`          // character spell nǐ
	VoicePath    string `json:"voice_path"`     // character voice
	Desc         string `json:"desc"`           // character desc
	SpellPicPath string `json:"spell_pic_path"` // spell pic
	LipPicPath   string `json:"lip_pic_path"`   // lip pic
	VideoPath    string `json:"video_path"`     // video
}

type QuestionWord struct {
	Name      string `json:"name"`       // word name
	Spell     string `json:"spell"`      // word spell
	VoicePath string `json:"voice_path"` // word voice
	VideoPath string `json:"video_path"` // word video
}

type UserLearnRecord struct {
	gorm.Model
	UserID     uint   `gorm:"user_id" json:"user_id"`
	QuestionID uint   `gorm:"question_id" json:"question_id"`
	Result     string `gorm:"result" json:"result"`
	Right      bool   `gorm:"right" json:"right"`
}
