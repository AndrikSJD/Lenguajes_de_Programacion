/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public class Circulo extends Figura {
    
    int radio;
    double pi;
    

    public Circulo(String nombre, int radio) {
        super(nombre);
        this.radio = radio;
        this.pi = 3.14;
    }

    @Override
    public int getArea() {
        return 0;
    }

    @Override
    public String toString() {
        return "Circulo{" +
                "radio=" + radio +
                ", pi=" + pi +
                ", nombre='" + nombre + '\'' +
                ", vertices=" + vertices +
                ", largo=" + largo +
                ", ancho=" + ancho +
                ", altura=" + altura +
                ", perimetro=" + perimetro +
                '}';
    }
}
