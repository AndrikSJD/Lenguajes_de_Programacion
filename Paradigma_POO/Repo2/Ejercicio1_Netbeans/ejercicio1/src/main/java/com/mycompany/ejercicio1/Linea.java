/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public class Linea extends Figura {
    
    public Linea(String nombre) {

        super(nombre);
    }

    @Override
    public int getArea() {
        return 0;
    }

    @Override
    public String toString() {
        return "Linea{" +
                "nombre='" + nombre + '\'' +
                ", vertices=" + vertices +
                ", largo=" + largo +
                ", ancho=" + ancho +
                ", altura=" + altura +
                ", perimetro=" + perimetro +
                '}';
    }
}
