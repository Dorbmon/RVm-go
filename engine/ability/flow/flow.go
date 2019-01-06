package flow

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
)

type Flow struct {
	vm *StructData.RVM
	process *StructData.Progress
	orderInt struct{
		Goto int
		Sign int
	}
}

func (this *Flow)RegisterAll(vm *StructData.RVM,Process *StructData.Progress)StructData.EngineError{
	this.vm = vm
	this.process = Process
	Goto := this.process.OrderLinker.GetAnRandomOrderInt()
	err := this.process.OrderLinker.RegisterOrder(Goto,"goto",this.Goto,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	Sign := this.process.OrderLinker.GetAnRandomOrderInt()
	err = this.process.OrderLinker.RegisterOrder(Sign,"sign",this.Sign,this.CheckFunction)
	return StructData.EmptyError
}
func (this *Flow)CheckFunction(OrderInt int,argument []string)StructData.EngineError{
	switch OrderInt{
	case this.orderInt.Goto:{
		if len(argument) != 1{
			return StructData.MakeError(EngineError.Bad,"Goto Only need one for sign to jump to")
		}
		break
	}
	case this.orderInt.Sign:{
		if len(argument) != 1{
			return StructData.MakeError(EngineError.Bad,"Sign Only need one for sign to jump to")
		}
		break
	}
	}
	return StructData.EmptyError
}
func (this *Flow)Goto(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	sign := Arguments[0]

}
func (this *Flow)Sign(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	sign := Arguments[0]
	//创建标记
	
}