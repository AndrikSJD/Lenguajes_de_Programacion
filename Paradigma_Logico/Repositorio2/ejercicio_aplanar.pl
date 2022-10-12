%Funcion aplanarL:
% Se encarga de aplanar una lista anidada dando como resultado una lista
% con todos los elementos
%CA: cabeza
%CO: cola
%L: nueva lista aplanada
%K: cola de la sublista

aplanarL([],[]).    %una lista vacia devuelve vacio ya que ya esta aplanada
aplanarL([[CA|CO]|K],L):- aplanarL([CA|[CO|K]],L). %si la cabeza es una sublista, extrae la cabeza y se llama con la cola
aplanarL([[]|CO],NL) :- aplanarL(CO,NL). %si la cabeza es una lista vacia, llama a la funcion con la cola
aplanarL([CA|CO],[CA|K]):- aplanarL(CO,K). %si se trata de un elemento simple lo agrega a la lista final y llama la funcion con la cola
