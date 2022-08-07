/*
*

Ejercicio #2, Andrik Solano Jimenez 2021091379

2)	Escriba el programa más eficiente que pueda para imprimir en pantalla la siguiente figura:


		*
	*	*	*
*	*	*	*	*
	*	*	*
		*


Para dicho fin, escriba y use una función que reciba de parámetro la cantidad de elementos de la línea del centro,
la cual debe ser impar positiva.
*/

package main

import (
	"fmt"
)

func printTriangle(levels int) {

	if levels%2 != 0 && levels > 0 {

		for i := 0; i < levels; i++ {
			for j := levels - 1 - i; j >= 0; j-- {
				fmt.Print(" ")

			}
			for k := 0; k <= i; k++ {
				fmt.Print("* ")

			}
			fmt.Println("")
		}
		for i := 0; i < levels; i++ {

			for j := 0; j <= i+1; j++ {
				fmt.Print(" ")
			}
			for k := levels - 2 - i; k >= 0; k-- {
				fmt.Print("* ")
			}
			fmt.Println("")

		}
		return

	}
	fmt.Println("Please try a positive odd number")
	return

}

func main() {

	printTriangle(3)

}
