/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public class Cuadrado extends Figura {

    int lado;
    
    public Cuadrado(String nombre, int lado) {

        super(nombre);
        this.lado = lado;
    }

    @Override
    public int getArea() {
        return lado * lado;
    }

    @Override
    public String toString() {
        return "Cuadrado{" +
                "nombre='" + nombre + '\'' +
                ", vertices=" + vertices +
                ", largo=" + largo +
                ", ancho=" + ancho +
                ", altura=" + altura +
                ", perimetro=" + perimetro +
                '}';
    }
}
