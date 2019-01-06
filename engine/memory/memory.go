package memory

import (
	"fmt"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"sync"
)

type Memory struct {
	Engine *StructData.RVM
	Master *StructData.Memory
	Variables map[string]*StructData.Variable
	UseLock sync.Mutex
}
//因为成本限制，禁止程序直接操控内存。
//并且RVM的内存回收直接交给Golang
func (this Memory) Init(Master *StructData.Memory,Engine *StructData.RVM){	//Master为主内存。如果为从内存将依附于主内存 如果允许让瑞雪VM管理全部内存，则可以设置long为0.默认为之后的全部内存
	if Master != nil{	//需要设置主内存
		this.Master = Master
	}
	this.Engine = Engine
	return
}
func New()*StructData.Memory{
	cMemory := &Memory{}
	cMemory.Variables = make(map[string]*StructData.Variable)
	temp := &StructData.Memory{}
	temp.RMemory = cMemory
	temp.SetVariable = cMemory.SetVariable
	temp.AddVariable = cMemory.AddVariable
	temp.GetType = cMemory.GetType
	temp.GetVariable = cMemory.GetVariable

	return temp
}
func (this Memory) AddVariable(Name string,Value *StructData.Value)StructData.EngineError{
	this.UseLock.Lock()
	if this.Master != nil{
		result := this.Master.AddVariable(Name,Value)
		this.UseLock.Unlock()
		return result
	}
	_,ok := this.Variables[Name]
	if ok{	//变量已经存在。说明代码有误，必须停止代码执行。
		this.UseLock.Unlock()
		return StructData.MakeError(EngineError.Bad,"Tried to define a existed variable :" + Name)
	}	//如无错误，开始制造变量
	if this.Variables == nil{
		fmt.Println("errorrrrr")
	}
	this.Variables[Name] = &StructData.Variable{}
	this.Variables[Name].Name = Name
	this.Variables[Name].Value = Value
	this.UseLock.Unlock()
	return StructData.EmptyError
}
func (this Memory) SetVariable(Name string,Value *StructData.Value)StructData.EngineError{
	if this.Master != nil{
		return this.Master.SetVariable(Name,Value)
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
		return this.Master.GetType(Name)
	}
	return this.Variables[Name].Value.Type
}
func (this Memory) GetVariable(Name string)*StructData.Value{
	v,ok := this.Variables[Name]
	if !ok{
		return nil
	}
	return v.Value
}