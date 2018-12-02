package TypeSystem

const(
	String = 0
	Int = iota
	UInt = iota
	Int64 = iota
	UInt64 = iota
	Byte = iota
	Unknown = iota
)
type Type uint

func StringToType(Type string)int{	//用于检验和转换类型
	switch Type{
	case "int":
		return Int
	case "uint":
		return UInt
	case "string":
		return String
	case "int64":
		return Int64
	case "uint64":
		return UInt64
	case "byte":
		return Byte
	case "unknown":
		return Unknown
	}
	return -1
}