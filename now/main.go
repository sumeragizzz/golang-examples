package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	timezone := flag.String("timezone", "Asia/Tokyo", "Timezone(ex: Asia/Tokyo)")
	flag.Parse()

	location, err := time.LoadLocation(*timezone)
	if err != nil {
		panic(err)
	}

	now := time.Now().In(location)
	fmt.Println(now.Format("2006/01/02 15:04:05 MST"))
}
