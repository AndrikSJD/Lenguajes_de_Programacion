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

	characters := 0

	words := strings.Split(text, " ")
	lines := strings.Split(text, "\n")

	for range text {
		characters++
	}

	fmt.Println("In the given text are: \n", characters, "characters\n", len(lines), "lines\n", len(words), " words")
}
func main() {

	Texto :=
		`Soy ginebrino de nacimiento, y mi familia es una de las más distinguidas de esa república. 
Durante muchos años mis antepasados habían sido consejeros y jueces, y mi padre había ocupado con gran honor 
y buena reputación diversos cargos públicos. Todos los que lo conocían lo respetaban por su integridad e infatigable 
dedicación. Pasó su juventud dedicado por completo a los asuntos de su país, y sólo al final de su vida pensó en el 
matrimonio y así dar al Estado unos hijos que pudieran perpetuar su nombre y sus virtudes.`
	textAnalizer(Texto)

}

/** RESULTADO TERMINAL:
In the given text are:
527 characters
5 lines
91  words
*/
