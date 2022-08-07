/*
*

Ejercicio #1, Andrik Solano Jimenez 2021091379

Haga un programa que cuente e indique el número de caracteres, el número de palabras y el número de líneas
de un texto cualquiera (obtenido de cualquier forma que considere). Considere hacer el cálculo de las tres
variables en el mismo proceso.
*/
package main

import (
	"fmt"
	"strings"
)

func textAnalizer(text string) {

	characters, lines := 0, 1

	words := strings.Split(text, " ")
	lines2 := strings.Split(text, "\n")

	for char := range text {

		if char >= 'a' && char <= 'z' {

			characters++

		} else if char >= 'A' && char <= 'Z' {

			characters++

		} else if char == ' ' {

			characters++

		} else if char == '\n' {

			lines++

		} else {

			characters++
		}
	}

	fmt.Println("In the text are: \n", characters+2, "characters\n", lines, "lines\n", len(words), " words", "\n lines22: ", len(lines2))
}
func main() {

	Texto := "Doña Uzeada de Rizo\n" +
		"Y muy delicioso\n" +
		"me encanta comer pizza"

	textAnalizer(Texto)

}
