package StructData

type RVM struct {
	MainRegister RegisterList
	Memory Memory
	Progress map[string]*Progress
	ProgressById map[uint64]*Progress	//与上方Progress是同一个。不过是索引方式不同。
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
type Memory struct{
	Master *Memory
}