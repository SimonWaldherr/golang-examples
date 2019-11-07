package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
	//1401403874
	fmt.Println(time.Now().Format(time.RFC850))
	//Friday, 30-May-14 00:51:14 CEST
	fmt.Println(time.Now().Format(time.RFC1123Z))
	//Fri, 30 May 2014 00:51:14 +0200
	fmt.Println(time.Now().Format("20060102150405"))
	//20140530005114
	fmt.Println(time.Unix(1401403874, 0).Format("02.01.2006 15:04:05"))
	//30.05.2014 00:51:14
	fmt.Println(time.Unix(123456789, 0).Format(time.RFC3339))
	//1973-11-29T22:33:09+01:00
	fmt.Println(time.Unix(1234567890, 0).Format(time.RFC822))
	//14 Feb 09 00:31 CET
}
