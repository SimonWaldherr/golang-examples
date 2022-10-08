package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	MinuteCheck := time.Now().Minute() + 1
	SpecVal := fmt.Sprintf("%d", MinuteCheck) + " * * * *"
	fmt.Println("The SpecVal is ", SpecVal)
	_, err := c.AddFunc(SpecVal, PrintNumber)
	if err != nil {
		fmt.Println("Could not register cron - PrintNumber. Error:", err)
	}
	c.Start()
	time.Sleep(time.Second * 70)
}

func PrintNumber() {
	fmt.Println("The function was triggered at time:", time.Now())
}
