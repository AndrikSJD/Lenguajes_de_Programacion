package com.mycompany.ejercicio1;

import java.util.ArrayList;

public class Grupo {

    private String nombre;
    private ArrayList<Figura> figuras_list;

    public Grupo(String nombre) {
        this.nombre = nombre;
        this.figuras_list = new ArrayList<>();
    }

    public ArrayList<Figura> getFiguras_list() {
        return figuras_list;
    }

    public void setFiguras_list(ArrayList<Figura> figuras_list) {
        this.figuras_list = figuras_list;
    }

    public void addFigura(Figura fig) {

        this.figuras_list.add(fig);


    }



    @Override
    public String toString() {
        return "Grupo{" +
                "nombre='" + nombre + '\'' +
                ", figuras_list=" + figuras_list +
                '}';
    }
}
