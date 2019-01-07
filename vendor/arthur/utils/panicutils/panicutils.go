package panicutils

// exp 为 false 则 panic what
func TrueOrPanic(exp bool, what interface{}) {
	if exp == false {
		panic(what)
	}
}

// exp 为 true 则 panic what
func FalseOrPanic(exp bool, what interface{}) {
	if exp == true {
		panic(what)
	}
}

// 有 err 则 panic
func OkOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
