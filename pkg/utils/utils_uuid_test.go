package utils

import (
	"fmt"
	"testing"
)

func TestGetGetUUID(t *testing.T) {
	id := NewUUID()
	fmt.Println(id)
}
