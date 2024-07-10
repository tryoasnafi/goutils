//go:build windows
// +build windows

package file

import (
	"os"
	"syscall"
	"time"
)

func statTime(fi os.FileInfo) Timestamp {
	file := fi.Sys().(*syscall.Win32FileAttributeData)

	return Timestamp{
		AccessTime: time.Unix(0, file.LastAccessTime.Nanoseconds()),
		ModifyTime: time.Unix(0, file.LastWriteTime.Nanoseconds()),
		// change time unsupported, use modify time instead
		ChangeTime: time.Unix(0, file.LastWriteTime.Nanoseconds()),
		BirthTime: time.Unix(0, file.CreationTime.Nanoseconds()),
	}
}
