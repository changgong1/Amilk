package result

// SecondResult is a result as the returning two result
type SecondResult interface {
	Init() // init a second result
	Ret()  // return the sencond result
}
