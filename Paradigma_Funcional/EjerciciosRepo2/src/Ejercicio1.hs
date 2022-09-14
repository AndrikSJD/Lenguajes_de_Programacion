{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Eta reduce" #-}
{-# HLINT ignore "Redundant bracket" #-}
module Ejercicio1
    ( main
    ) where

data Producto = Producto { nombre :: String
                    , cantidad :: Int
                    , precio :: Int
                     } deriving Show


listarExistenciaMinima minimo listaProductos =
    filter minimos listaProductos
    where
        minimos prod = ((cantidad prod) <= minimo)

--calcularMontoVenta nombreProducto cantidad
--CALCULAR EL MONTO A PAGAR POR LA CANTIDAD DE PRODUCTOS INDICADA MÁS EL 13% DE IMPUESTO

--funcion para aplicar impuesto de venta a productos cuyos valores sean superiores a 2000
aplicarIvaProducto precio =
    if precio > 2000 then
        precio + (precio* 13 `div` 100)
    else
        precio



--funcion para calcular el monto total a pagar, aplicandole iva a los productos con precio superior a 2000
montoTotalSinImpuesto listaProductos montoTotal=
    if (null listaProductos) then montoTotal
        else
            montoTotalSinImpuesto (tail listaProductos) (montoTotal+ ((cantidad(head listaProductos))*(precio((head listaProductos)))))

--funcion para calcular el monto total a pagar, aplicandole iva a los productos con precio superior a 2000
montoVentaTotal listaProductos montoTotal=
    if (null listaProductos) then montoTotal
        else
            montoVentaTotal (tail listaProductos) (montoTotal+ ((cantidad(head listaProductos))*(aplicarIvaProducto (precio((head listaProductos))))))


montoVentaTotalConImpuesto montoTotal  =  montoTotal + (montoTotal * 13 `div` 100) 



main :: IO ()
main = do
    let listaProductos=[
            (Producto "arroz" 15 2500),
            (Producto "frijoles" 4 2000),
            (Producto "leche" 8 1200),
            (Producto "cafe" 12 4500)]

    print(listarExistenciaMinima 10 listaProductos)

    print("El monto total a pagar por los productos sin impuestos es: ")
    print(montoTotalSinImpuesto listaProductos 0)

    print("El monto total a pagar por los productos con impuestos a los productos mayores a 2000 es: ")
    print(montoVentaTotal listaProductos 0)


    {-
    Terminal
    [Producto {nombre = "frijoles", cantidad = 4, precio = 2000},Producto {nombre = "leche", cantidad = 8, precio = 1200}]
    "El monto total a pagar por los productos sin impuestos es: "
    109100
    "El monto total a pagar por los productos con impuestos a los productos mayores a 2000 es: "
    120995
    -}