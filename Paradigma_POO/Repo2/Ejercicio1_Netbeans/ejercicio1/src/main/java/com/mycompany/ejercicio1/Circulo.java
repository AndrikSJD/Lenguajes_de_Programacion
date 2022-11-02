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
    
    int area;
    int numero_cuerdas;
    

    public Circulo(int area, int numero_cuerdas, String nombre, int vertices, int perimetro) {
        super(nombre, vertices, perimetro);
        this.area = area;
        this.numero_cuerdas = numero_cuerdas;
    }
    
}
