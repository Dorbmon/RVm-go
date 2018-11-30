package orderLinker

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/stack"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"math"
	"math/rand"
	"strconv"
)

type OrderLinker struct {
	order []*order
	orderString map[string]*int
}
type order struct{
	Function LinkerFunction
	ArgumentNumber int
}
type Argument struct {
	Data interface{}
	Type TypeSystem.Type
}
type LinkerFunction func(Arguments []string,Stack *stack.Stack)StructData.EngineError	//对方函数不需要判断参数个数，因为在传递调用之前一定是已经完成参数个数检测的。
func (this *OrderLinker)GetAnRandomOrderInt()int{
again:
	r := int(math.Abs(float64(rand.Intn(10000))))
	if this.order[r] == nil{
		return r
	}
	goto again
	return 0
}
func (this *OrderLinker)RegisterOrder(OrderInt int,OrderString string,ArgumentNumber int,LinkFunction LinkerFunction)(StructData.EngineError){
	if this.order[OrderInt] != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + strconv.Itoa(OrderInt) + " Because it has benn existed")
	}
	if this.orderString[OrderString] != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + OrderString + " Because it has benn existed")
	}
	this.order[OrderInt] = &order{}
	this.order[OrderInt].ArgumentNumber = ArgumentNumber
	this.order[OrderInt].Function = LinkFunction
	this.orderString[OrderString] = &OrderInt	//双索引系统，便于编译时的替换
	return StructData.EmptyError
}
func (this *OrderLinker)TranslateToInt(OrderString string)(int,StructData.EngineError){
	if this.orderString[OrderString] == nil{	//命令不存在
		return 0,StructData.MakeError(EngineError.Bad,"Undefined order :" + OrderString)
	}
	return *this.orderString[OrderString],StructData.EmptyError
}
func (this *OrderLinker)GetOrderArgumentNumber(OrderInt int)int{	//假设命令一定存在	获取对应命令的参数个数 用来编译时检查
	return this.order[OrderInt].ArgumentNumber
}
func (this *OrderLinker)GetFunction(OrderInt int)LinkerFunction{	//假设已经存在
	return this.order[OrderInt].Function
}