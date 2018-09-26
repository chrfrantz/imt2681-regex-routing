package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex,_ :=regexp.Compile("[A-Z][a-zå-ø]*$")

	res := regex.FindAllString("Velkommen i Gjøvik", -1);

	fmt.Println(res)

}
