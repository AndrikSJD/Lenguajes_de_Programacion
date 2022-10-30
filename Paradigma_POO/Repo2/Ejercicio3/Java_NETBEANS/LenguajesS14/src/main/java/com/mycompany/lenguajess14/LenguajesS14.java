/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Project/Maven2/JavaApp/src/main/java/${packagePath}/${mainClassName}.java to edit this template
 */
package com.mycompany.lenguajess14;

import static java.lang.System.in;
import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

/**
 *
 * @author AndrikSJ
 */
public class LenguajesS14 {

    static ArrayList<Num> numerosL = new ArrayList<Num>();
    static ArrayList<Op> operadoresL = new ArrayList<Op>();

    public static void main(String[] args) {

        leerExp();

    }

    public static void leerExp() {
        String expresion;
        Scanner leer = new Scanner(System.in);
        System.out.print("Introduzca la expresion: ");
        expresion = leer.nextLine();
        String data = expresion.replaceAll("\\s", ""); //quitar espacios en blanco, reemplaza " " por ""
        String[] data2 = data.split("(?=[-+*/])|(?<=[-+*/])"); //expresion regular para hacer el split antes o despues del simbolo matematico
        List<String> list = Arrays.asList(data2); //lista con numeros y operadores

        descomponer(list);

    }

    public static void descomponer(List<String> exp) {

        try {

            for (int i = 0; i < exp.size(); i++) {
                String val = exp.get(i);
                if (i % 2 == 0 && val.matches("[0-9]+") == false) {
                    throw new FormatoValidoException("Formato expresion infija invalido");

                } else {
                    if (val.matches("[a-z]")) {
                        throw new CaracterValidoException("Caracter invalido");
                    } else if (val.matches("[0-9]+")) {

                        Num temp = new Num(val);
                        numerosL.add(temp);
                    } else if (val.matches("[/*+-]")) {

                        Op temp = new Op(val);
                        operadoresL.add(temp);
                    } else {
                        throw new CaracterValidoException("Caracter invalido");
                    }
                }
            }

        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();

        }
        realizarOP();

    }

    public static void realizarOP() {

        try {

            if (numerosL.size() < operadoresL.size()) {
                throw new FormatoValidoException("Formato de expresion regular invalido");

            } else {

                while (numerosL.size() > 1) {
                    int temp1 = numerosL.get(0).getNumero();
                    numerosL.remove(0);
                    String tempOp = operadoresL.get(0).getOperador();
                    operadoresL.remove(0);
                    int temp2 = numerosL.get(0).getNumero();
                    numerosL.remove(0);

                    if (tempOp.equals("+")) {
                        temp1 += temp2;

                    } else if (tempOp.equals("-")) {
                        temp1 -= temp2;

                    } else if (tempOp.equals("*")) {
                        temp1 *= temp2;

                    } else {
                        temp1 /= temp2;

                    }
                    Num result = new Num(String.valueOf(temp1));
                    numerosL.add(0, result);
                }
                System.out.print("El resultado de la operación es: ");
                System.out.println(numerosL.get(0));

            }

        } catch (FormatoValidoException e) {
            System.out.println(e.getMessage());
            e.printStackTrace();

        }

    }

}
