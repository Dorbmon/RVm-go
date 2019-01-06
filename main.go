package main

import (
	"bufio"
	"fmt"
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/struct"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
	_"net/http/pprof"
)

func main(){
	go func() {
		//port is you coustom define.
		log.Println(http.ListenAndServe("localhost:7000", nil))
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())

	//some func or operation

	rvm := engine.RVM{}
	rvm.Init()
	ok,Id := rvm.CreateProgress("Ruixue")
	if !ok{
		fmt.Println("Error when create process")
	}
	fmt.Println("Id:",Id)
	var c StructData.Code
	/*
	for n := 0;n < 4;n ++{
		tmp := make([]string,5)
		c.Lines = append(c.Lines,tmp)
	}

	c.Lines[0][0] = "newvar"
	c.Lines[0][1] = "count"
	c.Lines[1][0] = "add"
	c.Lines[1][1] = "count"
	c.Lines[1][2] = "1"
	c.Lines[2][0] = "sign"
	c.Lines[2][1] = "fuck"
	c.Lines[3][0] = "goto"
	c.Lines[3][1] = "fuck"
	*/
	c = ReadCode()

	rvm.LoadUncompiledCode(Id,0,c)
	start := time.Now()
	err := rvm.RunCode(Id)
	if StructData.CheckError(err){
		fmt.Println(err)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}
func ReadCode()StructData.Code{
	fi, err := os.Open("test.rvm")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}
	rtCode := StructData.Code{}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		strLine := string(a)
		tmp := strings.Split(strLine," ")
		rtCode.Lines = append(rtCode.Lines,tmp)
		//rtCode.Lines[0] = tmp
	}
	return rtCode
}