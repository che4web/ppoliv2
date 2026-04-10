package models

type HeroCard struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:120;not null"`
	Description string `gorm:"size:220;not null"`
	LinkText    string `gorm:"size:40;not null"`
}

type Feature struct {
	ID    uint   `gorm:"primaryKey"`
	Icon  string `gorm:"size:16;not null"`
	Title string `gorm:"size:160;not null"`
}

type Step struct {
	ID     uint   `gorm:"primaryKey"`
	Number int    `gorm:"not null"`
	Title  string `gorm:"size:200;not null"`
}

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:200;not null"`
	Price string `gorm:"size:40;not null"`
	Meta  string `gorm:"size:120;not null"`
}

type Review struct {
	ID     uint   `gorm:"primaryKey"`
	Text   string `gorm:"size:300;not null"`
	Author string `gorm:"size:80;not null"`
	Date   string `gorm:"size:60;not null"`
}
