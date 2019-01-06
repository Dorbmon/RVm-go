package memory

import (
	"bytes"
	"github.com/Dorbmon/RVm/engine/error"
	//"github.com/Dorbmon/RVm/engine/memory"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
)

type Memory struct{
	vm *StructData.RVM
	Progress *StructData.Progress
	orderInt struct{
		NewVar int
		SetVar int
	}
}

func (this *Memory)RegisterAll(vm *StructData.RVM,linker *StructData.OrderLinker,Progress *StructData.Progress)StructData.EngineError{
	this.Progress = Progress
	this.vm = vm
	NewVar := linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(NewVar,"newvar",this.NewVar,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	SetVar := linker.GetAnRandomOrderInt()
	err = linker.RegisterOrder(SetVar,"setvar",this.SetVar,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}

	return StructData.EmptyError
}
func (this *Memory)NewVar(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	varName := Arguments[0]
	Value := &StructData.Value{}
	Value.Type = TypeSystem.Unknown
	return this.Progress.Memory.AddVariable(varName,&StructData.Value{})
}
func (this *Memory)CheckFunction(OrderInt int,Arguments []string)StructData.EngineError{

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
func (this *Memory)SetVar(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	varName := Arguments[0]
	value := &StructData.Value{}
	value.Value = AddString(Arguments[1:])
	value.Type = this.Progress.Memory.GetType(varName)
	return this.Progress.Memory.SetVariable(varName,value)

}
func AddString(Data []string)string{	//拼接文本
	var buffer bytes.Buffer
	for i := 0; i < len(Data); i ++{
		buffer.WriteString(Data[i])
	}
	return buffer.String()
}