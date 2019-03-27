package main

import (
	"fmt"
	"github.com/jinzhu/gorm"

)

func main()  {


	gorm.Open("", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"",
		"",
		8000,
		""))

	fmt.Println("asd")
}
