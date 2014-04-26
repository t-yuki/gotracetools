// This is a derivative work of roger peppe's stackgraph command.
// For more details, see http://code.google.com/p/rog-go/
//

package traceback

type StackStatus string

const (
	StackStatusChanReceive   StackStatus = "chan receive"
	StackStatusSemAcquire    StackStatus = "semacquire"
	StackStatusRunning       StackStatus = "running"
	StackStatusRunnable      StackStatus = "runnable"
	StackStatusSleep         StackStatus = "sleep"
	StackStatusFinalizerWait StackStatus = "finalizer wait"
	StackStatusSyscall       StackStatus = "syscall"
	StackStatusIOWait        StackStatus = "IO Wait"
)

// Traceback represents the per-goroutine stacktraces generated by GOTRACEBACK=1 environment.
type Traceback struct {
	Reason string
	Stacks []Stack
}

// Stack represents the call stack trace of a goroutine.
type Stack struct {
	ID     int
	Status StackStatus
	Calls  []Call
}

// Call represents a function call of a call stack trace.
type Call struct {
	Func   string
	Source string
	Line   int
	Args   []uint64
}
