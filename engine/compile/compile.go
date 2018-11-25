package compile

import (
	"github.com/Dorbmon/RVm/engine/compile/orderDefine"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
)

type Compiler struct{
	Code StructData.Code
}

func (this *Compiler)LoadCode(Code StructData.Code){
	this.Code = Code
	return
}
func (this *Compiler)Compile()(bool,StructData.EngineError,*StructData.CompiledCode){
	//开始编译
	if len(this.Code.Lines[0]) == 0{
		//没有代码
		return false,StructData.MakeError(EngineError.Bad,"There is no code in the Compiler."),&StructData.CompiledCode{}
	}
	//首先进行翻译，将所有的字符串指令翻译为int
	CompiledCode,err := this.TranslateStringToINT(this.Code)
	if StructData.CheckError(err){
		return false,err,&StructData.CompiledCode{}
	}
	return true,StructData.EmptyError,CompiledCode
}
func (this *Compiler)TranslateStringToINT(code StructData.Code)(*StructData.CompiledCode,StructData.EngineError){
	//开始处理
	var err StructData.EngineError
	result := &StructData.CompiledCode{}
	for Line := 0;Line < len(code.Lines);Line ++{
		//开始对每一行进行处理，根据硬性规定，每一行只有第一个单词为命令
		result.Lines[Line] = &StructData.CodeLine{}
		result.Lines[Line].Order,err = toInt(code.Lines[Line][0])
		if StructData.CheckError(err){	//出现不存在的指令
			return nil,err
		}
		result.Lines[Line].Data = code.Lines[Line][1:]
		continue
	}
	return result,StructData.EmptyError
}
func toInt(key string)(int,StructData.EngineError){
	switch key {
	case "SetSystem":{
		return orderDefine.SetSystem,StructData.EmptyError
	}
	case "NewVar":{
		return orderDefine.NewVar,StructData.EmptyError
	}
	case "SetVar":{
		return orderDefine.SetVar,StructData.EmptyError
	}
	default:{
		return 0,StructData.MakeError(EngineError.Bad,"There is no order called " + key)
	}
	}
}