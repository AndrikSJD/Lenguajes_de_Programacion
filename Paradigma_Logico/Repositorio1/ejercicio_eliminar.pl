
%Ejercicio de eliminar una lista

%E = elemento
%C = cola
%LF = lista final (sin el elemento
%Z = cabeza diferente
%C1 = cola 1
%C2 = cola 2



eliminarElemento(_, [ ], [ ]). %En caso de que la lista este vacia se devuelve
eliminarElemento(E, [E|C], LF) :-  eliminarElemento(E,C,LF),!. %En caso de que la cabeza de la lista sea igual al elemento se llama con cola
eliminarElemento(E, [Z|C1], [Z|C2]) :- eliminarElemento(E, C1, C2).%En caso de que la cabeza sea distinta del elemento, el elemento se le adjunta
                                                          %a la nueva lista y esto sucede si solo si se continua buscando eliminar en la cola
                                                          %de la nueva lista
