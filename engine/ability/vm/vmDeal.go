package vmDeal		//本包用来提供一些核心支持功能
import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/orderLinker"
	"github.com/Dorbmon/RVm/engine/stack"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

type VmDeal struct {
	vm *engine.RVM
	Progress *StructData.Progress
}
func (this VmDeal)ChangeProgress(NewProgress *StructData.Progress){
	this.Progress = NewProgress
}
func (this VmDeal)RegisterAll(vm *engine.RVM,linker *orderLinker.OrderLinker,Progress *StructData.Progress)StructData.EngineError{
	SetSystem := linker.GetAnRandomOrderInt()
	err := linker.RegisterOrder(SetSystem,"SetSystem",2,this.SetSystem)
	if StructData.CheckError(err){
		return err
	}

	this.Progress = Progress
	this.vm = vm
	return StructData.EmptyError
}

const (
	MaxDeepOfStack = 0
)
func (this VmDeal)SetSystem(Arguments []string,Stack *stack.Stack)StructData.EngineError{
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