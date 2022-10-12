
%Funcion Miembro:
%Indica si el elemento X forma parte de la lista Y
%X: elemento
%Y: lista de elementos

miembro(X,[X|_]):- !. %si el elemento es igual a la cabeza de la lista
miembro(X,[_|Y]):- miembro(X,Y). %si es diferente se llama recursivo con la cola


%Funcion subconjunto:
%Indica si la primer lista brindada es subconjunto de la segunda
%
%X: cabeza del subconjunto
%C1: cola del subconjunto
%L2: lista general que contiene el subconjunto
%

subconjunto([],_). %si la lista del subconjunto esta vacia es verdadero.
subconjunto([X|C1], L2) :- miembro(X,L2), subconjunto(C1,L2),!. %es subconjunto si                               %la cabeza es miembro de la lista y la cola tambien


