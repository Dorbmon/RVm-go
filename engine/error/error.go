package EngineError


const (	//错误等级
	Bad = 2	//需要暂停停止引擎的错误
	NotBad = iota
	Mid = iota	//此等级错误将调用异常代码进行处理 （还未完成异常部分）
)