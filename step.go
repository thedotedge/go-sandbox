package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Package struct {
	Id         int        `gorm:"primary_key"`
	Name       string     `sql:"not null;unique;size:255"`
	Url        string     `sql:"not null;size:255"`
	Created_at time.Time
}

func (p Package) TableName() string {
	return "api_package"
}

func (p Package) String() string {
	return fmt.Sprintf("id=%d name=%s url=%s", p.Id, p.Name, p.Url)
}

func main() {
	db, err := gorm.Open("mysql", "root@/bower?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true) // Disable table name's pluralization
	log.Printf("Database connected")
	db.CreateTable(&Package{})

	r := mux.NewRouter()
	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("%s - %s", request.RequestURI, request.UserAgent())
		fmt.Fprintf(response, "Hello world!")
	});
	r.HandleFunc("/insert", func(response http.ResponseWriter, request *http.Request) {
		bowerPackage := Package{Name: "text", Url: "https://github.com/golang", Created_at: time.Now()}
		db.NewRecord(bowerPackage)
		if err := db.Create(&bowerPackage).Error; err != nil {
			fmt.Fprintf(response, err.Error())
			log.Fatal(err.Error())
		} else {
			fmt.Fprintf(response, "%s", bowerPackage)
		}
	});
	http.Handle("/", r)

	var port string;
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "3001";
	}
	log.Printf("Listening on port " + port)
	http.ListenAndServe(":" + port, nil)
}
