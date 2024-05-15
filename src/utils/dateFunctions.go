package utils

import (
	"fmt"
	"time"
)

func GetYesterday() string {
	a := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println("Yesterday ", a)
	return a
}
