package main

import (
	"github.com/labstack/echo/v4"
	
	"calculator/math"
	"fmt"
)

func main() {

	// Echo instance
  e := echo.New()

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(1, 2)
	fmt.Println(sub)
}