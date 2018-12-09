package memory

import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/ability/memory"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
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
func (this Memory) AddVariable(Name string,Value *StructData.Value)StructData.EngineError{
	this.UseLock.Lock()
	if this.Master != nil{
		result := (*Memory)(this.Master).AddVariable(Name,Value)
		this.UseLock.Unlock()
		return result
	}
	_,ok := this.Variables[Name]
	if ok{	//变量已经存在。说明代码有误，必须停止代码执行。
		this.UseLock.Unlock()
		return StructData.MakeError(EngineError.Bad,"Tried to define a existed variable :" + Name)
	}	//如无错误，开始制造变量
	this.Variables[Name].Name = Name
	this.Variables[Name].Value = Value
	this.UseLock.Unlock()
	return StructData.EmptyError
}
func (this Memory) SetVariable(Name string,Value *StructData.Value)StructData.EngineError{
	if this.Master != nil{
		return (*Memory)(this.Master).SetVariable(Name,Value)
	}
	_,ok := this.Variables[Name]
	if !ok{	//变量不存在。说明代码有误，必须停止代码执行。
		return StructData.MakeError(EngineError.Bad,"Tried to use a undefined variable :" + Name)
	}	//如无错误，开始制造变量
	this.Variables[Name].Value = Value
	return StructData.EmptyError
}
func (this Memory) GetType(Name string)int{
	if this.Master != nil{
		return (*memory.Memory)(this.Master).
	}
}