/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Project/Maven2/JavaApp/src/main/java/${packagePath}/${mainClassName}.java to edit this template
 */

package com.mycompany.ejercicio1;

/**
 *
 * @author AndrikSJ
 */
public class Ejercicio1 {

    public static void main(String[] args) {


        Rectangulo rectangulotemp = new Rectangulo("Rectangulo 1");
        Circulo circulotemp = new Circulo("Circulo 1", 20);

        Cuadrado cuadradotemp = new Cuadrado("Cuadrado 1", 10);

        Texto textotemp1 = new Texto("Hola mundo");
        Texto textotemp2 = new Texto("Comida tailandesa");


        Grupo grupotemp = new Grupo("grupo1");

        grupotemp.addFigura(circulotemp);


        Hoja hojatemp = new Hoja("Hoja de matematica");

        hojatemp.addFigura(rectangulotemp);
        hojatemp.addGrupo(grupotemp);
        hojatemp.addTexto(textotemp1);
        hojatemp.addFigura(cuadradotemp);
        hojatemp.addTexto(textotemp2);


        System.out.println("Hoja creada" + hojatemp);



    }
}
