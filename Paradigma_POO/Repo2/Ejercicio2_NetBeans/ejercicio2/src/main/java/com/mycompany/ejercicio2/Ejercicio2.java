/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Project/Maven2/JavaApp/src/main/java/${packagePath}/${mainClassName}.java to edit this template
 */

package com.mycompany.ejercicio2;

import com.mycompany.ejercicio2.Logica;

/**
 *
 * @author AndrikSJ
 */
public class Ejercicio2 {

    public static void main(String[] args) {
        Logica temp = new Logica();
        temp.agregarLibro(234, "Cuentos de terror", "Juan Alcazar", true);
        temp.agregarLibro(555, "Frankenstein", "Mary Shelley", true);
        temp.agregarLibro(777, "La Odisea", "Homero", true);
        temp.agregarLibro(785, "Don Quijote", "Miguel de Cervantes Sabedra", true);
        temp.agregarLibro(444, "El Cuervo", "Edgar Allan Poe", true);
        temp.agregarLibro(888, "Mi amigo el oso", "Alberto Roedra", true);
        temp.agregarLibro(523, "Mi perro Roberto", "Hughes Join", true);
        temp.agregarLibro(812, "La casa encantada", "Koop Smith", true);
        temp.agregarLibro(521, "Cazando fantasmas", "Bu Jian", true);
        temp.agregarLibro(001, "Esperando el momento", "Fabelo Dois", true);
        temp.agregarSocio(9829,"Andrik","300 metros Sur de la Iglesia");
        temp.agregarSocio(1235,"Julia","300 metros Norte del Parque");
        temp.agregarSocio(1479,"Jose","Barrio Escalante, 50 metros del super Ana");
        temp.agregarSocio(2222,"Nicole","Barrio San Isidro");
        temp.prestarLibro(777, 9829, "22-01-2022");
        temp.prestarLibro(888, 9829, "05-10-2022");
        temp.prestarLibro(555, 9829, "19-07-2022");
        temp.prestarLibro(234, 1235, "09-03-2022");
        temp.prestarLibro(785, 1235, "12-09-2022");
        temp.prestarLibro(444, 1235, "10-04-2022");
        temp.prestarLibro(523, 1479, "22-08-2022");
        temp.prestarLibro(521, 1479, "29-11-2022");
        temp.prestarLibro(812, 2222, "27-05-2022");
        temp.prestarLibro(001, 2222, "30-01-2022");
        temp.prestarLibro(523, 2222, "16-07-2022"); //intentar reservar libro ya prestado ERROR
        System.out.println(temp.socios_Mas_Prestados());
        
        
        
    }
}
