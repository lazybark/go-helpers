package mock

import (
	"net"
	"time"
)

// MockTLSConnection mocks net.Conn interface
type MockTLSConnection struct {
	MWR MockWriteReader

	AskedToBeClosed bool

	LocAddr net.Addr
	RemAddr net.Addr
}

func (m *MockTLSConnection) Read(b []byte) (n int, err error) {
	return m.MWR.Read(b)
}

func (m *MockTLSConnection) Write(b []byte) (n int, err error) {
	return m.MWR.Write(b)
}

func (m *MockTLSConnection) Close() error {
	m.AskedToBeClosed = true

	return m.MWR.Close()
}

func (m *MockTLSConnection) LocalAddr() net.Addr                { return m.LocAddr }
func (m *MockTLSConnection) RemoteAddr() net.Addr               { return m.RemAddr }
func (m *MockTLSConnection) SetDeadline(t time.Time) error      { return nil }
func (m *MockTLSConnection) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockTLSConnection) SetWriteDeadline(t time.Time) error { return nil }
