package StructData

import (
	"github.com/Dorbmon/RVm/engine/type"
	"os"
	"sync"
	"time"
)

type RVM struct {
	MainRegister RegisterList
	Memory *Memory
	Progress map[string]*Progress
	ProgressById map[uint64]*Progress	//与上方Progress是同一个。不过是索引方式不同。
	DebugFilePath string
	OutputFileName string
	OutputWriter *os.File
	Compile func (OrderLinker *OrderLinker)(bool,EngineError,*CompiledCode)
	RunCode func (ProgressId uint64)(EngineError)
}
type RegisterList struct{
	CodePointer uint64	//指向内存中当前的代码位置 代号r1
}
type Code struct{
	Lines [][]string	//每一行和每一个空格都分开
}
type CodeArea struct{
	Codes []*Code
	NowPointer *Code	//指向当前执行的代码
}
type OrderLinker struct {
	Order []*Order
	OrderString map[string]*int
	GetAnRandomOrderInt func ()int
	RegisterOrder func (OrderInt int,OrderString string,LinkFunctions LinkerFunction,CheckFunctions CheckFunction)(EngineError)
	TranslateToInt func (OrderString string)(int,EngineError)
	GetFunction func (OrderInt int)LinkerFunction
	GetCheckFunction func (OrderInt int)CheckFunction
}
type Order struct{
	Function LinkerFunction
	CheckFunction CheckFunction
}
type Stack struct{
	TopNode *Node
	MaxDeep int
	NowDeep int
	Pop func ()(*StackObject,EngineError)
	Push func (Data interface{},Type int)(EngineError)
	SetMaxDeep func (deep int)
	Empty func ()bool
}
type Node struct{
	NodeAfter *Node
	Data *StackObject
}
type LinkerFunction func(Arguments []string,Stack *Stack)EngineError	//对方函数不需要判断参数个数，因为在传递调用之前一定是已经完成参数个数检测的。
type CheckFunction func(OrderInt int,Arguments []string)EngineError	//用来在编译时期检测参数类型等信息
type Progress struct{	//单个进程
	Id uint64
	Name string
	Slience []Slience
	Memory Memory
	Compiler *Compiler
	CompiledCode *CompiledCode
	OrderLinker *OrderLinker
	Stack *Stack
}
type Compiler struct{
	RealObj interface{}	//用来储存真实对象
	Code Code
	LoadCode func (Code Code)
	Compile func (OrderLinker *OrderLinker)(bool,EngineError,*CompiledCode)
	TranslateStringToINT func (code Code,OrderLinker *OrderLinker)(*CompiledCode,EngineError)
}
type Slience struct{	//进程切片
	From uint64
	Long uint64
}
type Value struct{
	Value interface {}
	Type int
}
type Variable struct{
	Value *Value
	Name string
}
type Memory struct{
	Engine *RVM
	Master *Memory
	Variables map[string]*Variable
	UseLock sync.Mutex
}

type EngineError struct{
	Class int	//错误等级，对应error.go中的第一个const
	Data string //具体错误信息
	Time string	//发生时间
}
type CodeValue struct {
	Value interface{}
	Type TypeSystem.Type
}
type CompiledCode struct {
	Lines []*CodeLine
}
type CodeLine struct {
	Order int
	Data []string
}
var EmptyError EngineError
func CheckError(error EngineError)bool{
	return error.Class != 0
}
func MakeError(Class int,Data string)EngineError{
	temp := EngineError{}
	temp.Time = time.Now().String()
	temp.Data = Data
	temp.Class = Class
	return temp
}

type StackObject struct{	//栈中数据对象
	Data interface{}
	Type int
}