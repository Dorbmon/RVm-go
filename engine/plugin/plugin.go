package plugin	//作为引擎的插件系统 为了加快运行速度，允许使用原生语言！
//目前 开发的调用框架的子系统只支持Windows，Linux准备开发！

import (
	"github.com/Dorbmon/RVm/engine"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/engine/type"
	"github.com/Dorbmon/RVm/struct"
	"os"
	"path/filepath"
	"strconv"
)



//回调函数的定义
type Pop func ()PopReturn
type PopReturn struct {
	Type int
	Value string
}
type Push func (DataType int,Data string)PushReturn

type PushReturn struct {
	ok bool
	CallBackError
}
type CallBackError struct {
	Error string
}
type PluginController struct {
	vm *engine.RVM
	pop Pop
	push Push
	process *StructData.Progress
}

func New(vm *engine.RVM,DoInit bool)*PluginController{
	temp := PluginController{}
	temp.vm = vm
	if DoInit{
		temp.Init()
	}
	return &temp
}
func (this *PluginController)SetProcess(Process *StructData.Progress){
	this.process = Process
	return
}
func (this *PluginController)Init(){
	this.pop = func()PopReturn{
		Data,err := this.process.Stack.Pop()
		if StructData.CheckError(err){
			return PopReturn{0,err.Data}
		}
		return PopReturn{Data.Type,Data.Data.(string)}
	}
	this.push = func(DataType int,Data string)PushReturn{
		if !TypeSystem.TypeExist(DataType){	//不存在类型
			return PushReturn{false,CallBackError{"Undefined Type of " + strconv.Itoa(DataType)}}
		}
		this.process.Stack.Push(Data,DataType)
		return PushReturn{true,CallBackError{}}
	}
	return
}
func (this *PluginController)Load(DirName string)StructData.EngineError{	//相对于引擎目录
	Files,err := getDirList(DirName)
	if err != nil{
		return StructData.MakeError(EngineError.Mid,err.Error())
	}
	for n := 0;n < len(Files);n ++{
		FileName := Files[n]	//按照RVM规范，目录名和插件名一致！
		//判断对应的DLL文件是否存在
		ok,err := fileExist(FileName)
		if err != nil{
			return StructData.MakeError(EngineError.Mid,err.Error())
		}
		if !ok {
			return StructData.MakeError(EngineError.Mid,"There is no file named \"" + FileName +".o or .dll\" in the engine dir")
		}

	}

}
func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}
func fileExist(FileName string)(bool,error){
	FileInfo,err := os.Stat(FileName)
	if err != nil{
		return false,err
	}
	return !FileInfo.IsDir(),nil
}