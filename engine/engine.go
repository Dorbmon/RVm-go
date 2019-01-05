package engine

import (
	"fmt"
	"github.com/Dorbmon/RVm/engine/compile"
	"github.com/Dorbmon/RVm/engine/error"
	"github.com/Dorbmon/RVm/struct"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)
type RVM struct {
	MainRegister StructData.RegisterList
	Memory *StructData.Memory
	Progress map[string]*StructData.Progress
	ProgressById map[uint64]*StructData.Progress	//与上方Progress是同一个。不过是索引方式不同。
	DebugFilePath string
	OutputFileName string
	OutputWriter *os.File
	Inited bool
}
func (this *RVM)Init(){
	//初始化debug信息输出目录
	FileDir := GetCurPath()
	ok,_ := PathExists(FileDir + string(os.PathSeparator) + "RJSdebug.txt")
	var File *os.File
	var err error
	if !ok{
		File,err = os.Create(FileDir + string(os.PathSeparator) + "RJSdebug.txt")

	}else{
		File,err = os.Open(FileDir + string(os.PathSeparator) + "RJSdebug.txt")
	}
	if err != nil{	//发生错误。
		fmt.Println(err)
		os.Exit(0)
	}
	this.OutputWriter = File
	this.Progress = make(map[string]*StructData.Progress)
	this.ProgressById = make(map[uint64]*StructData.Progress)

	return
}
func (this *RVM)CreateProgress(Name string)(Ok bool,ProgressId uint64){
	ProgressId = uint64(0)
	if Name == ""{	//无名称形式启动进程
		again:
			ProgressId = rand.Uint64()
			if this.ProgressById[ProgressId].Id != 0{
				goto again
			}
		//创建成功 马上进行进程初始化。
		this.ProgressById[ProgressId].Id = ProgressId
		this.ProgressById[ProgressId].Slience = make([]StructData.Slience,1)
	}else {
		_, ok := this.Progress[Name]
		if ok { //该进程名称已经被占用
			return false, 0
		}
		ProgressId = uint64(0)
	again2:
		ProgressId = rand.Uint64()
		if _,ok := this.ProgressById[ProgressId];ok {
			goto again2
		}
		//NewProgress := StructData.Progress{}
		this.ProgressById[ProgressId] = new(StructData.Progress)
		this.Progress[Name] = this.ProgressById[ProgressId]
		this.Progress[Name].Id = ProgressId
		this.Progress[Name].Slience = make([]StructData.Slience, 1)
		this.Progress[Name].Name = Name

	}
	//初始化指令链接系统，并进行初始化链接
	this.ProgressById[ProgressId].OrderLinker = &StructData.OrderLinker{}
	this.ProgressById[ProgressId].Stack = &StructData.Stack{}
	return true, ProgressId
}
func (this *RVM)LoadUncompiledCode(ProgressId uint64,From uint64,Code StructData.Code)(bool,StructData.EngineError){
	//查找进程
	Progress,ok := this.ProgressById[ProgressId]
	if !ok{
		return false,StructData.MakeError(EngineError.Bad,"Can't Find that Progress")
	}
	//开始载入代码。并且编译代码
	Progress.Compiler = compile.New()
	ok,EngineErr,CompiledCode := Progress.Compiler.Compile(Progress.OrderLinker)
	if !ok{
		this.ThrowError(EngineErr)
		return false,EngineErr
	}
	Progress.CompiledCode = CompiledCode
	return false,StructData.EmptyError
}
func (this *RVM)RunCode(ProgressId uint64)(StructData.EngineError){
	//寻找进程
	Progress,ok := this.ProgressById[ProgressId]
	if !ok{
		return StructData.MakeError(EngineError.Bad,"Can't Find that Progress")
	}
	if Progress.CompiledCode == nil{	//还没有代码
		return StructData.MakeError(EngineError.Bad,"No Code in Progress " + strconv.FormatUint(ProgressId,64))
	}
	if Progress.CompiledCode.Lines == nil{
		return StructData.MakeError(EngineError.Bad,"Error Code.No code in the stack.")
	}
	//开始执行
	for Line := 0;;Line ++{
		if Progress.CompiledCode.Lines[Line] == nil{
			continue
		}
		NowLine := Progress.CompiledCode.Lines[Line]
		//获取到指令立刻调用OrderLinker中对应的函数
			Function := Progress.OrderLinker.GetFunction(NowLine.Order)
			err := Function(NowLine.Data,Progress.Stack)
			if StructData.CheckError(err){	//引擎出现错误
				this.ThrowError(err)
			}
			continue
	}
}
func (this RVM)ThrowError(Error StructData.EngineError){
	//根据错误等级来判断处理方式 首先先输出错误
	Data := strconv.Itoa(Error.Class) + " " + Error.Time + "" + Error.Data
	this.OutputWriter.Write([]byte(Data))
	if Error.Class == EngineError.Bad{
		fmt.Println(Data)
		os.Exit(0)
	}
}
func New()*RVM{
	temp := &RVM{}
	temp.Init()
	return temp
}
func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}