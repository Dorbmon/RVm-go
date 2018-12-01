package orderDefine

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

const(
	SetSystem = iota
	NewVar
	SetVar
)

//这里定义每一条命令需要的参数个数
const(
	NSetSystem = 2
	NNewVar = 1
	NSetVar = 0
)

//定义提供给各个库的参数检测函数的错误信息
func ArgumentTypeError(FuncName string,ArgumentNumber int,ShouldType string)StructData.EngineError{
	return StructData.MakeError(EngineError.Bad,"Argument of " + FuncName + " No." + strconv.Itoa(ArgumentNumber) + " Should be " + ShouldType)
}
func ArgumentNumberError(FuncName string,ArgumentNumber int,ShouldNumber int)StructData.EngineError{
	return StructData.MakeError(EngineError.Bad,"Argument " + strconv.Itoa(ArgumentNumber) + " Should be " + strconv.Itoa(ShouldNumber))
}