package product

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBProduct() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_product"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Product{})
	return db, nil
}
