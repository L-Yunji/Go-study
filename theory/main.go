// Variables and Constants

package main

import (
	"fmt"
	"strings"
)

func superAdd(numbers ...int) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, up := lenAndUpper("yunji")
	fmt.Println(totalLenght, up)
}
