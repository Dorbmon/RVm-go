package flow

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/hash"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

type Flow struct {
	vm *StructData.RVM
	process *StructData.Progress
	orderInt struct{
		Goto int
		Sign int
		JIO int
	}
	//Goto相关
	//Signs map[string]int
	Sisns hash.ValueHashTable
}

func (this *Flow)RegisterAll(vm *StructData.RVM,Process *StructData.Progress)StructData.EngineError{	//初始化工作
	this.vm = vm
	this.process = Process
	//this.Signs = make(map[string]int)
	Goto := this.process.OrderLinker.GetAnRandomOrderInt()
	err := this.process.OrderLinker.RegisterOrder(Goto,"goto",this.Goto,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	Sign := this.process.OrderLinker.GetAnRandomOrderInt()
	err = this.process.OrderLinker.RegisterOrder(Sign,"sign",this.Sign,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
	JIO := this.process.OrderLinker.GetAnRandomOrderInt()
	this.orderInt.JIO = JIO
	err = this.process.OrderLinker.RegisterOrder(JIO,"jio",this.JIO,this.CheckFunction)
	if StructData.CheckError(err){
		return err
	}
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
	//Line,ok := this.Signs[sign]
	Line := this.Sisns.Get(sign)
	if Line == nil{	//不存在标记
		return StructData.MakeError(EngineError.Bad,"There has not been a sign named " + sign + "before line" + strconv.Itoa(*this.process.Line))
	}
	*this.process.Line = Line.(int)
	return StructData.EmptyError
}
func (this *Flow)Sign(Arguments []string,Stack *StructData.Stack)StructData.EngineError{
	sign := Arguments[0]
	//创建标记 判断标记是否已经存在
	Line := this.Sisns.Get(sign)
	if Line != nil {	//已经存在标记
		return StructData.MakeError(EngineError.Bad,"There has been a sign named " + sign + "on line" + strconv.Itoa(Line.(int)))
	}
	this.Sisns.Put(sign,*this.process.Line)
	//this.Signs[sign] = *this.process.Line
	//fmt.Println("sign...ed")
	return StructData.EmptyError
}
func (this *Flow)JIO(Arguments []string,stack *StructData.Stack)StructData.EngineError{
	sign := Arguments[0]
	//Line,ok := this.Signs[sign]
	Line := this.Sisns.Get(sign)
	if Line == nil{	//不存在标记
		return StructData.MakeError(EngineError.Bad,"There has not been a sign named " + sign + " before line" + strconv.Itoa(*this.process.Line))
	}
	Top,err := stack.Pop()
	if StructData.CheckError(err){
		return err
	}
	if Top.Type != TypeSystem.Int{
		return StructData.MakeError(EngineError.Bad,"Error type of the top of the stack.jio need int")
	}
	if Top.Data.(int) == 1{
		*this.process.Line = Line.(int)
	}
	return StructData.EmptyError
}