package mock

type MockFile struct {
	MWR MockWriteReader
}

func (m *MockFile) Write(b []byte) (n int, err error) {
	return m.MWR.Write(b)
}

func (m *MockFile) WriteString(s string) (n int, err error) {
	return m.MWR.WriteString(s)
}

func (m *MockFile) Close() error {
	return m.MWR.Close()
}

func (m *MockFile) Read(b []byte) (n int, err error) {
	return m.MWR.Read(b)
}

func (m *MockFile) SetLastRead(n int) {
	m.MWR.SetLastRead(n)
}
