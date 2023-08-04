package mock

import "io"

type MockWriteReader struct {
	Bytes            []byte
	lastRead         int
	ReturnEOF        bool //FIle will return EOF at once
	DontReturEOFEver bool
}

func (l *MockWriteReader) Write(b []byte) (n int, err error) {
	l.Bytes = append(l.Bytes, b...)

	return len(b), nil
}

func (l *MockWriteReader) WriteString(s string) (n int, err error) {
	return l.Write([]byte(s))
}

func (l *MockWriteReader) Close() error {
	return nil
}

func (l *MockWriteReader) Read(b []byte) (n int, err error) {
	if l.ReturnEOF {

		return n, io.EOF

	} else {
		n = copy(b, l.Bytes[l.lastRead:])
		l.lastRead = l.lastRead + n
	}

	if l.lastRead >= len(l.Bytes) && !l.DontReturEOFEver {
		return n, io.EOF
	}

	return n, nil
}

func (l *MockWriteReader) SetLastRead(n int) {
	l.lastRead = n
}
