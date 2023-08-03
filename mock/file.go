package mock

import "io"

type MockFile struct {
	Bytes     []byte
	lastRead  int
	ReturnEOF bool //FIle will return EOF at once
}

func (l *MockFile) Write(b []byte) (n int, err error) {
	l.Bytes = append(l.Bytes, b...)

	return 0, nil
}

func (l *MockFile) WriteString(s string) (n int, err error) {
	return l.Write([]byte(s))
}

func (l *MockFile) Close() error {
	return nil
}

func (l *MockFile) Read(b []byte) (n int, err error) {
	if l.ReturnEOF {

		return n, io.EOF

	} else {
		n = copy(b, l.Bytes[l.lastRead:])
		l.lastRead = l.lastRead + n
	}

	if l.lastRead >= len(l.Bytes) {
		return n, io.EOF
	}

	return
}
