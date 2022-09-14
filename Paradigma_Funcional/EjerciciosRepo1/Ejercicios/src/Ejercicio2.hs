{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use null" #-}
{-# HLINT ignore "Use guards" #-}
{-# HLINT ignore "Redundant if" #-}
{-# HLINT ignore "Redundant ==" #-}
{-# HLINT ignore "Redundant bracket" #-}

module  Ejercicio2
( main
) where



    {- funcion verificador recorre la lista buscando el elemento indicado

        parametros:

        elemento: elemento buscado
        lista: lista brindada
    
    -}
    verificador elemento lista  
     | lista == [] = False --si es vacia devuelve F
     | elemento == head lista = True --si la cabeza es igual al elemento devuelve True
     | otherwise = 
        verificador elemento (tail lista) --avanza


    {-funcion subconjunto, verifica si los elementos del subconjunto pertenecen al conjunto

     Parametros:
        lista1: sublista buscada
        lista2: lista general
        flag: bandera
    
    -}
    siSubconjunto lista1 lista2 flag =   
        if lista1 == [] then  True --si el subcojunto es vacio entonces es true
        else if flag == False then False --si la bandera es falsa entonces devuelve false
        else
            if (verificador (head lista1) lista2) then --si el verificador devuelve true con el primer elemento del subconjunto
                siSubconjunto (tail lista1) lista2 True --avanza al siguiente elemento del subconjunto
            else
                False --sino devuelve false



    main :: IO ()
    main = do

        print(siSubconjunto [4,2,3] [2,5,3,1,8,4] True)

    --Terminal : True