//go:build darwin
// +build darwin

package file

import (
	"os"
	"syscall"
	"time"
)

func statTime(fi os.FileInfo) Timestamp {
	file := fi.Sys().(*syscall.Stat_t)

	return Timestamp{
		AccessTime: time.Unix(file.Atimespec.Unix()),
		ModifyTime: time.Unix(file.Mtimespec.Unix()),
		ChangeTime: time.Unix(file.Ctimespec.Unix()),
		BirthTime: time.Unix(file.Birthtimespec.Unix()),
	}
}
