package vmDeal		//本包用来提供一些核心支持功能 与本地系统进行交互
import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

type VmDeal struct {
	vm *StructData.RVM
	Progress *StructData.Progress
	OrderInt struct{
		SetSystem int
	}
}
func (this VmDeal)ChangeProgress(NewProgress *StructData.Progress){
	this.Progress = NewProgress
}
func (this VmDeal)RegisterAll(vm *StructData.RVM,linker *StructData.OrderLinker,Progress *StructData.Progress)StructData.EngineError{
	SetSystem := linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(SetSystem,"SetSystem",this.SetSystem,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	this.OrderInt.SetSystem = SetSystem
	this.Progress = Progress
	this.vm = vm
	return StructData.EmptyError
}
func (this VmDeal)CheckFunction(OrderInt int,argument []string)StructData.EngineError{
	switch OrderInt{
	case this.OrderInt.SetSystem:{	//该命令第一个为整数，第二个任意
		_,err := strconv.Atoi(argument[0])
		if err != nil{
			return StructData.MakeError(EngineError.Bad,err.Error())
		}
		break
	}
	default:{
		return StructData.MakeError(EngineError.Bad,"Undefined Setting of " + strconv.Itoa(OrderInt))
	}
	}
	return StructData.EmptyError
}

const (
	MaxDeepOfStack = 0
)
func (this VmDeal)SetSystem(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	ElementName := Arguments[0]	//设置的属性名称
	Value := Arguments[1]	//属性值
	ElementId,err := strconv.Atoi(ElementName)
	if err != nil{
		return StructData.MakeError(EngineError.Bad,err.Error())
	}
	switch ElementId{
	case MaxDeepOfStack:{	//设置栈的最大深度
		//转换为Int
		v,_ := strconv.Atoi(Value)
		Stack.SetMaxDeep(v)
		break
	}
	}
	return StructData.EmptyError
}