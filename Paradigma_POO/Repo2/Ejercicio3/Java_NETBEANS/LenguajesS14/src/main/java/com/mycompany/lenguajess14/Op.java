/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.lenguajess14;

/**
 *
 * @author AndrikSJ
 */
public class Op {
    
    String operador;
            
    public Op(String op){
        this.operador = op;
    }      

    public String getOperador() {
        return operador;
    }

    @Override
    public String toString() {
        return  operador;
    }
    
}
