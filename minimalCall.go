/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: ./cfiles/libcint.a
#include <stdlib.h>
#include <math.h>
#include "./cfiles/cint_funcs.h"
#include "./cfiles/cint.h"
#cgo LDFLAGS: -lm
*/
import "C"

import (
        "fmt"
        "math"
        _"unsafe"
)

func Factorial(n uint64)(result uint64) {
        if (n > 0) {
                result = n * Factorial(n-1)
                return result
        }
        return 1
}

func GtoNorm (n uint64, a float64) float64{
        res := math.Pow(2, (2*float64(n)+3))*float64(Factorial(n+1))*math.Pow((2*a),(float64(n)+1.5))/(float64(Factorial(2*n+2))*math.Sqrt(math.Pi))
        return math.Sqrt(res)
}

func main(){
        cn := C.int(1)

        cexp := C.double(0.5)

        res := C.CINTgto_norm(cn, cexp)
        
        fmt.Println("Parameters (n, exp):", int(cn), float64(cexp))
        fmt.Println("Libcint result:", float64(res))
        fmt.Println("Go result:", GtoNorm(1, 0.5))
}
