package StructData

import (
	"github.com/Dorbmon/RVm/engine"
	"os"
	"sync"
)

type RVM struct {
	MainRegister RegisterList
	Memory *Memory
	Progress map[string]*Progress
	ProgressById map[uint64]*Progress	//与上方Progress是同一个。不过是索引方式不同。
	DebugFilePath string
	OutputFileName string
	OutputWriter *os.File
}
type RegisterList struct{
	CodePointer uint64	//指向内存中当前的代码位置 代号r1
}
type Code struct{
	Lines [][]string	//每一行和每一个空格都分开
}
type Progress struct{	//单个进程
	Id uint64
	Name string
	Slience []Slience
	Memory Memory
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
	Engine *engine.RVM
	Master *Memory
	Variables map[string]*Variable
	UseLock sync.Mutex
}

type EngineError struct{
	Class int	//错误等级，对应error.go中的第一个const
	Data string //具体错误信息
	Time string	//发生时间
}