package ability

import (
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