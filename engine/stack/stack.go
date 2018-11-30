package stack

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
)

type Stack struct{
	TopNode *node
	MaxDeep uint
}

type node struct{
	NodeAfter *node
	Data *StructData.StackObject
}
func (this *Stack)Empty()bool{	//返回是否为空
	return this.TopNode == nil
}
func (this *Stack)SetMaxDeep(deep int){	//不能删除已经存在的数据

}
func (this *Stack)Push(Data interface{},Type int)(StructData.EngineError){
	var TargetNode *node
	if this.TopNode == nil{	//栈中为空
		this.TopNode = &node{}
		TargetNode = this.TopNode
	}else{
		tempNode := this.TopNode
		this.TopNode = &node{}
		this.TopNode.NodeAfter = tempNode
	}
	TargetNode.Data = &StructData.StackObject{}
	TargetNode.Data.Data = Data
	TargetNode.Data.Type = Type
	return StructData.EmptyError
}
func (this *Stack)Pop()(*StructData.StackObject,StructData.EngineError){
	if this.TopNode == nil{
		return &StructData.StackObject{},StructData.MakeError(EngineError.Mid,"There is nothing in the stack.Check Your code")
	}
	Data := this.TopNode.Data
	if this.TopNode.NodeAfter == nil{
		this.TopNode = nil
	}else{
		this.TopNode = this.TopNode.NodeAfter
	}
	return Data,StructData.EmptyError
}