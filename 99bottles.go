package main

import (
	"fmt"
	"strconv"
)

func main() {
	key := "ff0965a16d38556b5e961784090b60c2"
	for i := 100; i > 0; i-- {
		amt := strconv.Itoa(i)
		fmt.Println(amt + " bottles of beer on the wall, " + amt + " bottles of beer, take one down, pass it around, " + amt + " bottles of beer on the wall!")
	}

}
