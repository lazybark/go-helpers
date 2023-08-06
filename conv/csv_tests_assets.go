package conv

type testCSVString struct {
	keyColValue  string
	stringValues []string
}

const (
	csvDivider = ";"
)

var (
	csvColsBench = []string{"id", "name", "email", "some1", "some2", "some3", "some4", "some5", "some6", "some7"}
	csvCols      = []string{"id", "name", "email"}
	csvLines     = []testCSVString{
		{
			keyColValue:  "1",
			stringValues: []string{"1", "user 1", "user1@gmail.com"},
		},
		{
			keyColValue:  "2",
			stringValues: []string{"2", "user 2", "user2@gmail.com"},
		},
		{
			keyColValue:  "3",
			stringValues: []string{"3", "user 3", "user3@gmail.com"},
		},
		{
			keyColValue:  "yy",
			stringValues: []string{"yy", "dfsdfsdf", "uyjhgrtyhj"},
		},
	}
)
