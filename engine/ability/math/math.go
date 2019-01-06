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
		Comparevc int
		Comparevcless int
	}
}
func (this Math)RegisterAll(vm *StructData.RVM,linker *StructData.OrderLinker,Process *StructData.Progress)StructData.EngineError{
	this.Process = Process
	this.Vm = vm
	this.orderInt.Add = linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(this.orderInt.Add,"add",this.Add,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	this.orderInt.Comparevc = linker.GetAnRandomOrderInt()
	err = linker.RegisterOrder(this.orderInt.Comparevc,"comparevc",this.Comparevc,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	this.orderInt.Comparevcless = linker.GetAnRandomOrderInt()
	err = linker.RegisterOrder(this.orderInt.Comparevcless,"comparevcless",this.Comparevcless,this.CheckFunction)
	if StructData.CheckError(err){
		return err
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
		break
	}
	case this.orderInt.Comparevc:{
		if len(Arguments) != 2{
			return StructData.MakeError(EngineError.Bad,"comparevc need double arguments for the name of variable and the const number")
		}
		//查看是否为数字
		_,err := strconv.Atoi(Arguments[1])
		if err != nil{
			return StructData.MakeError(EngineError.Bad,"the NO.2 argument of comparevc must be int")
		}
	}
	case this.orderInt.Comparevcless:{
		if len(Arguments) != 2{
			return StructData.MakeError(EngineError.Bad,"comparevcless need double arguments for the name of variable and the const number")
		}
		//查看是否为数字
		_,err := strconv.Atoi(Arguments[1])
		if err != nil{
			return StructData.MakeError(EngineError.Bad,"the NO.2 argument of comparevcless must be int")
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
	var oldint int
	if oldData.Type == TypeSystem.String{
		oldint,_ = strconv.Atoi(oldData.Value.(string))
	}else{
		oldint = oldData.Value.(int)
	}

	oldData.Value = oldint + Addn
	oldData.Type = TypeSystem.Int
	return this.Process.Memory.SetVariable(VariableName,oldData)
}
func (this Math)Comparevc(Arguments []string,Stack *StructData.Stack)StructData.EngineError{	//compare variable and const
	//获取变量
	VariableName := Arguments[0]
	Var := this.Process.Memory.GetVariable(Arguments[0])
	if Var.Value == nil{
		Var.Value = 0
		Var.Type = TypeSystem.Int
	}else{	//如果变量不为空，则进行类型判断
		if Var.Type != TypeSystem.Int && Var.Type != TypeSystem.Int64 && Var.Type != TypeSystem.UInt && Var.Type != TypeSystem.UInt64{
			return StructData.MakeError(EngineError.Bad,"Bad type of variable:" + VariableName)
		}
	}
	//开始比较
	nConstNumber,_ := strconv.Atoi(Arguments[1])
	vValue := Var.Value.(int)
	var err StructData.EngineError
	if vValue > nConstNumber{
		err = Stack.Push(1,TypeSystem.Int)
	}else{
		err = Stack.Push(0,TypeSystem.Int)
	}
	return err
}
func (this Math)Comparevcless(Arguments []string,Stack *StructData.Stack)StructData.EngineError{	//compare variable and const
	//获取变量
	VariableName := Arguments[0]
	Var := this.Process.Memory.GetVariable(Arguments[0])
	if Var.Value == nil{
		Var.Value = 0
		Var.Type = TypeSystem.Int
	}else{	//如果变量不为空，则进行类型判断
		if Var.Type != TypeSystem.Int && Var.Type != TypeSystem.Int64 && Var.Type != TypeSystem.UInt && Var.Type != TypeSystem.UInt64{
			return StructData.MakeError(EngineError.Bad,"Bad type of variable:" + VariableName)
		}
	}
	//开始比较
	nConstNumber,_ := strconv.Atoi(Arguments[1])
	vValue := Var.Value.(int)
	var err StructData.EngineError
	if vValue < nConstNumber{
		err = Stack.Push(1,TypeSystem.Int)
	}else{
		err = Stack.Push(0,TypeSystem.Int)
	}
	return err
}