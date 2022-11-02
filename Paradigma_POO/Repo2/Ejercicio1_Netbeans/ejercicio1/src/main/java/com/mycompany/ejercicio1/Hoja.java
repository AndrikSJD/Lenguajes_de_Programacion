package com.mycompany.ejercicio1;

import java.util.ArrayList;

public class Hoja {


    private String nombre;
    private ArrayList<Texto> textos_list;
    private ArrayList<Grupo> grupos_list;
    private ArrayList<Figura> figuras_list;

    public Hoja(String nombre) {
        this.nombre = nombre;
        this.grupos_list = new ArrayList<>();
        this.textos_list = new ArrayList<>();
        this.figuras_list = new ArrayList<>();

    }

    public ArrayList<Figura> getFiguras_list() {
        return figuras_list;
    }

    public ArrayList<Texto> getTextos_list() {
        return textos_list;
    }

    public ArrayList<Grupo> getGrupos_list() {
        return grupos_list;
    }

    public void addFigura(Figura figura) {
        figuras_list.add(figura);
    }

    public void addTexto(Texto texto) {
        textos_list.add(texto);
    }

    public void addGrupo(Grupo grupo) {
        grupos_list.add(grupo);
    }

    @Override
    public String toString() {
        return "Hoja {\n" +
                "nombre='" + nombre + '\'' +
                ",\n textos_list=" + textos_list +
                ",\n grupos_list=" + grupos_list +
                ",\n figuras_list=" + figuras_list + "\n" +
                '}';
    }
}
