package main

import (
	"github.com/gnewton/mesh2sqlite3/lib"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	db, err := gorm.Open("sqlite3", "mesh2019_sqlite3.db")
	if err != nil {
		log.Fatal(err)

	}

	count := 0
	for i := 0; i < 14; i++ {
		var trees []*lib.MeshTree
		db.Where("depth = ?", i).Find(&trees)
		log.Println("")
		log.Println("------")
		log.Println(i)
		log.Println(len(trees))
		log.Println(trees)
		count += len(trees)
	}
	log.Println("Count=")
	log.Println(count)

}
