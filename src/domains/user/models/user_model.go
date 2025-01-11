package models

type User struct {
	Id         int    `gorm:"id"`
	TelegramId int64  `gorm:"telegram_id"`
	Password   string `gorm:"password"`
}
