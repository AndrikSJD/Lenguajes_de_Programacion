package main

import (
	"fmt"
)

/**
4)	Escriba un programa que utilice una estructura y un arreglo/slice para almacenar en memoria un inventario de una tienda que vende zapatos.
Cada zapato debe contar con la información de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44. La estructura debe
llamarse "calzado". El programa debe poder almacenar la información “quemada” para 10+ zapatos del inventario y un stock de N cantidad de unidades de
dicho zapato (quiere decir por ejemplo que puede existir en inventario el zapato marca Nike, talla 42 y que cuesta 50000 colones, pero además que puede
existir no solo un par de esos, si no N pares del mismo zapato).

Haga una función que venda zapatos (eliminando del stock cada vez que haya venta e indicando que no se puede vender porque ya o hay stock) y pruébela
dentro del main haciendo las ventas que usted considere para demostrar su funcionamiento.

*/
type calzado struct {
	marca    string
	talla    int
	precio   int
	cantidad int
}

//Slice del inventario general
var inventoryCalzado = []calzado{{"nike", 34, 50000, 4},
	{"adidas", 37, 60000, 7}, {"Gucci", 39, 90000, 3},
	{"DrMarteens", 42, 12000, 10}, {"Reebook", 44, 29500, 7},
	{"Champion", 35, 35000, 15}, {"UnderArmour", 43, 80000, 5},
	{"Timberland", 36, 92000, 17}, {"Assics", 39, 78000, 20},
	{"Vans", 37, 24500, 25}, {"Caterpillar", 40, 105000, 7}}

//Funcion para buscar un calzado en la lista del inventario, retorna el indice del mismo en el inventario
func searchCalzado(inventory []calzado, brand string, size int) int {

	for i, shoe := range inventory {
		if shoe.marca == brand && shoe.talla == size {
			return i
		}

	}
	return -1
}

//Funcion para añadir un calzado al inventario
func addCalzado(brand string, size int, price int, stock int) {
	if size >= 34 && size <= 44 {
		shoe := calzado{marca: brand, talla: size, precio: price, cantidad: stock}
		inventoryCalzado = append(inventoryCalzado, shoe)
	} else {
		fmt.Println("Talla de calzado invalida, solo se aceptan tallas de la 34 a la 44")
	}

}

//Function para vender zapatos
func sellCalzado(inventory *[]calzado, brand string, size int) {

	shoe := searchCalzado(*inventory, brand, size)
	if shoe != -1 {

		if (*inventory)[shoe].cantidad >= 1 {
			(*inventory)[shoe].cantidad -= 1
			fmt.Println("Compra de " + brand + " realizada con exito")

		} else {
			fmt.Println("No se pudo concretar la venta debido a la falta de stock del calzado " + brand)

		}

	} else {
		fmt.Println("No se pudo concretar la venta debido a que no se encuentra el calzado buscado")

	}

}

func main() {

	sellCalzado(&inventoryCalzado, "nike", 34)
	sellCalzado(&inventoryCalzado, "nike", 34)
	sellCalzado(&inventoryCalzado, "nike", 34)
	sellCalzado(&inventoryCalzado, "nike", 34)
	sellCalzado(&inventoryCalzado, "nike", 34)
	sellCalzado(&inventoryCalzado, "Gucci", 39)
	sellCalzado(&inventoryCalzado, "Gucci", 39)
	sellCalzado(&inventoryCalzado, "Gucci", 39)
	sellCalzado(&inventoryCalzado, "Gucci", 39)
	sellCalzado(&inventoryCalzado, "Vans", 37)
	sellCalzado(&inventoryCalzado, "Vans", 37)
	sellCalzado(&inventoryCalzado, "Vans", 37)
	sellCalzado(&inventoryCalzado, "Vans", 37)

}

/** Resultado terminal
Compra de nike realizada con exito
Compra de nike realizada con exito
Compra de nike realizada con exito
Compra de nike realizada con exito
No se pudo concretar la venta debido a la falta de stock del calzado nike
Compra de Gucci realizada con exito
Compra de Gucci realizada con exito
Compra de Gucci realizada con exito
No se pudo concretar la venta debido a la falta de stock del calzado Gucci
Compra de Vans realizada con exito
Compra de Vans realizada con exito
Compra de Vans realizada con exito
Compra de Vans realizada con exito
*/
