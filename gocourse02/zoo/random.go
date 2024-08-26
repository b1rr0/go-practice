package zoo

import (
	"math/rand/v2"
	"reflect"
)

func RandomBool() bool {
	return rand.Uint64()&1 == 0
}

func RandomOneZero() int {
	return rand.Int() % 2
}

func GetType(myvar interface{}) any {
	return reflect.TypeOf(myvar)
}
