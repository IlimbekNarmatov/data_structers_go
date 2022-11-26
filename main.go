package main

import (
    "fmt"
    "data_structures/errors"
)

func main(){
    e := errors.IndexOutOfRange{}
	fmt.Println(e.Error())
}
