package file

import (
	"os"
	"time"
)

type Timestamp struct {
	AccessTime time.Time
	ModifyTime time.Time
	ChangeTime time.Time
	BirthTime  time.Time
}

// StatTime returns the file access time, modified time, change time, and birth time.
// But linux doesn't have birth time, and windows doesn't have change time,
// use modified time as the fallback. 
func StatTime(fi os.FileInfo) Timestamp {
	return statTime(fi)
}
