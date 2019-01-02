package orderLinker

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"math"
	"math/rand"
	"strconv"
)

type OrderLinker struct {
	Order []*StructData.Order
	OrderString map[string]*int
}
type order StructData.Order
type Argument struct {
	Data interface{}
	Type TypeSystem.Type
}
type LinkerFunction StructData.LinkerFunction	//对方函数不需要判断参数个数，因为在传递调用之前一定是已经完成参数个数检测的。
type CheckFunction StructData.CheckFunction	//用来在编译时期检测参数类型等信息
func (this *OrderLinker)GetAnRandomOrderInt()int{
again:
	r := int(math.Abs(float64(rand.Intn(10000))))
	if this.Order[r] == nil{
		return r
	}
	goto again
	return 0
}
func (this *OrderLinker)RegisterOrder(OrderInt int,OrderString string,LinkFunctions LinkerFunction,CheckFunctions CheckFunction)(StructData.EngineError){
	if this.Order[OrderInt] != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + strconv.Itoa(OrderInt) + " Because it has benn existed")
	}
	if this.OrderString[OrderString] != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + OrderString + " Because it has benn existed")
	}
	this.Order[OrderInt] = &StructData.Order{}
	this.Order[OrderInt].Function = (StructData.LinkerFunction)(LinkFunctions)
	this.Order[OrderInt].CheckFunction = (StructData.CheckFunction)(CheckFunctions)
	this.OrderString[OrderString] = &OrderInt	//双索引系统，便于编译时的替换
	return StructData.EmptyError
}
func (this *OrderLinker)TranslateToInt(OrderString string)(int,StructData.EngineError){
	if this.OrderString[OrderString] == nil{	//命令不存在
		return 0,StructData.MakeError(EngineError.Bad,"Undefined order :" + OrderString)
	}
	return *this.OrderString[OrderString],StructData.EmptyError
}
func (this *OrderLinker)GetFunction(OrderInt int)StructData.LinkerFunction{	//假设已经存在
	return this.Order[OrderInt].Function
}
func (this *OrderLinker)GetCheckFunction(OrderInt int)StructData.CheckFunction{
	return this.Order[OrderInt].CheckFunction
}