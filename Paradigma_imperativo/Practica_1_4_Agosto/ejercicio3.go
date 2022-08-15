package main

import (
	"fmt"
)

/**
3)	Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la derecha según sea el caso en función al parámetro
que se reciba. Los parámetros deben ser un puntero a un arreglo previamente creado, la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).

A partir de esta función, escriba un programa que haga varias rotaciones cualesquiera de una secuencia de elementos previamente creada que sea inmutable.
Al final debe imprimirse el resultado de cada rotación además de la secuencia original.

*/

//funcion para invertir lista
func inverse(array []string) {
	length := len(array)
	for i := 0; i < length/2; i++ {
		(array)[i], (array)[length-1-i] = (array)[length-1-i], (array)[i]

	}
}

//funcion para rotar lista
func rotateList(listToRotate *[]string, loopNumber int, directionOfRotation string) {
	originalList := make([]string, len(*listToRotate))
	originalList = *listToRotate

	if directionOfRotation == "izq" {
		for ; loopNumber > 0; loopNumber-- {
			left := (*listToRotate)[0]
			*listToRotate = (*listToRotate)[1:]
			*listToRotate = append(*listToRotate, left)
		}
		fmt.Println("The original list was: ", originalList)
		fmt.Println("The list was rotated: ", loopNumber, " times to the left")
		fmt.Println("The new looped list is: ", *listToRotate)

	} else if directionOfRotation == "der" {
		fmt.Println("The original list was: ", originalList)
		loopNumber = loopNumber % len(*listToRotate)
		if loopNumber == 0 {
			return
		}
		inverse(*listToRotate)
		inverse((*listToRotate)[:loopNumber])
		inverse((*listToRotate)[loopNumber:])
		fmt.Println("The list was rotated: ", loopNumber, " times to the right")
		fmt.Println("The new looped list is: ", *listToRotate)

	}

}

func main() {

	charList := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	charList2 := []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	rotateList(&charList, 3, "der")
	rotateList(&charList2, 3, "izq")

}

/** Resultado
The original list was:  [a b c d e f g h]
The list was rotated:  3  times to the right
The new looped list is:  [f g h a b c d e]
The original list was:  [1 2 3 4 5 6 7 8]
The list was rotated:  0  times to the left
The new looped list is:  [4 5 6 7 8 1 2 3]

*/
