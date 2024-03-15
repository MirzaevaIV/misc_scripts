package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: ./cfiles/libcint.a
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include "./cfiles/cint_funcs.h"
#include "./cfiles/cint.h"
#cgo LDFLAGS: -lm  -lquadmath
int cint1e_ipnuc_cart(double *buf, int *shls, int *atm, int natm, int *bas, int nbas, double *env);
*/
import "C"

import (
        "fmt"
)

/* general contracted DZ basis [3s1p/2s1p] for H2
    exponents    contract-coeff
S   6.0          0.7               0.4
    2.0          0.6               0.3
    0.8          0.5               0.2
P   0.9          1.
*/


func CalcContractionCoeffs(l int, c []float64) []float64{
        res := make([]float64, len(c))
        for i := range res{
                res[i] = float64(C.CINTgto_norm(C.int(l), C.double(c[i])))
        }
        return res
}

func CConvertInt(a [][]int) [][]C.int{
        res := make([][]C.int, len(a))
        for i:= range res{
                res[i] = make([]C.int, len(a[i]))
        }
        for i := range a{
                for j := range a[i]{
                        res[i][j] = C.int(a[i][j])
                }
        }
        return res
}

func CConvertDouble(a []float64) []C.double{
        res := make([]C.double, len(a))
        for i := range res{
                res[i] = C.double(a[i])
        }
        return res
}

func main(){
        var atm [][]int
        var bas [][]int
        env := make([]float64,20)
        env_coord := 20 //index of current position in the env

        //add atom 1
        atm = append(atm, []int{1, env_coord, 0, 0, 0, 0})
        env = append(env, 0.0, 0.0, 0.8)
        env_coord += 3
        //add atom 2
        atm = append(atm, []int{1, env_coord, 0, 0, 0, 0})
        env = append(env, 0.0, 0.0, -0.8)
        env_coord += 3

        //basis for atom 1
        // 3s -> 2s (3 primitives, 2 contraction coefficients sets)
        env = append(env, 6.0, 2.0, 0.8) //exponents
        contr_coef := CalcContractionCoeffs(0, []float64{ 6.0, 2.0, 0.8})
        env = append(env, 0.7*contr_coef[0], 0.6*contr_coef[1], 0.5*contr_coef[2]) //1st set of contraction coefficients
        env = append(env, 0.4*contr_coef[0], 0.3*contr_coef[1], 0.2*contr_coef[2]) //2nd set of contraction coefficients
        bas = append(bas, []int{0, 0, 3, 2, 0, env_coord, env_coord+3, 0}) //    ATOM_OF, ANG_OF, NPRIM_OF, NCTR_OF, KAPPA_OF, PTR_EXP, PTR_COEFF
        env_coord += 9

        // 1p -> 1p (1 primitive, 1 contraction coefficients set)
        env = append(env, 0.9) //exponents
        contr_coef = CalcContractionCoeffs(1, []float64{0.9})
        env = append(env, 1.0*contr_coef[0]) //contraction coefficients
        bas = append(bas, []int{0, 1, 1, 1, 0, env_coord, env_coord+1, 0})
        env_coord += 2

        // basis for atom 2 (same as for atom 1)
        bas = append(bas, bas[0], bas[1])

        fmt.Println("atm:", atm)
        fmt.Println("bas:", bas)
        fmt.Println("env:", env)
        fmt.Println("Current env coordinate:", env_coord)

        // convert to C type array
        cbas := CConvertInt(bas)
        catm := CConvertInt(atm)
        cenv := CConvertDouble(env)

        di := C.CINTcgto_cart(C.int(0), &cbas[0][0])
        dj := C.CINTcgto_cart(C.int(1), &cbas[0][0])
        fmt.Println("di, dj:", int(di), int(dj))


        var shls [2]C.int
        shls[0] = C.int(0)
        shls[1] = C.int(1)

        buf := make([]C.double, int(di*dj*3))

        natm := C.int(len(atm))
        nbas := C.int(len(bas))

        fmt.Println("Number of atoms:", int(natm),"\nNumber of basis functions:",  int(nbas))
        _ = C.cint1e_ipnuc_cart(&buf[0], &shls[0], &catm[0][0], natm, &cbas[0][0], nbas, &cenv[0])

        fmt.Println("Result:")
        for i := 0; i< int(di*dj*3); i++{
                fmt.Println (float64(buf[i]))
        }
}
