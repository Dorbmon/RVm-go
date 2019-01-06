package ability

import (
	"bytes"
	flow2 "github.com/Dorbmon/RVm/engine/ability/flow"
	"github.com/Dorbmon/RVm/engine/ability/math"
	//"github.com/Dorbmon/RVm/engine"
	memory2 "github.com/Dorbmon/RVm/engine/ability/memory"
	"github.com/Dorbmon/RVm/engine/ability/vm"
	"github.com/Dorbmon/RVm/struct"
)

func LoadFunction(VM *StructData.RVM,Progress *StructData.Progress){
	vm := new(vmDeal.VmDeal)
	vm.RegisterAll(VM,Progress.OrderLinker,Progress)
	Progress.Abilities["vm"] = vm
	memory := new(memory2.Memory)
	memory.RegisterAll(VM,Progress.OrderLinker,Progress)
	Progress.Abilities["memory"] = memory
	maths := new(math.Math)
	maths.RegisterAll(VM,Progress.OrderLinker,Progress)
	Progress.Abilities["maths"] = maths
	flow := flow2.Flow{}
	flow.RegisterAll(VM,Progress)
	Progress.Abilities["flow"] = flow
	return
}
func AddString(Data []string)string{	//拼接文本
	var buffer bytes.Buffer
	for i := 0; i < len(Data); i ++{
		buffer.WriteString(Data[i])
	}
	return buffer.String()
}