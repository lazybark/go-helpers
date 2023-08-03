package fsw

import (
	"io"
	"io/fs"
	"syscall"
	"time"
)

// IFile represents file in the filesystem that's used to log events
type IFile interface {
	Chdir() error
	Chmod(mode fs.FileMode) error
	Chown(uid int, gid int) error
	Close() error
	Fd() uintptr
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadDir(n int) ([]fs.DirEntry, error)
	ReadFrom(r io.Reader) (n int64, err error)
	Readdir(n int) ([]fs.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Stat() (fs.FileInfo, error)
	Sync() error
	SyscallConn() (syscall.RawConn, error)
	Truncate(size int64) error
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	WriteString(s string) (n int, err error)
}

// IFileWriter is an interface to mock file writing operations
type IFileWriter interface {
	Close() error
	Write(b []byte) (n int, err error)
	WriteString(s string) (n int, err error)
}

// IFileReader is an interface to mock file reading operations
type IFileReader interface {
	Close() error
	Read(b []byte) (n int, err error)
}
