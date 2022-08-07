package main

import (
	"fmt"
)

func rotateList(listToRotate []string, loopNumber int, directionOfRotation string) {

	var originalList []string = listToRotate

	if directionOfRotation == "izq" {

		for i := 0; i < len(listToRotate); i++ {

			originalList[i-1] = listToRotate[i]

		}

	} else {

		fmt.Println("Right")

	}

}

func main() {

	char_list := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	rotateList(char_list, 3, "izq")

}
