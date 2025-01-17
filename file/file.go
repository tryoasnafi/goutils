package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyPerm copies the file permission from source file to destination file
// This function will alter destination file mode.
func CopyPerm(src, dst string) error {
	fs, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat src file %s: %v", src, err)
	}
	if err := os.Chmod(dst, fs.Mode()); err != nil {
		return fmt.Errorf("failed to chmod dst file: %v", err)
	}
	return nil
}

// CopyTimestamp copies access and modification time of source file 
// to destination file.
func CopyTimestamp(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat src file %s: %v", src, err)
	}
	fileTime := StatTime(fi)
	if err := os.Chtimes(dst, fileTime.AccessTime, fi.ModTime()); err != nil {
		return fmt.Errorf("failed to change access and modify time: %v", err)
	}
	return nil
}

// Copy copy file from source file to destination file. If destination
// file already exists, it will be truncates.
func Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed open source file: %v", err)
	}
	defer srcFile.Close()
	
	if err := mkdir(dst); err != nil {
		return fmt.Errorf("failed mkdir destination path: %v", err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed create destination file: %v", err)
	}
	defer dstFile.Close()

	if n, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed copy file: written: %d, %v", n, err)
	}

	dstFile.Sync()

	return nil
}

// Move move file with preserve permission and timestamp 
func Move(src, dst string) error {
	if err := Copy(src, dst); err != nil {
		return err
	}
	if err := CopyPerm(src, dst); err != nil {
		return err
	}
	if err := CopyTimestamp(src, dst); err != nil {
		return err
	}
    if err := os.Remove(src); err != nil {
        return fmt.Errorf("failed remove source file: %v", err)
    }
    return nil
}

// mkdir make path directories if not exists with permission 0777
// respect umask
func mkdir(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0777)
	}
	return nil
}