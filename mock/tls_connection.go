package mock

import (
	"io"
	"net"
	"time"
)

// MockTLSConnection mocks net.Conn interface
type MockTLSConnection struct {
	BytesToRead  []byte //Bytes will be returned, mocking network reading
	lastRead     int    //Number of last byte that was read
	WriteBytesTo []byte //Bytes will be written, mocking network write

	ReturnEOF bool //ReturnEOF marks that Read() should return EOF

	AskedToBeClosed bool

	LocAddr net.Addr
	RemAddr net.Addr
}

func (m *MockTLSConnection) Read(b []byte) (n int, err error) {
	if m.ReturnEOF {
		return 0, io.EOF
	} else {
		n = copy(b, m.BytesToRead[m.lastRead:])
		m.lastRead = m.lastRead + n
	}

	return
}

// Write replaces current buffer with b
func (m *MockTLSConnection) Write(b []byte) (n int, err error) {
	m.WriteBytesTo = b

	return len(b), nil
}

func (m *MockTLSConnection) Close() error {
	m.AskedToBeClosed = true

	return nil
}

func (m *MockTLSConnection) LocalAddr() net.Addr                { return m.LocAddr }
func (m *MockTLSConnection) RemoteAddr() net.Addr               { return m.RemAddr }
func (m *MockTLSConnection) SetDeadline(t time.Time) error      { return nil }
func (m *MockTLSConnection) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockTLSConnection) SetWriteDeadline(t time.Time) error { return nil }
