package StudentEnum

type Medal int16

func (t Medal) ToInt16() int16 {
	return int16(t)
}

const (
	E   Medal = 1
	D   Medal = 2
	C   Medal = 3
	B   Medal = 4
	A   Medal = 5
	S   Medal = 6
	SS  Medal = 7
	SSS Medal = 8
)
