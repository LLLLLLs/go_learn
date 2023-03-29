package define

type Test struct {
}

func (t Test) Exported() {

}

func (t Test) unexported() {

}
