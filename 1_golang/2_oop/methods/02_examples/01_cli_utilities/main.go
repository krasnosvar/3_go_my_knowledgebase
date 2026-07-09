package main

import "fmt"

type Version struct{ Major int }

func (v Version) String() string { return fmt.Sprintf("v%d", v.Major) }
func main()                      { fmt.Println(Version{1}.String()) }
