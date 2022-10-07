package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, _ := xjwt.CreateToken(1578215274794979328, 1)
	fmt.Println(token)
}
