/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio2;

import java.util.ArrayList;

/**
 *
 * @author AndrikSJ
 */
public class Socio {
    
    int numero_socio;
    String nombre;
    String direccion;
    ArrayList<Libro> libros_prestados;

    public Socio(int numero_socio, String nombre, String direccion) {
        this.numero_socio = numero_socio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.libros_prestados = new ArrayList<Libro>();
    }

    public int getNumero_socio() {
        return numero_socio;
    }

    public String getNombre() {
        return nombre;
    }

    public String getDireccion() {
        return direccion;
    }

    public ArrayList<Libro> getLibros_prestados() {
        return libros_prestados;
    }

    public void setLibros_prestados(ArrayList<Libro> libros_prestados) {
        this.libros_prestados = libros_prestados;
    }

    @Override
    public String toString() {
        return "Socio{" + "numero_socio=" + numero_socio + ", nombre=" + nombre + ", direccion=" + direccion + ",\n libros_prestados=" + libros_prestados + "\n}";
    }
    
    
}
