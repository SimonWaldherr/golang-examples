package main

import (
	"fmt"
	"github.com/mxk/go-sqlite/sqlite3"
	"os"
	"strconv"
	"time"
)

type Table struct {
	Id   int
	Text string
	Time int64
}

func main() {
	sql, _ := sqlite3.Open("./sqlite.sqlite")
	mode := os.Args[1]
	var qu Table
	if mode == "select" {
		q, err := sql.Query("SELECT * FROM \"main\".\"test\" ORDER BY id DESC LIMIT 1;")
		if err != nil {
			fmt.Printf("Error while getting data: %s\n", err)
		} else {
			err = q.Scan(&qu.Id, &qu.Text, &qu.Time)
			if err != nil {
				fmt.Printf("Error while getting row data: %s\n", err)
			}
			fmt.Println("Last inserted data is: " + qu.Text)
			t := time.Unix(qu.Time, 0)
			fmt.Println("Inserted on " + t.Format(time.RFC1123Z))
		}
	} else {
		if len(os.Args) > 2 {
			data := os.Args[2]
			now := strconv.FormatInt(time.Now().Unix(), 10)
			sql.Query("INSERT INTO \"main\".\"test\" (\"text\", \"time\") VALUES('" + data + "', '" + now + "')")
			fmt.Println(data + " inserted")
		}
	}
}
