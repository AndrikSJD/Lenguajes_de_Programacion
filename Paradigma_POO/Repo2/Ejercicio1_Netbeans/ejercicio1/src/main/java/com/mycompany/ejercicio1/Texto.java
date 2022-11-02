package com.mycompany.ejercicio1;

public class Texto {
    String texto;

    public Texto(String texto) {
        this.texto = texto;
    }

    @Override
    public String toString() {
        return "Texto{" +
                "texto='" + texto + '\'' +
                '}';
    }

    public void setTexto(String texto) {
        this.texto = texto;
    }
}

