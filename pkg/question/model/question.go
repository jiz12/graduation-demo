package model

type QuestionSingleRelT struct {
	ID         int    `gorm:"primaryKey;index:id_idx;column:id;type:int;not null"`
	Name       string `gorm:"column:name;type:longtext;not null"`
	Text       string `gorm:"column:text;type:longtext;not null"`
	QuesID     int    `gorm:"column:ques_id;type:int;not null"`
	Type       string `gorm:"column:type;type:varchar(12);not null"`
	TypeName   string `gorm:"column:type_name;type:varchar(12);not null"`
	SortNum    int    `gorm:"column:sort_num;type:int;not null"`
	Resolution string `gorm:"column:resolution;type:longtext"`
	Score      int    `gorm:"column:score;type:int;not null"`
	IsCorrect  bool   `gorm:"column:is_correct;type:tinyint(1)"`
}

// TableName 获取数据库表名
func (m *QuestionSingleRelT) TableName() string {
	return "question_single_rel_t"
}

