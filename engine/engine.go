package engine

import(
	"errors"
	"fmt"
	"github.com/Dorbmon/RVm/engine/error"
	StructData "github.com/Dorbmon/RVm/struct"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)
type RVM StructData.RVM
func (this RVM)Init(){
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
	return
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