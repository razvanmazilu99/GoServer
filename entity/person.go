package entity

type Person struct {
	Name string `gorm:"type: varchar(100); not null" json:"name"`
	Age  int `gorm:"type:int"`
	ID string `gorm:"type: varchar(10); not null; primary key"`
}
