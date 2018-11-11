package memory

import "github.com/Dorbmon/RVm/struct"

type Memory StructData.Memory

func (this Memory) Init(Master *StructData.Memory,From uint64,Long uint64){	//Master为主内存。如果为从内存将依附于主内存 如果允许让瑞雪VM管理全部内存，则可以设置long为0.默认为之后的全部内存
	if Master != nil{	//需要设置主内存
		this.Master = Master
	}

}