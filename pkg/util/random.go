package util

import (
    "math/rand/v2"
	"dichos.com/pkg/repository"
)

func GetRandomNumber() int {
	
    return rand.IntN(repository.TamanioLista()) // devuelve un número entre 0 y 99
}

