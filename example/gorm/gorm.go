package main

import (
	"fmt"
	"github.com/SamuelTissot/sqltime"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
)

// redefine the gorm base model
type Base struct {
	ID        uint          `gorm:"primary_key"`
	CreatedAt sqltime.Time  `gorm:"type:timestamp"`
	UpdatedAt sqltime.Time  `gorm:"type:timestamp"`
	DeletedAt *sqltime.Time `gorm:"type:timestamp"`
}

type TestModel struct {
	Base
	Data string
}

func main() {
	// change the dataSourceName value
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//setup database
	setup(db)

	//create model
	e := TestModel{
		Data: "foo bar fizz",
		Base: Base{
			CreatedAt: sqltime.Now(),
			UpdatedAt: sqltime.Now(),
		},
	}
	db.Create(&e)

	// fetch the same model
	var f TestModel
	db.Where(&TestModel{Base: Base{ID: e.ID}}).First(&f)

	// compare
	if reflect.DeepEqual(e, f) {
		fmt.Println("Success the models are equal")
	} else {
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println("boo, not equal")
	}
}

func setup(db *gorm.DB) {
	db.Exec(`
			CREATE TABLE IF NOT EXISTS test_models
			(
				id           SERIAL PRIMARY KEY ,
				data          varchar(255),
				created_at    timestamp    NOT NULL DEFAULT now(),
				updated_at    timestamp    NOT NULL DEFAULT now(),
				deleted_at    timestamp
			);
		`)
}
