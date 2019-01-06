package math

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)
type Math struct{
	Vm *StructData.RVM
	Process *StructData.Progress
	orderInt struct {
		Add int
	}
}
func (this Math)RegisterAll(vm *StructData.RVM,linker *StructData.OrderLinker,Process *StructData.Progress)StructData.EngineError{
	this.Process = Process
	this.Vm = vm
	this.orderInt.Add = linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(this.orderInt.Add,"add",this.Add,this.CheckFunction)
	if StructData.CheckError(err){
	}
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
func (this Math)Add(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	VariableName := Arguments[0]
	oldData := this.Process.Memory.GetVariable(Arguments[0])
	if oldData.Value == nil{
		oldData.Value = ""
	}
	Addn,_ := strconv.Atoi(Arguments[1])
	oldint,_ := strconv.Atoi(oldData.Value.(string))
	oldData.Value = oldint + Addn
	oldData.Type = TypeSystem.Int
	return this.Process.Memory.SetVariable(VariableName,oldData)
}
