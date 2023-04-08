package product

type Product struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"varchar(300)" json:"nama_product"`
	Harga       int `gorm:"int(200)" json:"harga"`
}
