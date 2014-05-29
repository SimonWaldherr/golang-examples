package main

import (
	"github.com/mxk/go-sqlite/sqlite3"
	"strconv"
	"time"
	"fmt"
	"os"
)

type Table struct {
	Id int
	Text string
	Time int64
}

func main() {
	sql, _ := sqlite3.Open("./sqlite.sqlite")
	var qu Table
	if os.Args[1] == "select" {
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
		now := strconv.FormatInt(time.Now().Unix(), 10)
		sql.Query("INSERT INTO \"main\".\"test\" (\"text\", \"time\") VALUES('" + os.Args[2] + "', '" + now + "')")
		fmt.Println(os.Args[2] + " inserted")
	}
}