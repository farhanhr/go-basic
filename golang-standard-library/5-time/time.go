package main
import (
	"time"
	"fmt"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2009, time.August, 18, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	// di golang untuk melakukan format date, cukup unik yaitu dengan menspesifikasikan seperti yang tertulis
	// pada dokumentasi, misalnya jika ingin format yyyy maka perlu menuliskan "2006" pada formatter
	formatter := "2006-01-02 15:04:05"

	value := "2025-02-22 23:27:27"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
		} else {
		fmt.Println(valueTime)
	}
}