package main

import (
	"fmt"
)

type persona struct {
	nombre    string
	apellido  string
	cromosoma []byte
}

var lista_personas_array [5]persona
var lista_personas_slice []persona

func crearPersona(n string, a string, c []byte) persona {
	p := persona{nombre: n, apellido: a, cromosoma: c} //funcion para gregar personas en el slice
	return p
}

func agregarPersona(p persona) {
	lista_personas_slice = append(lista_personas_slice, p) //devuelve, cambia el puntero hay que devolver
}

func llenar1() { //datos arreglos quemado
	lista_personas_array[0].nombre = "Oscar1"
	lista_personas_array[0].apellido = "Viquez1"
	lista_personas_array[0].cromosoma = []byte{'1', '0', '1', '0', '0', '1', '1', '1', '0', '1'}

	lista_personas_array[1].nombre = "Oscar2"
	lista_personas_array[1].apellido = "Viquez2"
	lista_personas_array[1].cromosoma = []byte{'1', '1', '1', '0', '1', '0', '1', '0', '0', '1'}

	lista_personas_array[2].nombre = "Oscar3"
	lista_personas_array[2].apellido = "Viquez3"
	lista_personas_array[2].cromosoma = []byte{'0', '1', '1', '0', '0', '1', '0', '1', '1', '0'}

	lista_personas_array[3].nombre = "Oscar4"
	lista_personas_array[3].apellido = "Viquez4"
	lista_personas_array[3].cromosoma = []byte{'0', '0', '1', '1', '0', '0', '0', '1', '0', '0'}

	lista_personas_array[4].nombre = "Oscar5"
	lista_personas_array[4].apellido = "Viquez5"
	lista_personas_array[4].cromosoma = []byte{'1', '1', '0', '1', '1', '1', '0', '0', '0', '1'}
}

func llenar2() { //meter datos slice, solo se metio 1
	agregarPersona(crearPersona("Oscar1", "Viquez1", []byte{'1', '0', '1', '0', '0', '1', '1', '1', '0', '1'}))
}

// PRIMERA IMPLEMENTACIÓN primeras 3 con arreglo
// recorrido tradicional con range... pruebe hacerlo utilizando la función len(arreglo)
func masParecido(lista [5]persona, muestra [10]byte) persona {
	cont, max := 0, 0 //max cantidad genes iguales
	masparecido := 0
	for i, individuo := range lista {
		cont = 0
		for j, gen := range individuo.cromosoma {
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido]
}

// SEGUNDA IMPLEMENTACIÓN
// usando punteros a arreglos... un poco más difícil de leer
// se recibe un puntero a la lista de personas
func masParecido2(lista *[5]persona, muestra [10]byte) persona {
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range *lista {
		cont = 0
		for j, gen := range individuo.cromosoma {
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido]
}

// TERCERA IMPLEMENTACIÓN
// USANDO PUNTEROS A LOS REGISTROS DE CADA UNO DE LOS ELEMENTOS DEL ARREGLO
// variable | tipo
func (person *persona) verificar(muestra [10]byte) int {
	cont := 0
	for j, gen := range person.cromosoma { //puntero cada uno de los elementos de la lista persona
		if gen == muestra[j] {
			cont++
		}
	}
	return cont
}

func masParecido3(lista [5]persona, muestra [10]byte) persona {
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range lista {
		cont = individuo.verificar(muestra)
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido]
}

// CUARTA IMPLEMENTACIÓN
// Usando Slices... Es mejor en Go utilizar Slice en vez de punteros y hay más flexibilidad de trabajar con el arreglo
func masParecido4(lista []persona, muestra []byte) persona {
	cont, max := 0, 0
	masparecido := 0
	for i, individuo := range lista {
		cont = 0
		//for =  indice | elemnt
		for j, gen := range individuo.cromosoma {
			if gen == muestra[j] {
				cont++
			}
		}
		if cont > max {
			masparecido = i
			max = cont
		}

	}
	return lista[masparecido]
}

func main() {
	llenar1()
	llenar2()
	fmt.Printf("Cromosoma Persona #2 ARRAY: %s\n", string(lista_personas_array[1].cromosoma))
	//fmt.Println("Cromosoma Persona #2 ARRAY: ", string(lista_personas_array[1].cromosoma))
	//getchar()

	// + = , en impresiones
	fmt.Println("Más parecido: ", masParecido(lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	fmt.Println("Más parecido: ", masParecido2(&lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	fmt.Println("Más parecido: ", masParecido3(lista_personas_array, [10]byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	//fmt.Println(masParecido4(lista_personas_array[:], []byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))

	fmt.Printf("Cromosoma Persona #1 SLICE: %s\n", string(lista_personas_array[1].cromosoma))

	fmt.Println("Más parecido: ", masParecido4(lista_personas_slice, []byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
	agregarPersona(crearPersona("Oscar2", "Viquez2", []byte{'1', '1', '0', '1', '0', '0', '1', '1', '0', '1'}))
	fmt.Println("Más parecido: ", masParecido4(lista_personas_slice, []byte{'1', '0', '0', '1', '0', '0', '0', '1', '1', '0'}))
}
