package main

import (
	"fmt"
	"./frac"
)

func main() {
	f,err:=frac.New(15,5)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
