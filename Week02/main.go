package main

import (
	"Goworks/Week02/Service"
	"fmt"
)

func main() {

	value, err := Service.ContinueWrapTheError()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%s", value)
}
