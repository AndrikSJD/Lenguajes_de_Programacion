/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.ejercicio2;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

/**
 *
 * @author AndrikSJ
 */
public class Logica {

    ArrayList<Socio> sociosList = new ArrayList<Socio>();
    ArrayList<Libro> librosList = new ArrayList<Libro>();
    ArrayList<Prestamo> prestamosL = new ArrayList<Prestamo>();

    public void agregarSocio(int numero_socio, String nombre, String direccion) {
        Socio tempS = new Socio(numero_socio, nombre, direccion);
        sociosList.add(tempS);
    }

    public void agregarLibro(int codigo, String titulo, String autor, boolean disponibilidad) {
        Libro libroTemp = new Libro(codigo, titulo, autor, disponibilidad);
        librosList.add(libroTemp);
    }

    public boolean reservar_lib(int codeLib, int socioP) {

        for (int i = 0; i < librosList.size(); i++) {
            if (librosList.get(i).getCodigo() == codeLib && librosList.get(i).isDisponibilidad() == true) {
                librosList.get(i).setDisponibilidad(false);
                librosList.get(i).setSocio_prestado(socioP);
                return true;
            }

        }
        return false;
    }

    public Socio buscar_Socio(int numSocio) {
        for (int i = 0; i < sociosList.size(); i++) {
            if (sociosList.get(i).getNumero_socio() == numSocio) {
                return sociosList.get(i);

            }
        }
        return null;
    }

    public Libro buscar_Libro(int codeLib) {
        for (int i = 0; i < librosList.size(); i++) {
            if (librosList.get(i).getCodigo() == codeLib) {
                return librosList.get(i);
            }
        }
        return null;

    }

    public void prestarLibro(int codeLib, int numeroSocio, String fecha) {

        if (buscar_Socio(numeroSocio) != null && buscar_Libro(codeLib) != null) {
            boolean check = reservar_lib(codeLib, numeroSocio);
            if (check) {
                Libro libroP = buscar_Libro(codeLib);
                Socio socioP = buscar_Socio(numeroSocio);
                ArrayList<Libro> listaLSocio = socioP.getLibros_prestados();
                listaLSocio.add(libroP);
                socioP.setLibros_prestados(listaLSocio);
                Prestamo nuevoPrestamo = new Prestamo(codeLib, numeroSocio, fecha);
                prestamosL.add(nuevoPrestamo);
                reservar_lib(codeLib, numeroSocio);
                System.out.println("Libro prestado correctamente al socio " + socioP.getNombre());
            } else {

                System.out.println("El libro solicitado ya se encuentra reservado, intente con otro");

            }

        } else {

            System.out.println("El libro para prestar o el socio ingresado no existe. \nPor favor ingreselos e intente denuveo");
        }

    }

    public List<Socio> socios_Mas_Prestados() {

        List<Socio> masP = sociosList.stream()
                .filter(socio -> socio.getLibros_prestados().size() >= 3)
                .collect(Collectors.toList());

        return masP;
    }

}
