package stack

import (
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"strconv"
)

func New()*StructData.Stack{
	cStack := &Stack{}
	cStack.MaxDeep = 0
	temp := &StructData.Stack{}
	temp.RStack = cStack
	temp.Push = cStack.Push
	temp.Pop = cStack.Pop
	temp.SetMaxDeep = cStack.SetMaxDeep
	temp.Empty = cStack.Empty
	return temp
}
type Stack struct{	//真实栈结构
	TopNode *StructData.Node
	MaxDeep int
	NowDeep int
	Fuck string
}
func (this *Stack)Empty()bool{	//返回是否为空
	return this.NowDeep == 0
}
func (this *Stack)SetMaxDeep(deep int){	//不能删除已经存在的数据
	this.MaxDeep = deep
	return
}
func (this *Stack)Push(Data interface{},Type int)(StructData.EngineError){
	if this.MaxDeep == 0{	//不限制栈深度
		goto Push
	}
	if this.NowDeep < this.MaxDeep{

	}else{
		return StructData.MakeError(EngineError.Mid,"There is no space for stack,Because Max Deep is" + strconv.Itoa(this.MaxDeep))
	}
	Push:
	var TargetNode *StructData.Node
	if this.Empty(){	//栈中为空
		this.TopNode = new(StructData.Node)
		TargetNode = this.TopNode
	}else{
		tempNode := *this.TopNode
		this.TopNode = new(StructData.Node)
		this.TopNode.NodeAfter = &tempNode
	}
	TargetNode.Data = new(StructData.StackObject)
	TargetNode.Data.Data = Data
	TargetNode.Data.Type = Type
	this.NowDeep ++
	return StructData.EmptyError
}
func (this *Stack)Pop()(StructData.StackObject,StructData.EngineError){
	if this.Empty(){
		return StructData.StackObject{},StructData.MakeError(EngineError.Mid,"There is nothing in the stack.Check Your code")
	}
	Data := *this.TopNode.Data
	if this.TopNode.NodeAfter == nil{
		this.TopNode = nil
	}else{
		this.TopNode = this.TopNode.NodeAfter
	}
	this.NowDeep --
	return Data,StructData.EmptyError
}