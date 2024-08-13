package utils

import (
	"math/rand"
	"reflect"
)


func RandomBool() bool {
	return rand.Uint64()&1 == 0
}

func RandomOneZero() int {
	return rand.Intn(1)
}

func GetType(myvar interface{}) any {
	return reflect.TypeOf(myvar)
}
