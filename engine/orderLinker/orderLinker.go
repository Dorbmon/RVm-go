package orderLinker

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/hash"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"math"
	"math/rand"
	"strconv"
)
const(
	MaxOrderNumber = 100000	//最大指令数，理论上越大越好（因为RVM的算法导致）过小会导致引擎大量时间用在申请ID上，甚至死循环
)
type OrderLinker struct {
	Order hash.ValueHashTable
	// 旧RVM
	//Order map[int]*StructData.Order
	OrderString map[string]int
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
	r := int(math.Abs(float64(rand.Intn(MaxOrderNumber))))
	if this.Order.Get(r) == nil{
		return r
	}
	goto again
	//this.Order[r] = &StructData.Order{}
	this.Order.Put(r,StructData.Order{})
	return 0
}
func New()*StructData.OrderLinker{
	cLinker := &OrderLinker{}

	cLinker.OrderString = make(map[string]int)

	temp := &StructData.OrderLinker{}
	temp.CLinker = cLinker
	temp.RegisterOrder = cLinker.RegisterOrder
	temp.GetAnRandomOrderInt = cLinker.GetAnRandomOrderInt
	temp.GetCheckFunction = cLinker.GetCheckFunction
	//cLinker.Order = temp.Order
	temp.GetFunction = cLinker.GetFunction
	temp.TranslateToInt = cLinker.TranslateToInt
	return temp
}
func (this *OrderLinker)RegisterOrder(OrderInt int,OrderString string,LinkFunctions StructData.LinkerFunction,CheckFunctions StructData.CheckFunction)(StructData.EngineError){
	if this.Order.Get(OrderInt) != nil{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + strconv.Itoa(OrderInt) + " Because it has benn existed")
	}
	if this.OrderString[OrderString] != 0{	//该命令已经存在
		return StructData.MakeError(EngineError.Bad,"Error Linking of Func " + OrderString + " Because it has benn existed")
	}
	temp := &StructData.Order{}
	temp.Function = LinkFunctions
	temp.CheckFunction = CheckFunctions
	this.OrderString[OrderString] = OrderInt	//双索引系统，便于编译时的替换
	this.Order.Put(OrderInt,temp)
	return StructData.EmptyError
}
func (this *OrderLinker)TranslateToInt(OrderString string)(int,StructData.EngineError){
	if this.OrderString[OrderString] == 0{	//命令不存在
		return 0,StructData.MakeError(EngineError.Bad,"Undefined order :" + OrderString)
	}
	return this.OrderString[OrderString],StructData.EmptyError
}
func (this *OrderLinker)GetFunction(OrderInt int)StructData.LinkerFunction{	//假设已经存在
	return this.Order.Get(OrderInt).(*StructData.Order).Function
	//return this.Order[OrderInt].Function
}
func (this *OrderLinker)GetCheckFunction(OrderInt int)StructData.CheckFunction{
	return this.Order.Get(OrderInt).(*StructData.Order).CheckFunction
	//return this.Order[OrderInt].CheckFunction
}