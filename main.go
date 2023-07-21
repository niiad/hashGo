package main

import (
	"fmt"
	"hashGo/base"
)

func main() {
	radius := 5.0
	area, err := base.AreaOfCircle(radius)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(area)
}
