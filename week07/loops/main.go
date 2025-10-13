package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month time.Time = time.Time(now.Month())
	var day int = now.Day()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

	univ := "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
