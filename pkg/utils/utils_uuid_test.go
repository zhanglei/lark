package utils

import (
	"fmt"
	"testing"
)

func GetGetUUID(t *testing.T) {
	id := NewUUID()
	fmt.Println(id)
}
