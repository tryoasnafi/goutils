package file

import (
	"fmt"
	"os"
	"time"
)

// CopyPerm copies the file permission from source file to destination file
// This function will alter destination file mode.
func CopyPerm(src, dst string) error {
	fs, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("CopyPerm: failed to stat src file %s: %v", src, err)
	}
	if err := os.Chmod(dst, fs.Mode()); err != nil {
		return fmt.Errorf("CopyPerm: failed to chmod dst file: %v", err)
	}
	return nil
}

// CopyTimestamp copies access and modification time of source file 
// to destination file.
func CopyTimestamp(src, dst string) error {
	// when stat or access a file the atime will be updated
	// so the atime can use the current time
	atime := time.Now()
	fs, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("CopyTimestamp: failed to stat src file %s: %v", src, err)
	}
	if err := os.Chtimes(dst, atime, fs.ModTime()); err != nil {
		return fmt.Errorf("CopyTimestamp: failed to change access and modify time: %v", err)
	}
	return nil
}
