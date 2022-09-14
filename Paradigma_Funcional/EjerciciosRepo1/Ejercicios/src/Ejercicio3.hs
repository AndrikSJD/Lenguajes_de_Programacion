{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use null" #-}

module  Ejercicio3
( main
) where



    {-Funcion agregarE, agrega los elementos de una lista a una nueva lista
    
    Parametros:
        subLista : sublista de la lista anidada
        listaPlana: nueva lista
    -}
    agregarE :: Eq a => [a] -> [a] -> [a]
    agregarE subLista listaPlana =
        if subLista == [] then listaPlana --si es vacio es decir ya se agregaron todos, devuelve la nueva lista
        else
            agregarE (tail subLista) (listaPlana ++ [head subLista]) --se llama a la funcion con la con la cola de la lista y a la nueva lista se le suma la cabeza de la lista


    {-funcion aplanarLista

    Parametros:
        listaAnidada: lista anidada brindada
        listaPlana: nueva lista aplanada
    
    -}
    aplanarLista :: Eq a => [[a]] -> [a] -> [a]
    aplanarLista listaAnidada listaPlana  =
        if listaAnidada == [] then  listaPlana --si la lista anidada esta vacia devuelve la aplanada
        else
            aplanarLista (tail listaAnidada) (agregarE (head listaAnidada) listaPlana) --llama a la funcion con la cola de la anidada y con la llamada a la funcion
                                                                                    --de la agregacion de los elementos de la cabeza de la anidada a la listaPlana



    main :: IO ()
    main = do

        let listaAnidada = [[2],[3,4],[5], [7,8], [9], [10,11,12,13]]

        print(aplanarLista listaAnidada [])

    --Terminal: [2,3,4,5,7,8,9,10,11,12,13]