package jwd_test

import (
	"fmt"
	"github.com/9uuso/go-jaro-winkler-distance"
)

func main() {
	res := jwd.Calculate("DIXON", "DICKSONX")
	fmt.Println(res)
}
