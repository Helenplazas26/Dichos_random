package main

import (
    "fmt"
    "math/rand/v2"
)

func GetRandomNumber() int {
    return rand.IntN(100) // devuelve un número entre 0 y 99
}

func main() {
    fmt.Println(GetRandomNumber())
}
