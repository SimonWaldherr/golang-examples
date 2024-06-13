package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/postgres"
)

func main() {

	host := "192.168.0.16"
	port := 5432
	user := "postgres-dev"
	password := "s3cr3tp4ssw0rd"
	dbname := "dev"
	sslmode := "disable"

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)

	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	type Article struct {
		Title         string
		TitleTokens   string    `gorm:"type:tsvector;column:title_tokens"`
		PublishDate   time.Time `gorm:"column:publish_date"`
		Summary       string
		SummaryTokens string `gorm:"type:tsvector;column:summary_tokens"`
		TopImage      string `gorm:"column:top_image"`
	}

	db.DropTableIfExists(&Article{})
	db.AutoMigrate(&Article{})
	now := time.Now()

	articles := []Article{
		Article{Title: "Title of Pack my box with five dozen liquor jugs.", PublishDate: now, Summary: "Summary of Pack my box with five dozen liquor jugs.", TopImage: "url"},
		Article{Title: "Title of Jackdaws love my big sphinx of quartz.", PublishDate: now, Summary: "Summary of Jackdaws love my big sphinx of quartz.", TopImage: "url"},
		Article{Title: "Title of The five boxing wizards jump quickly.", PublishDate: now, Summary: "Summary of The five boxing wizards jump quickly.", TopImage: "url"},
		Article{Title: "Title of How vexingly quick daft zebras jump!", PublishDate: now, Summary: "Summary of How vexingly quick daft zebras jump!", TopImage: "url"},
	}

	for _, a := range articles {
		db.Exec("INSERT INTO articles (title, title_tokens, publish_date, summary, summary_tokens, top_image) VALUES (?, to_tsvector(?), ?, ?, to_tsvector(?), ?)", a.Title, a.Title, now, a.Summary, a.Summary, a.TopImage)
	}

	type Result struct {
		Title       string
		PublishDate time.Time
		TopImage    string
	}

	var result Result
	rows, err := db.Raw("SELECT title, publish_date, top_image FROM articles WHERE title_tokens @@ to_tsquery('jump & quick')").Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &result)
		fmt.Println(result)
	}
}
