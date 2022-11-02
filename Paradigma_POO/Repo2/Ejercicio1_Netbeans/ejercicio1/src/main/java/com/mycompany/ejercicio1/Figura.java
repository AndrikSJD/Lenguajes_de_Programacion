/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public abstract class Figura {

    String nombre;
    int vertices;
    int largo;
    int ancho;
    int altura;
    int perimetro;

    public Figura(String nombre ) {
        this.nombre = nombre;
        this.vertices = 0;
        this.perimetro = 0;
        this.altura = 0;
        this.ancho = 0;
        this.largo = 0;

    }

    public abstract int getArea();

    public abstract String toString();
}
