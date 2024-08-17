package main

import (
	"fmt"
	"os"

	"udemy.course.dev/v1/info"
)

func pointer_test() {
	gb_value := 4096
	var gb_value_pointer *int

	gb_value_pointer = &gb_value
	crud_pointer_value(gb_value_pointer)

}

func crud_pointer_value(gb_value *int) {
	var choice int
	var tmp_value int
	info.Pointer_menu_info()
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Pointer value ", *gb_value)
	case 2:
		fmt.Scan(&tmp_value)
		*gb_value = *gb_value + tmp_value
		fmt.Println("New pointer value", *gb_value)
	case 3:
		fmt.Scan(&tmp_value)
		*gb_value = *gb_value - tmp_value
		fmt.Println("New pointer value", *gb_value)
	case 4:
		os.Exit(0)
	}

}
