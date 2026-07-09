package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	regionFlag := flag.String("region", "", "region")
	flag.Parse()
	region := "eu-central-1"
	if v := os.Getenv("APP_REGION"); v != "" {
		region = v
	}
	if *regionFlag != "" {
		region = *regionFlag
	}
	fmt.Println(region)
}
