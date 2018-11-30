package orderLinker

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

type OrderLinker struct {
	order []*order
}
type order struct{
	Function LinkerFunction
	ArgumentNumber int
}
type Argument struct {
	Data interface{}
	Type TypeSystem.Type
}
type LinkerFunction func(Arguments []Argument)	//对方函数不需要判断参数个数，因为在传递调用之前一定是已经完成参数个数检测的。
func (this *OrderLinker)RegisterOrder(OrderInt int,ArgumentNumber int,LinkFunction LinkerFunction)(StructData.EngineError){
	if this.order[OrderInt] != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + strconv.Itoa(OrderInt) + " Because it has benn existed")
	}
	this.order[OrderInt] = &order{}
	this.order[OrderInt].ArgumentNumber = ArgumentNumber
	this.order[OrderInt].Function = LinkFunction
	return StructData.EmptyError
}
