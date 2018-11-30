package compile

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/orderLinker"
	"github.com/Dorbmon/RVm/struct"
)

type Compiler struct{
	Code StructData.Code
}

func (this *Compiler)LoadCode(Code StructData.Code){
	this.Code = Code
	return
}
func (this *Compiler)Compile(OrderLinker *orderLinker.OrderLinker)(bool,StructData.EngineError,*StructData.CompiledCode,){
	//开始编译
	if len(this.Code.Lines[0]) == 0{
		//没有代码
		return false,StructData.MakeError(EngineError.Bad,"There is no code in the Compiler."),&StructData.CompiledCode{}
	}
	//首先进行翻译，将所有的字符串指令翻译为int
	CompiledCode,err := this.TranslateStringToINT(this.Code,OrderLinker)
	if StructData.CheckError(err){
		return false,err,&StructData.CompiledCode{}
	}
	return true,StructData.EmptyError,CompiledCode
}
func (this *Compiler)TranslateStringToINT(code StructData.Code,OrderLinker *orderLinker.OrderLinker)(*StructData.CompiledCode,StructData.EngineError){
	//开始处理
	var err StructData.EngineError
	result := &StructData.CompiledCode{}
	for Line := 0;Line < len(code.Lines);Line ++{
		//开始对每一行进行处理，根据硬性规定，每一行只有第一个单词为命令
		result.Lines[Line] = &StructData.CodeLine{}
		result.Lines[Line].Order,err = toInt(code.Lines[Line][0],len(code.Lines[Line]) - 1,OrderLinker)
		if StructData.CheckError(err){	//出现不存在的指令
			return nil,err
		}
		result.Lines[Line].Data = code.Lines[Line][1:]
		continue
	}
	return result,StructData.EmptyError
}
var ErrorArgument = StructData.MakeError(EngineError.Bad,"Error number of arguments")
func toInt(key string, NumberOfArguments int,OrderLinker *orderLinker.OrderLinker)(int,StructData.EngineError){		//并且检查参数个数
//作为虚拟机，需要进行高强度的检测。 包括前后参数的设置
	//获取对应的指令
	orderInt,err := OrderLinker.TranslateToInt(key)
	if StructData.CheckError(err){
		return 0,err
	}
	orderArgumentNumber := OrderLinker.GetOrderArgumentNumber(orderInt)
	if orderArgumentNumber == -1 {	//任意参数个数
		return orderInt,StructData.EmptyError
	}
	if orderArgumentNumber != NumberOfArguments{	//参数个数不匹配
		return 0,ErrorArgument
	}
	return orderInt,StructData.EmptyError
}