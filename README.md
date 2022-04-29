# go-helpers
Simple lib make life with go easier.
## What's inside
Hasher for files, strings and byte slices.
Hash types: MD5, SHA1, SHA256, SHA512
```
"github.com/lazybark/go-helpers/hasher"
HashFilePath(path string, ht HashType, bl int) (hashed string, err error)
HashFile(file *os.File, ht HashType, bl int) (hashed string, err error)
HashString(s string, ht HashType) string
HashBytes(b []byte, ht HashType) string
```
Method (t \*HashType) CheckType() bool: returns 'true' if specified hash type exists in the lib. <br>
Method (t HashType) String() string: returns hash name or "illegal" if type is incorrect.
