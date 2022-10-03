package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, _ := xjwt.CreateToken(1, 1)
	fmt.Println(token)
}
