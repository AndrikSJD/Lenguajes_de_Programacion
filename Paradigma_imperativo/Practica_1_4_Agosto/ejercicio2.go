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

func printDiamond(levels int) {

	if levels%2 != 0 && levels > 0 {

		for row := 0; row < levels; row += 2 {

			var whiteSpacesLimit int = levels - row
			for empty := 0; empty <= whiteSpacesLimit; empty++ {
				fmt.Print(" ")

			}
			lessAsterisk := row
			for column := 0; column <= lessAsterisk; column++ {
				fmt.Print("* ")
			}
			fmt.Println("")

		}
		for row := 0; row < levels-2; row += 2 {

			var whiteSpacesLimit int = 4 + row
			for empty := 0; empty <= whiteSpacesLimit-1; empty++ {
				fmt.Print(" ")

			}
			lessAsterisk := (levels - whiteSpacesLimit) + 1
			for column := 0; column <= lessAsterisk; column++ {
				fmt.Print("* ")
			}
			fmt.Println("")

		}
	} else {
		fmt.Println("Please try a positive odd number")
		return
	}

}

func main() {

	printDiamond(5)

}

/** RESULTADO
      *
    * * *
  * * * * *
    * * *
      *
*/
