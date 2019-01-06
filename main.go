package main

import (
	"fmt"
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/struct"
)

func main(){
	rvm := engine.RVM{}
	rvm.Init()
	ok,Id := rvm.CreateProgress("Ruixue")
	if !ok{
		fmt.Println("Error when create process")
	}
	fmt.Println("Id:",Id)
	var c StructData.Code
	for n := 0;n < 2;n ++{
		tmp := make([]string,5)
		c.Lines = append(c.Lines,tmp)
	}
	c.Lines[0][0] = "newvar"
	c.Lines[0][1] = "count"
	c.Lines[1][0] = "add"
	c.Lines[1][1] = "count"
	c.Lines[1][2] = "1"
	rvm.LoadUncompiledCode(Id,0,c)
	err := rvm.RunCode(Id)
	if StructData.CheckError(err){
		fmt.Println(err)
	}
}
