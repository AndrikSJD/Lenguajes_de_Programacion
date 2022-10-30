/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.lenguajess14;

/**
 *
 * @author AndrikSJ
 */
public class Num {
    
    int numero;

    public Num(String n) {
        this.numero = Integer.parseInt(n);
    }

    public int getNumero() {
        return numero;
    }

    @Override
    public String toString() {
        return String.valueOf(numero);
    }
    
    
    
    
    
}
