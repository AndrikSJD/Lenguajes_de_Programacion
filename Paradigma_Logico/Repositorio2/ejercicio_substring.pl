%
%Ejercicio 1
%
%

%funcion que se encarga de comprobar si la cadena SB es subcadena de P
comparar([],_).
comparar([SB|C1],[P|C2]) :-
    SB = P,
    comparar(C1,C2).



% funcion que se encarga de devolver los valores de la lista en
% donde la cadena brindada es subcadena de los mismos
substring(_,[],[]).
substring(SB,[L|C1],[L|F]) :- %se agrega el valor de la lista si la cadena SB es subcadena de dicho elemento

     string_codes(SB,S1), %se convierten a listas de caracteres
     string_codes(L,S2),
     comparar(S1,S2), %se comparan
     substring(SB,C1,F), !. %se llama recursivo

substring(SB,[_|C1],F) :- %caso contrario se llama con la cola de la lista
     substring(SB,C1,F), !.



%substring("hol",["hola","holanda","hosa"],X).
%X = ["hola", "holanda"].




