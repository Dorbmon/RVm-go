package memory

import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/ability"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/memory"
	"github.com/Dorbmon/RVm/engine/orderLinker"
	"github.com/Dorbmon/RVm/engine/stack"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
)

type Memory struct{
	vm *engine.RVM
	Progress *StructData.Progress
	orderInt struct{
		NewVar int
		SetVar int
	}
}

func (this Memory)RegisterAll(vm *engine.RVM,linker *orderLinker.OrderLinker,Progress *StructData.Progress)StructData.EngineError{
	NewVar := linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(NewVar,"newvar",this.NewVar,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	SetVar := linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(SetVar,"setvar",this.SetVar,this.CheckFunction)
	this.Progress = Progress
	this.vm = vm
	return StructData.EmptyError
}
func (this Memory)NewVar(Arguments []string,Stack *stack.Stack)StructData.EngineError{
	varName := Arguments[0]
	Memory := (memory.Memory)(this.Progress.Memory)
	Value := &StructData.Value{}
	Value.Type = TypeSystem.Unknown
	return Memory.AddVariable(varName,&StructData.Value{})
}
func (this Memory)CheckFunction(OrderInt int,Arguments []string)StructData.EngineError{

	switch OrderInt{
	case this.orderInt.NewVar:{
		if len(Arguments) != 2{
			return StructData.MakeError(EngineError.Bad,"NewVar only need one argument for variable name and type")
		}

	}
	case this.orderInt.NewVar:{}
		if len(Arguments) < 2{
			return StructData.MakeError(EngineError.Bad,"NewVar  need double argument for variable name and value")
		}
	}
	return StructData.EmptyError
}
func (this Memory)SetVar(Arguments []string,Stack *stack.Stack)StructData.EngineError{
	varName := Arguments[0]
	return (memory.Memory)(this.Progress.Memory).SetVariable(varName,ability.AddString(Arguments[1:]))
}