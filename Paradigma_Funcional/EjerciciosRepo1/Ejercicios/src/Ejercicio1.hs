{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Redundant map" #-}
{-# HLINT ignore "Use guards" #-}
{-# HLINT ignore "Redundant bracket" #-}
{-# HLINT ignore "Redundant if" #-}
{-# HLINT ignore "Redundant ==" #-}
{-# HLINT ignore "Eta reduce" #-}
module  Ejercicio1
    ( main
    ) where


    {-Funcion validarF, devuelve un valor booleano sobre si una subcadena se encuentra dentro de una lista de cadenas

    Parametros:
        subcadena:subcadena buscada
        elemento: elemento de la lista actual
        indice: indice de la subcadena
        indice2: indice del elemento
        flag: bandera

    -}
    validarF subcadena elemento indice indice2 flag  =
        if ((indice + 1) > (length subcadena))  then True --Valida si la subcadena ya se recorrio toda es decir que ya se encontraron los elementos dentro de la cadena 
            else if ((indice2+1) > (length elemento)) --Valida si el elemento de la lista ya se recorrio toda mientras que la subcadena no, lo que indica que 
                then False                            --la subcadena no se encuentra en dicho elemento y devuelve false
                else
                    if (subcadena !!indice) == (elemento !!indice2) then --compara si el caracter de la subcadena en el indice x condicide con el caracter del elemento en el indice y 
                        validarF (subcadena) (elemento) (indice+1)(indice2+1) True --vuelve a llamar la funcion con ambos indices en la siguiente posicion y con la bandera en true
                    else --si no son iguales
                        if (flag == True && (indice2+1) < (length elemento) && indice > 0) then --compara si la bandera es true es decir en la iteracion anterior si eran iguales, pero aun le queda por recorrer al elemento entonces podria coincidir mas adelante
                            validarF (subcadena) (elemento) (indice-1)(indice2+1) False  --ademas pregunta si el indice es mayor a 0, es decir que ya se habia avanzado, entonces lo devuelve hacia atras para reiniciarlo
                        else
                            validarF (subcadena) (elemento) (indice)(indice2+1) False --caso normal, sigue recorriendo el elemento en busca de coincidencias


    siContiene subc lista = filter (\x -> validarF subc x 0 0 False ) lista


    main :: IO ()
    main = do
        let lista = ["gelatina", "laguna", "cola", "hielo", "gulhumna", "ba", "le"]

        print(siContiene "la" lista)

--Terminal : ["gelatina","laguna","cola"]