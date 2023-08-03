package mock

type MockAddr struct {
	Ntwrk string // name of the network (for example, "tcp", "udp")
	Str   string // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

func (m *MockAddr) Network() string { return m.Ntwrk }
func (m *MockAddr) String() string  { return m.Str }
