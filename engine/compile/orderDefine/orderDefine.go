package orderDefine

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