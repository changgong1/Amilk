package result

// SecondResult is a result as the returning two result
type SecondResult interface {
	InitResult() // init a second result
	Ret()  // return the sencond result
}
