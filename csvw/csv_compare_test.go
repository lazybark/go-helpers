package csvw

import (
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCSVString struct {
	keyColValue  string
	stringValues []string
}

const (
	csvDivider = ";"
)

var (
	csvCols1       = []string{"id", "name", "email", "some", "another"}
	csvCols2       = []string{"id", "name", "email", "extra"}
	csvColsCompare = []string{"id", "name", "email"}

	//deletedString1 = []string{"4", "user 4", "user4@gmail.com", "some4", "another3"}

	csvLines1 = []testCSVString{
		{
			keyColValue:  "1",
			stringValues: []string{"1", "user 1", "user1@gmail.com", "some1", "another1"},
		},
		{
			keyColValue:  "2",
			stringValues: []string{"2", "user 2", "user2@gmail.com", "some2", "another2"},
		},
		{
			keyColValue:  "3",
			stringValues: []string{"3", "user 3", "user3@gmail.com", "some3", "another3"},
		},
		{
			keyColValue:  "4",
			stringValues: []string{"4", "user 4", "user4@gmail.com", "some4", "another3"},
		},
		{
			keyColValue:  "yy",
			stringValues: []string{"yy", "dfsdfsdf", "uyjhgrtyhj", "someY", "anotherY"},
		},
	}
	csvLines2 = []testCSVString{
		{
			keyColValue:  "1",
			stringValues: []string{"1", "user 1", "user1@gmail.com", "extra1"},
		},
		{
			keyColValue:  "2",
			stringValues: []string{"2", "user 2", "user2@gmail.com1", "extra2"},
		},
		{
			keyColValue:  "4",
			stringValues: []string{"4", "user 4", "user4@gmail.com1", "some4", "another3"},
		},
		{
			keyColValue:  "yy",
			stringValues: []string{"yy", "dfsdfsdf", "uyjhgrtyhj", "extraY"},
		},
	}
)

//There are some static cases here and the test could be made much more complex with much more test cases.
//I will come back when i have free time for this task. Right now it's not crucial as i don't intent to change any
//of csvw in nearest futue

func TestCSVCompareAndWrite(t *testing.T) {
	file1 := &mock.MockWriteReader{}
	file2 := &mock.MockWriteReader{}

	//Head / cols
	for _, v := range csvCols1 {
		file1.Bytes = append(file1.Bytes, []byte(v+csvDivider)...)
	}
	file1.Bytes = append(file1.Bytes, '\n')

	for _, v := range csvCols2 {
		file2.Bytes = append(file2.Bytes, []byte(v+csvDivider)...)
	}
	file2.Bytes = append(file2.Bytes, '\n')

	for _, v := range csvLines1 {
		for _, svv := range v.stringValues {
			file1.Bytes = append(file1.Bytes, []byte(svv+csvDivider)...)
		}
		file1.Bytes = append(file1.Bytes, '\n')
	}
	for _, v := range csvLines2 {
		for _, svv := range v.stringValues {
			file2.Bytes = append(file2.Bytes, []byte(svv+csvDivider)...)
		}
		file2.Bytes = append(file2.Bytes, '\n')
	}

	name1 := "file1"
	name2 := "file2"
	c, err := CompareCSVs(file1, file2, name1, name2, csvDivider, csvDivider, csvCols1[0], csvColsCompare...) //STATIC here - change if head changes
	require.NoError(t, err)

	for key, cc := range csvColsCompare {
		assert.Equal(t, cc, c.compareCols[key])
	}
	assert.Equal(t, name1, c.one)
	assert.Equal(t, name2, c.two)
	assert.Equal(t, csvDivider, c.Divider)
	assert.Equal(t, csvCols1[0], c.keyCol)
	assert.Equal(t, len(csvLines1), c.totalOne)
	assert.Equal(t, len(csvLines2), c.totalTwo)

	assert.Equal(t, 2, c.diff)         //STATIC here - change if assets change
	assert.Equal(t, 2, c.same)         //STATIC here - change if assets change
	assert.Equal(t, 1, len(c.deleted)) //STATIC here - change if assets change

	assert.Equal(t, "3", c.deleted[0]["id"]) //STATIC here - change if assets change

	_, ok := c.diffFields["email"] //STATIC here - change if assets change
	assert.True(t, ok)
	_, ok = c.diffFields["id"] //STATIC here - change if assets change
	assert.False(t, ok)
	_, ok = c.diffFields["name"] //STATIC here - change if assets change
	assert.False(t, ok)

	stat := c.DifferentFieldsStat() //STATIC here - change if assets change
	diffCount, ok := stat["email"]
	assert.True(t, ok)
	assert.Equal(t, 2, diffCount)

	fileDiff := &mock.MockWriteReader{}
	err = c.WriteDifferent(fileDiff)
	require.NoError(t, err)

	//STATIC here - change if assets change
	assert.Equal(t,
		"doc;id;name;email;email_d;\n//;\nfile1;2;user 2;user2@gmail.com;TRUE;\nfile2;2;user 2;user2@gmail.com1;TRUE;\n//;\nfile1;4;user 4;user4@gmail.com;TRUE;\nfile2;4;user 4;user4@gmail.com1;TRUE;\n",
		string(fileDiff.Bytes))

	fileDel := &mock.MockWriteReader{}
	err = c.WriteDeleted(fileDel)
	require.NoError(t, err)
	assert.Equal(t, "id;id;name;email;\n3;3;user 3;user3@gmail.com;\n", string(fileDel.Bytes))
}
