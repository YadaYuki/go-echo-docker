package entities

type Todo struct {
	ID    int    `gorm:"column:id;primaryKey"`
	Title string `gorm:"column:title"`
}

type Todos []Todo
