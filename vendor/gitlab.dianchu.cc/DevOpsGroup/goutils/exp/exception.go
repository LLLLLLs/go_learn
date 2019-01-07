package exp

import (
	"fmt"
	"log"

	"github.com/facebookgo/stack"
)

type ExType int

var (
	//dunno     = []byte("???")
	//centerDot = []byte("·")
	//dot       = []byte(".")
	//slash     = []byte("/")

	DEBUG = true
	//logger = logs.New("exp")
)

const (
	//LoginEx     ExType = 1 + iota
	//NotMemberEx
	//NotMethodEx

	InnerEx = 10000
	TestEx  = 999999
)

type Exception struct {
	ExType
	Info    string
	Stack   string
	Message string
}

func NewException(typ ExType) Exception {
	return NewCustomException(typ, "", "")
}
func NewCustomException(typ ExType, msg, stack string) Exception {
	return Exception{typ, msg, stack, fmt.Sprintf("[Recovery] panic recovered:\n%s\n%s ", msg, stack)}
}

func NewInnerEx(msg, stack string) Exception {
	return Exception{InnerEx, msg, stack, fmt.Sprintf("[Recovery] panic recovered:\n%s\n%s ", msg, stack)}
}

func Try(fun func(), handler func(ex Exception)) {
	defer func() {
		if err := recover(); err != nil {

			switch err.(type) {
			case Exception:
				handler(err.(Exception))
			default:
				c := stack.Callers(3)
				//stack := stack(4)
				//errmsg := fmt.Sprintf("[Recovery] panic recovered:\n%s\n%s ", err, c)
				// 如果非DEBUG模式 错误堆栈将不会直接打印在控制台 需要上一层recover捕获堆栈后输出，否则没有错误堆栈输出
				if DEBUG {
					log.Printf("[Recovery] panic recovered:\n%s\n%s", err, c)
				}
				handler(NewInnerEx(fmt.Sprintf("%s", err), c.String()))
			}
		}
	}()
	fun()
}

func SetDebug(debug bool) {
	DEBUG = debug
}

//
//func stack(skip int) []byte {
//	buf := new(bytes.Buffer) // the returned data
//	// As we loop, we open files and read them. These variables record the currently
//	// loaded file.
//	var lines [][]byte
//	var lastFile string
//	for i := skip; ; i++ { // Skip the expected number of frames
//		pc, file, line, ok := runtime.Caller(i)
//		if !ok {
//			break
//		}
//		// Print this much at least.  If we can't find the source, it won't show.
//		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
//		if file != lastFile {
//			data, err := ioutil.ReadFile(file)
//			if err != nil {
//				continue
//			}
//			lines = bytes.Split(data, []byte{'\n'})
//			lastFile = file
//		}
//		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
//	}
//	return buf.Bytes()
//}

//// source returns a space-trimmed slice of the n'th line.
//func source(lines [][]byte, n int) []byte {
//	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
//	if n < 0 || n >= len(lines) {
//		return dunno
//	}
//	return bytes.TrimSpace(lines[n])
//}
//
//// function returns, if possible, the name of the function containing the PC.
//func function(pc uintptr) []byte {
//	fn := runtime.FuncForPC(pc)
//	if fn == nil {
//		return dunno
//	}
//	name := []byte(fn.Name())
//	// The name includes the path name to the package, which is unnecessary
//	// since the file name is already included.  Plus, it has center dots.
//	// That is, we see
//	//	runtime/debug.*T·ptrmethod
//	// and want
//	//	*T.ptrmethod
//	// Also the package path might contains dot (e.g. code.google.com/...),
//	// so first eliminate the path prefix
//	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
//		name = name[lastslash+1:]
//	}
//	if period := bytes.Index(name, dot); period >= 0 {
//		name = name[period+1:]
//	}
//	name = bytes.Replace(name, centerDot, dot, -1)
//	return name
//}
