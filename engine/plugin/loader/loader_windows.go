package loader	//Only Windows
import "syscall"

type Loader struct {
	Handle syscall.Handle
}
type Function struct {
	function uintptr
}
func New(FileName string)(*Loader,error){
	temp := Loader{}
	Handle,err := syscall.LoadLibrary(FileName)
	if err != nil{
		return nil,err
	}
	temp.Handle = Handle
	return &temp,nil
}
func (this *Loader)GetFunc(FuncName string)(Function,error){
	proc,err := syscall.GetProcAddress(this.Handle,FuncName)
	if err != nil{
		return Function{},err
	}
	
}



