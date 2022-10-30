
from Num import *
from Op import *
from CaracterValidoException import *
from FormatoValidoException import *
import re

listaNum = []

listaOp = []

def leerExp():
    try:
        expresion = input("Ingrese la expresion matematica: ")
        expresion_without_spaces = expresion.replace(" ", "") #quitar los espacios
        list = re.split("(?=[-+*/])|(?<=[-+*/])", expresion_without_spaces) #expresion regular para convertir en lista
        
        for i in range(0,len(list)):
            if (i%2 != 0 and )




        descomponer(list);
    except FormatoValidoException as err:
        print(err.args)



def descomponer(exp):
    try:
        for element in exp:
            if (re.search("[a-z]", element)):
                raise CaracterValidoException("Caracter invalido")
            elif (re.search("[0-9]+",element)):
                newN =  Num(element)
                listaNum.append(newN)
            elif(re.search("[/*+-]",element)):
                newOp = Op(element)
                listaOp.append(newOp)
            else:
                raise CaracterValidoException("Caracter invalido")

    except CaracterValidoException as err:
        print(err.args)
    
    print(listaNum)
    print(listaOp)

        






