//go:build linux
// +build linux

package file

import (
	"os"
	"syscall"
	"time"
)

func statTime(fi os.FileInfo) Timestamp {
	file := fi.Sys().(*syscall.Stat_t)
	return Timestamp{
		AccessTime: time.Unix(file.Atim.Unix()),
		ModifyTime: time.Unix(file.Mtim.Unix()),
		ChangeTime: time.Unix(file.Ctim.Unix()),
		// birth time unsupported, use modify time instead
		BirthTime: time.Unix(file.Mtim.Unix()),
	}
}
