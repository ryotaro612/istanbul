/*
#cgo CFLAGS: -Wall extern void nsexec(); void
<中略> init(void) { nsexec(); }
*/
import "C"
