package entities

type Todo struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

type Todos []Todo
