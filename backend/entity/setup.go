package entity

import (
	//"fmt"
	//"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Student{},
		&Teacher{},
		&Teacher_assessment{},
		&Teaching_duration{},
		&Content_difficulty_level{},
		&Content_quality{},
	)

	db = database
	phone_1, err := bcrypt.GenerateFromPassword([]byte("0935463156"), 14)
	phone_2, err := bcrypt.GenerateFromPassword([]byte("0917271607"), 14)

	// student data
	db.Model(&Student{}).Create(&Student{
		Name:         "นายเอกลักษณ์ มาดี",
		S_ID:         "B6300001",
		College_year: 8003,
		Gpx:          3.00,
		//Faculty_ID:    1004,
		Date_of_birth: "20/06/2001",
		Phone:         string(phone_1),
		Parent:        "Banjoong",

		//Teacher_ID:    5001,
		//Officer_ID:    0101,
	})
	db.Model(&Student{}).Create(&Student{
		Name:         "นางสาวลัดดา ตั้งใจ",
		S_ID:         "B6300002",
		College_year: 8003,
		Gpx:          3.00,
		//Faculty_ID:    1004,
		Date_of_birth: "26/07/2001",
		Phone:         string(phone_2),
		Parent:        "Sawat",

		//Teacher_ID:    5001,
		//Officer_ID:    0102,
	})
	// teacher_data data
	db.Model(&Teacher{}).Create(&Teacher{
		//Faculty_id:                1004,
		Level: "Mr.",
		Name:  "สมชาย ทันสมัย",
		Email: "Somchay@gmail.com",
		//Graduate_faculty_level_id: 2002,
		//Officer_id:                0001,
	})
	db.Model(&Teacher{}).Create(&Teacher{
		//Faculty_ID:                1001,
		Level: "Mrs.",
		Name:  "สมหญิง จุลพล",
		Email: "Somying@gmail.com",
		//Graduate_faculty_level_id: 2002,
		//Officer_id:                0001,
	})

	//Teaching_duration data

	duration_one := Teaching_duration{
		Description: "ใช้เวลานานเกินไปในการสอน",
	}
	db.Model(&Teaching_duration{}).Create(&duration_one)

	duration_two := Teaching_duration{
		Description: "มีการใช้เวลาบางเนื้อหานานเกินไป",
	}
	db.Model(&Teaching_duration{}).Create(&duration_two)

	duration_three := Teaching_duration{
		Description: "เวลาเหมาะสมกับเนื้อหาที่เรียน",
	}
	db.Model(&Teaching_duration{}).Create(&duration_three)

	//content_quality data

	quality_one := Content_quality{
		Description: "เนื้อหาควรที่จะมีการอัพเดต",
	}
	db.Model(&Content_quality{}).Create(&quality_one)

	quality_two := Content_quality{
		Description: "เนื้อหาอยู่ในระดับปกติ",
	}
	db.Model(&Content_quality{}).Create(&quality_two)

	quality_three := Content_quality{
		Description: "เนื้อหามีความเหมาะสมกับรายวิชา",
	}
	db.Model(&Content_quality{}).Create(&quality_three)

	//content difficulty level data

	difficulty_one := Content_difficulty_level{
		Description: "ยากเกินไป",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_one)

	difficulty_two := Content_difficulty_level{
		Description: "พอใช้",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_two)

	difficulty_three := Content_difficulty_level{
		Description: "เหมาะสม",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_three)

	difficulty_four := Content_difficulty_level{
		Description: "เข้าใจง่าย",
	}
	db.Model(&Content_difficulty_level{}).Create(&difficulty_four)

	var Khunjira Student
	var Natiluk Student
	db.Raw("SELECT * FROM Students WHERE name = ? ", "Khunjira Pantuket").Scan(&Khunjira)
	db.Raw("SELECT * FROM Students WHERE name = ? ", "Natiluk Srisakkwa").Scan(&Natiluk)

	var Somchay Teacher
	var Somying Teacher
	db.Raw("SELECT * FROM Teachers WHERE Name = ? ", "Somchay Tunsamai").Scan(&Somchay)
	db.Raw("SELECT * FROM Teachers WHERE Name = ? ", "Somying Junpon").Scan(&Somying)

}
