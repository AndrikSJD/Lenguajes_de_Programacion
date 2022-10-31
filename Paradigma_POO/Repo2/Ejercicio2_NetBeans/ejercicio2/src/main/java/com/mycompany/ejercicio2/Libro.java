/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio2;

/**
 *
 * @author AndrikSJ
 */
public class Libro {
    
    int codigo;
    String titulo;
    String autor;
    boolean disponibilidad;
    int socio_prestado;

    public Libro(int codigo, String titulo, String autor, boolean disponibilidad) {
        this.codigo = codigo;
        this.titulo = titulo;
        this.autor = autor;
        this.disponibilidad = disponibilidad;
    }

    public void setSocio_prestado(int socio_prestado) {
        this.socio_prestado = socio_prestado;
    }

    public int getCodigo() {
        return codigo;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getAutor() {
        return autor;
    }

    public boolean isDisponibilidad() {
        return disponibilidad;
    }

    public int getSocio_prestado() {
        return socio_prestado;
    }

    public void setDisponibilidad(boolean disponibilidad) {
        this.disponibilidad = disponibilidad;
    }

    @Override
    public String toString() {
        return "Libro{" + "codigo=" + codigo + ", titulo=" + titulo + ", autor=" + autor + ", disponibilidad=" + disponibilidad + ", socio_prestado=" + socio_prestado + '}';
    }
    
            
    
}
