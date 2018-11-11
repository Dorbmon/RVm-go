package engine

import(
	"errors"
	StructData "github.com/Dorbmon/RVm/struct"
	"math/rand"
)
type RVM StructData.RVM
func (this RVM)Init(){

}
func (this RVM)CreateProgress(Name string)(Ok bool,ProgressId uint64){
	if Name == ""{	//无名称形式启动进程
		ProgressId := uint64(0)
		again:
			ProgressId = rand.Uint64()
			if this.ProgressById[ProgressId].Id != 0{
				goto again
			}
		//创建成功 马上进行进程初始化。
		this.ProgressById[ProgressId].Id = ProgressId
		this.ProgressById[ProgressId].Slience = make([]StructData.Slience,1)
		return true,ProgressId
	}
	_,ok := this.Progress[Name]
	if ok {	//该进程名称已经被占用
		return false,0
	}
	ProgressId = uint64(0)
again2:
	ProgressId = rand.Uint64()
	if this.ProgressById[ProgressId].Id != 0{
		goto again2
	}
	NewProgress := StructData.Progress{}
	this.ProgressById[ProgressId] = &NewProgress
	this.Progress[Name] = &NewProgress
	this.Progress[Name].Id = ProgressId
	this.Progress[Name].Slience = make([]StructData.Slience,1)
	this.Progress[Name].Name = Name
	return true,ProgressId
}
func (this RVM)LoadCode(ProgressId uint64,From uint64,Code StructData.Code)(ok bool,err error){
	//查找进程
	Progress,ok := this.ProgressById[ProgressId]
	if !ok{
		return false,errors.New("Can't Find that Progress")
	}
	Progress.

}
func New()*RVM{
	temp := &RVM{}
	temp.Init()
	return temp
}