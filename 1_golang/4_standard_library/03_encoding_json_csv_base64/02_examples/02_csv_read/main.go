package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	r := csv.NewReader(strings.NewReader("id,name\n1,go\n"))
	recs, _ := r.ReadAll()
	fmt.Println(recs[1][1])
}
