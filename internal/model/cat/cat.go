package cat

type Cat struct {
	Id    uint `gorm:"primaryKey"`
	Title string
}
