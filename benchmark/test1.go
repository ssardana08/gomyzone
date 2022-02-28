package benchmark

type X struct {
	a        string
	b        int
	c        int64
	d        float64
	e        string
	f        int
	g        string
	h        float64
	sentence string
}

func returnStruct() X {
	return X{
		a:        "a",
		b:        4756575,
		c:        586876,
		d:        586768,
		e:        "ehhsjhsuio7081uijslkjnlmsp89i92i0298usljhlksk",
		f:        78977,
		g:        "ghlkjsjouyipsy9696082568561802t2ihbskgiylt6206718yhuyouyfygugougouyfou",
		h:        108076543456.1,
		sentence: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
	}
}

func returnStructPointer() *X {
	return &X{
		a:        "a",
		b:        4756575,
		c:        586876,
		d:        586768,
		e:        "ehhsjhsuio7081uijslkjnlmsp89i92i0298usljhlksk",
		f:        78977,
		g:        "ghlkjsjouyipsy9696082568561802t2ihbskgiylt6206718yhuyouyfygugougouyfou",
		h:        108076543456.1,
		sentence: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
	}
}

func fetchStruct() {
	x := returnStruct()
	transformStruct(x)
}

func fetchStructPointer() {
	x := returnStructPointer()
	transformStructPointer(x)
}

func transformStruct(x X)         {}
func transformStructPointer(x *X) {}
