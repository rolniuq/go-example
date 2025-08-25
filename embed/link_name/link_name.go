package linkname

import (
	_ "unsafe"
)

//go:linkname RuntimeNow time.Now
func RuntimeNow() (sec int64, nsec int32)

