package benchmark

const bigStructSize = 1

type bigStruct struct {
	a [bigStructSize]int
}

func newBigStruct() bigStruct {
	var b bigStruct
	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return b
}
func newBigStructPtr() *bigStruct {
	var b bigStruct
	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return &b
}
