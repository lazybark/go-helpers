# go-helpers
[![Test](https://github.com/lazybark/go-helpers/actions/workflows/test.yml/badge.svg)](https://github.com/lazybark/go-helpers/actions/workflows/test.yml)

go-helpers is a small and simple lib which i use for my everyday projects. It has packages to make life easier with using same solution for similar small tasks.

What's inside:
* No-pointer time (npt)
* Security (sec)
* Generators (gen)
* mock
* CSV worker (csvw)
* CSV comparer as CLI app (cmd/csv-comparer)
* Google API helper (gapi)
* Converters (conv)
* Hasher (hasher)
* CLI formatter (cli/clf)
* Semantic version (semver)
* Filesystem worker (fsw)

## No-pointer time

NPT is a time-holding object that does not have timezone pointer (default time does - *Location). It's more memory-effective than default Go time.Time and can be used in apps that store time in memory only and do not serialize it. It can be a logger package that works with huge amount of events or an action buffer that doesn't need extra precision.

Absence of pointers makes NPT more memory-friendly due to reduced GC load. It's the only reason of NPT's existence.

(The main problem of pointers is that GC could chase them in memory in some cases and if you have an app that stores lots of time records, it's performance may be reduced due to that effect. So, if you don't need location data and your app uses time only internally, you may want to replace it with NPT)

Important: to keep the package simpler and faster, NPT does not provide exact precision up to nano. Max precision is up to a second. It's enough for most tasks, but if your app depends on deeper precision - time.Time is still your choice.

Difference between outputs will look like that:
```
time.Time: 2023-08-03 16:33:13.0728246 +0000 UTC m=+0.005171401
npt.NPT: 2023-08-03 16:33:13 +0000 UTC
```

In term of nanoseconds difference will be:
```
time.Time: 1691080393072824600
npt.NPT: 1691080393000000000
```

### How to use NPT

Examples can be found in [cmd/npt-examples](https://github.com/lazybark/go-helpers/blob/main/cmd/npt-examples/npt.go)

Calling `Now()` will create a new NPT from current moment in time
```
t := npt.Now()
fmt.Println("Now it's:", t.Time())
```
`ToNow()` will set internals of NPT to current moment
`Time()` will return time.Time object from NPT internals
```
time.Sleep(2 * time.Second)
t.ToNow()
fmt.Println("Now it's:", t.Time())
```
`FromTime()` will set NPT to specified time value
```
t.FromTime(time.Now().Add(time.Hour))
fmt.Println("And now it's:", t.Time())
```
`Add()` will add specified duration to NPT
```
t.Add(time.Hour)
fmt.Println("And now it's:", t.Time())
```

## sec package

sec has functions to hash/compare passwords and generate cryptographically secure random strings using rand.Reader.

`HashAndSaltPasswordString(pwd string)` & `HashAndSaltPassword(pwd []byte)` will return password hash or error.

`ComparePasswords(hashedPwd string, plainPwd string)` & `ComparePasswordBytes(byteHash []byte, plainPwdBytes []byte)` will return 'true' if password matches hash or error.

`GenerateRandomString(n int)` will return string of n length filled with random symbols from english alphabet and numbers. So it's non-extended ASCII printable characters only (excluding special symbols) and number of symbols in the string will be equal to it's length in bytes.

But if you need to use specific character set, you can call `GenerateRandomStringFromSet(n int, charSet []byte)` providing your own set of one of predefined in [gen/charaterSets.go](https://github.com/lazybark/go-helpers/blob/main/gen/charaterSets.go). But it's important to understand that some (most) languages use symbols that longer than 1 byte. So resulting string may not be readable by ~~pathetic biological creatures like us~~ humans.

`åäö` - this string, for example, takes 6 bytes, not 3. So if you try to take random bytes from here, you will possibly get something like `�å`

Use tools [like this one](https://mothereff.in/byte-counter) to check 

## gen package

gen has methods to generate random strings & bytes that work the same way as in `sec` package above, but are not cryptographically secure. They share way of generating via rand.Reader, but no checks are made during the process. So result may be insecure.

Rules same to `sec` package above apply to resulting data sets.

* `GenerateRandomString(n int)` & `GenerateRandomBytes(n int)` uses english letters + digits
* `GenerateRandomStringFromSet(n int, charSet []byte)` & `GenerateRandomBytesFromSet(n int, charSet []byte)` use any character set you provide
* `GenerateRandomStringSet(lens []int)` & `GenerateRandomBytesSet(lens []int)` will provide a slice of random strings of english letters + digits
* `GenerateRandomStringSetFromSet(lens []int, charSet []byte)` & `GenerateRandomBytesSetFromSet(lens []int, charSet []byte)` use any character set you provide

## mock package

mock is a set of structs that implement specific interfaces and can be used in tests.

### MockWriteReader

MockWriteReader has methods to read/write into exported Bytes field. Field `ReturnEOF` will make `Read()` method to return `io.EOF` on any call. Field `DontReturEOFEver` will stop `Read()` from returnin `io.EOF` even if `Bytes` buffer was fully read.

You can reuse same MockWriteReader to read by simply calling `SetLastRead(int)` - it will set internal last read byte number to the one provided.

Full list of methods:

* `Write(b []byte) (n int, err error)`
* `WriteString(s string) (n int, err error)`
* `Close() error`
* `Read(b []byte) (n int, err error)`
* `SetLastRead(n int)`

Example declaration:

```
someText := "This is some file here"
wreader := mock.MockWriteReader{
  Bytes: []byte(someText),
}
```

### MockFile

MockFile uses MockWriteReader as `MWR` field to mock read/write operations.

### MockHTTPWriter

MockHTTPWriter is meant to implement `http.ResponseWriter` interface. It can be useful in various test cases with RESTful API methods that do not return any value to external function but write output directly to HTTP client.

Methods:

* `Header() http.Header` - returns `http.Header` in case it was set before or just nil map in other cases
* `Write(b []byte) (int, error)` - adds to `Data` field
* `WriteHeader(statusCode int)` - sets `StatusCode` field
* `AssertAndFlush(t *testing.T, assertWith interface{})` - uses `assert.Equal()` to check if current buffer data equals to given example and then cleans Data field
* `Flush()` - just cleans Data field

### MockAddr

MockAddr mocks `net.Addr`. Methods:

* `Network() string`
* `String() string`

### MockTLSConnection

MockTLSConnection implements `net.Conn` interface and uses MockWriteReader to mock connection read/write ops. Methods:

* `Read(b []byte) (n int, err error)`
* `Write(b []byte) (n int, err error)`
* `Close() error` - sets AskedToBeClosed field to `true`
* `LocalAddr() net.Addr` - returns LocAddr field
* `RemoteAddr() net.Addr` - returns RemAddr field
* `SetDeadline(t time.Time) error` - always nil
* `SetReadDeadline(t time.Time) error` - always nil
* `SetWriteDeadline(t time.Time) error` - always nil

## csvw package

CSV worker is the package to perform operations over CSV files.

### CSVBuilder

CSVBuilder uses strings.Builder to create CSV strings and can write them to a file or mock struct. Use NewCSVBuilder(separator string) to create ready to use builder.

Methods:

* `OpenFile(path string, truncate bool) (err error)` - opens a file to APPEND
* `UseFile(f fsw.IFileWriter)` - uses already opened file
* `Close() error` - closes file that was opened by builder or set by
* `AddCell(str ...string) (err error)` - adds new cell to current string (with separator at the end)
* `AddLine(str string) (err error)` - adds whole line to current buffer (with separator and '\n' at the end)
* `NewLine() (err error)` - adds line break to current string
* `Reset()` - cleans current string
* `String() string` - returns current string data
* `WriteBuffer() (int, error)` - writes current byte buffer into opened file and cleans the buffer right after
* `Write(bts []byte) (int, error)` - writes bytes directly into file (no line break at the end)
* `WriteString(s string) (int, error)` - writes s directly into file (no line break at the end)
* `WriteLine(bts []byte) (int, error)` - writes bytes directly into file and adds line break after last byte
* `WriteLineString(s string) (int, error)` - writes s directly into file and adds line break after last byte
* `WriteInto(w io.Writer) (int, error)`  - writes buffer into w and resets the buffer

### CSV comparer

Function CompareCSVs(fOne, fTwo fsw.IFileReader, pathOne, pathTwo, dividerOne, dividerTwo, keyCol string, compareCols ...string) takes fOne as base csv dataset and fTwo as changed dataset. Then compares column by column (compareCols) using keyCol as line primary ID. Generates a Compared struct that can write results into file if needed.

Different is the model that holds data about exactly two rows that differ (compared by keyCol). RowOne with data from fOne and RowTwo from fTwo. Cols here stores list of columns that have different data.

Model Compared holds data of two compared csv files, including statistic. Methods:

* `WriteDifferent(file fsw.IFileWriter) error` - writes into file full list of rows that differ from first to second file
* `WriteDeleted(file fsw.IFileWriter) error` - writes into file full list of deleted rows (that exist in first file, but not in second)
* `DifferentRows() []Different` - returns list of rows that differ (compared by keyCol)
* `DeletedRows() []map[string]string` - returns list of deleted rows (exist in first file, but not in second)
* `TotalRowsInFirstFile() int` - number of rows in first file
* `TotalRowsInSecondFile() int` - number of rows in second file
* `DifferentRowsCount() int` - number of rows that differ from document to document, but have same keyCol value
* `SameRowsCount() int` - number of rows that are same in both documents
* `DeletedRowsCount() int` - number of rows that exist in first document, but not in second
* `DifferentFieldsStat() map[string]int` - list of column names with number of how many rows have different value in each column

Working example can found at [cmd/csv-comparer-examples.go](https://github.com/lazybark/go-helpers/tree/main/cmd/csv-comparer-examples).

Live command-line tool - at [cmd/csv-comparer](https://github.com/lazybark/go-helpers/tree/main/cmd/csv-comparer) (requires `go build`).
