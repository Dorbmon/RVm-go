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