package main

import (
	"fmt"
	"github.com/Dorbmon/RVm/engine"
)

func main(){
	rvm := engine.RVM{}
	rvm.Init()
	ok,Id := rvm.CreateProgress("Ruixue")
	if !ok{
		fmt.Println("Error when create process")
	}
}
