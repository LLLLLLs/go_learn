//@author: lls
//@time: 2020/04/27
//@desc:

package generate

import _ "golang.org/x/tools/go/packages"

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
