package ability

import (
	"bytes"
	"github.com/Dorbmon/RVm/engine"
	vmDeal "github.com/Dorbmon/RVm/engine/ability/vm"
	"github.com/Dorbmon/RVm/engine/orderLinker"
	"github.com/Dorbmon/RVm/struct"
)

func LoadFunction(VM *engine.RVM,OrderLinker *orderLinker.OrderLinker,Progress *StructData.Progress){
	vm := vmDeal.VmDeal{}
	vm.RegisterAll(VM,OrderLinker,Progress)
	return
}




func AddString(Data []string)string{	//拼接文本
	var buffer bytes.Buffer
	for i := 0; i < len(Data); i ++{
		buffer.WriteString(Data[i])
	}
	return buffer.String()
}