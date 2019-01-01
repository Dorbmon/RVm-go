package math

import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/memory"
	"github.com/Dorbmon/RVm/engine/orderLinker"
	"github.com/Dorbmon/RVm/engine/stack"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)
type Math struct{
	Vm *engine.RVM
	Process *StructData.Progress
	orderInt struct {
		Add int
	}
}
func (this Math)RegisterAll(vm *engine.RVM,linker *orderLinker.OrderLinker,Process *StructData.Progress)StructData.EngineError{
	this.orderInt.Add = linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(this.orderInt.Add,"newvar",this.Add,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	this.Process = Process
	this.Vm = vm
	return StructData.EmptyError
}
func (this Math)CheckFunction(OrderInt int,Arguments []string)StructData.EngineError{
	switch OrderInt{
	case this.orderInt.Add:{
		if len(Arguments) != 2{
			return StructData.MakeError(EngineError.Bad,"Add need double for adding.Variable , Adding Number;")
		}
		_,err2 := strconv.Atoi(Arguments[1])
		if err2 != nil{
			return StructData.MakeError(EngineError.Bad,"Error Type Of Data For Add.")
		}
	}
	}
	return StructData.EmptyError
}
func (this Math)Add(Arguments []string,Stack *stack.Stack)StructData.EngineError{
	VariableName := Arguments[0]
	oldData := this.Process.Memory.Variables[Arguments[0]].Value
	Addn,_ := strconv.Atoi(Arguments[0])
	oldData.Value = oldData.Value.(int) + Addn
	return (memory.Memory)(this.Process.Memory).SetVariable(VariableName,oldData)
}
