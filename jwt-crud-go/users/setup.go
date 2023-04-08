package users

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBUsers() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_users"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})
	return db, nil
}