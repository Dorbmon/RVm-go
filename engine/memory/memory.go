package memory

import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"time"
)

type Memory StructData.Memory
//因为成本限制，禁止程序直接操控内存。
//并且RVM的内存回收直接交给Golang
func (this Memory) Init(Master *StructData.Memory,Engine *engine.RVM){	//Master为主内存。如果为从内存将依附于主内存 如果允许让瑞雪VM管理全部内存，则可以设置long为0.默认为之后的全部内存
	if Master != nil{	//需要设置主内存
		this.Master = Master
	}
	this.Engine = Engine
	return
}
func (this Memory) AddVariable(Name string,Value *StructData.Value){
	this.UseLock.Lock()
	if this.Master != nil{
		(*Memory)(this.Master).AddVariable(Name,Value)
		this.UseLock.Unlock()
		return
	}
	_,ok := this.Variables[Name]
	if ok{	//变量已经存在。说明代码有误，必须停止代码执行。
		this.UseLock.Unlock()
		NewError := StructData.EngineError{}
		NewError.Class = EngineError.Bad
		NewError.Time = time.Now().String()
		NewError.Data = "Tried to define a existed variable :" + Name

		this.Engine.ThrowError(NewError)
		return
	}	//如无错误，开始制造变量
	this.Variables[Name].Name = Name
	this.Variables[Name].Value = Value
	this.UseLock.Unlock()
	return
}
func (this Memory) SetVariable(Name string,Value *StructData.Value){
	if this.Master != nil{
		(*Memory)(this.Master).SetVariable(Name,Value)
		return
	}
	_,ok := this.Variables[Name]
	if !ok{	//变量不存在。说明代码有误，必须停止代码执行。
		this.UseLock.Unlock()
		NewError := StructData.EngineError{}
		NewError.Class = EngineError.Bad
		NewError.Time = time.Now().String()
		NewError.Data = "Tried to use a undefined variable :" + Name
		this.Engine.ThrowError(NewError)
		return
	}	//如无错误，开始制造变量
	this.Variables[Name].Value = Value
	return
}