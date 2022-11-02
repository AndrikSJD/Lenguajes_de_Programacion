/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public class Elipse extends Figura {
    
    public Elipse(String nombre) {

        super(nombre);
    }

    @Override
    public int getArea() {
        return 0;
    }

    @Override
    public String toString() {
        return "Elipse{" +
                "nombre='" + nombre + '\'' +
                ", vertices=" + vertices +
                ", largo=" + largo +
                ", ancho=" + ancho +
                ", altura=" + altura +
                ", perimetro=" + perimetro +
                '}';
    }
}
