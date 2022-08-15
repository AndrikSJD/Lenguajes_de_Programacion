package main

import (
	"fmt"
	"sort"
)

/**
Amplie el funcionamiento del ejercicio de Productos visto en clase para que el programa ahora permita:
a.	A partir de la lista de productos con mínimas existencias de inventario, ampliar dicho inventario con la
compra de más unidades de dicho producto hasta que cumplan con el mínimo establecido de manera constante. Se sugiere crear una
función denominada “aumentarInventarioDeMinimos(listaMinimos)”.

b.	Crear una función que ordene la lista de productos usando como llave para el ordenamiento cualquiera de los elementos de la
estructura producto. La lista/slice debe quedar modificada al finalizar el método. Se solicita investigar y hacer uso de alguna función predefinida
de alguna librería del lenguaje Go.

*/

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

//crear tipo de dato lsta productos
type listaProductos []producto

//variable tipo lproducto para usar l* despues
var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	producto := producto{nombre: nombre, cantidad: cantidad, precio: precio}
	comprobarProducto := (*l).buscarProducto(nombre)
	if comprobarProducto != -1 {
		(*l)[comprobarProducto].cantidad += cantidad
		if (*l)[comprobarProducto].precio != precio {
			(*l)[comprobarProducto].precio = precio
		}

	} else {
		*l = append(*l, producto)
	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
			if (*l)[prod].cantidad == 0 {
				*l = append((*l)[:prod], (*l)[prod+1:]...)

			}
		} else {
			fmt.Println("No se puede vayor cantidad de productos que los que hay en existencia")
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMinimos() listaProductos {

	var listaMinimos []producto // tipo listaProductos == []producto
	for i, producto := range *l {

		if producto.cantidad <= existenciaMinima {
			listaMinimos = append(listaMinimos, (*l)[i])
		}

	}
	return listaMinimos
}

func aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for i, producto := range listaMinimos {
		actualExistence := producto.cantidad
		listaMinimos[i].cantidad += existenciaMinima - actualExistence

	}

}

func (l *listaProductos) ordernarArreglo() {
	sort.Slice(*l, func(x, y int) bool {
		return (*l)[x].cantidad > (*l)[y].cantidad
	})

}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	lProductos.ordernarArreglo()
	fmt.Println(lProductos)
	listaMinimos := lProductos.listarProductosMinimos()
	aumentarInventarioDeMinimos(listaMinimos)
	fmt.Println(listaMinimos)
	//lProductos.venderProducto("arroz", 15)
	//fmt.Println(lProductos)
	//fmt.Println(lProductos)
	//lProductos.agregarProducto("azucar", 20, 1500)
	//fmt.Println(lProductos)
	//lProductos.venderProducto("frijoles", 4)
	//fmt.Println(lProductos)
	//lProductos.venderProducto("leche", 10)
}

/** Resultado terminal
[{arroz 30 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500}]
[{arroz 30 2500} {café 12 4500} {leche 8 1200} {frijoles 4 2000}]
[{leche 10 1200} {frijoles 10 2000}]
*/

//tener cuidado con los archivos a la manera de escribir, ver que sea agregar y no sobreescribir porque sino en cada iteracion del ciclo se sobreescribe
//uno sobre otro. Indagar sobre permisos de lectura.

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los los productos se almacenen en archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista
