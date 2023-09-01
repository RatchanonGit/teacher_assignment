package entity

import (
	//"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex"`
	S_ID          string `gorm:"uniqueIndex"`
	College_year  uint
	Gpx           uint
	Date_of_birth string
	Phone         string
	Parent        string

	Teacher_ID *uint
	Officer_ID *uint
	Faculty_ID *uint
	Teacher    Teacher `gorm:"references:id"`
	//Officer Officer `gorm:"references:id"`
	//Faculty Faculty `gorm:"references:id"`

	Teacher_assessment []Teacher_assessment `gorm:"foreignKey:Student_ID"`
}

type Teacher struct {
	gorm.Model

	Level string
	Name  string `gorm:"uniqueIndex"`
	Email string

	Graduate_faculty_level_ID *uint
	Officer_ID                *uint
	Faculty_ID                *uint
	//Graduate_faculty_level_ID Graduate_faculty_level_ID `gorm:"references:id"`
	//Officer_ID Officer_ID `gorm:"references:id"`

	Teacher_assessment []Teacher_assessment `gorm:"foreignKey:Teacher_ID"`
}

type Teaching_duration struct {
	gorm.Model
	Description string `gorm:"uniqueIndex"`

	Teacher_assessment []Teacher_assessment `gorm:"foreignKey:Teaching_duration_ID"`
}

type Content_difficulty_level struct {
	gorm.Model
	Description string `gorm:"uniqueIndex"`

	Teacher_assessment []Teacher_assessment `gorm:"foreignKey:Content_difficulty_level_ID"`
}

type Content_quality struct {
	gorm.Model
	Description string `gorm:"uniqueIndex"`

	Teacher_assessment []Teacher_assessment `gorm:"foreignKey:Content_quality_ID"`
}

type Teacher_assessment struct {
	gorm.Model
	Comment string

	Student_ID *uint
	Student    Student `gorm:"references:id"`

	Teacher_ID *uint
	Teacher    Teacher `gorm:"references:id"`

	Teaching_duration_ID *uint
	Teaching_duration    Teaching_duration `gorm:"references:id"`

	Content_difficulty_level_ID *uint
	Content_difficulty_level    Content_difficulty_level `gorm:"references:id"`

	Content_quality_ID *uint
	Content_quality    Content_quality `gorm:"references:id"`
}
