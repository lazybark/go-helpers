# go-helpers
Simple lib to make life with Go easier.
## What's inside

### Hasher

Hasher for files, strings and byte slices.
Hash types: MD5, SHA1, SHA256, SHA512
```
package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/hasher"
)

func main() {
	// Get all types of hash
	SHA256, err := hasher.HashFilePath("example_file", hasher.SHA256, 8192)
	if err != nil {
		fmt.Println(err)
	}
	MD5, err := hasher.HashFilePath("example_file", hasher.MD5, 8192)
	if err != nil {
		fmt.Println(err)
	}
	SHA1, err := hasher.HashFilePath("example_file", hasher.SHA1, 8192)
	if err != nil {
		fmt.Println(err)
	}
	SHA512, err := hasher.HashFilePath("example_file", hasher.SHA512, 8192)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File hashes:")
	fmt.Println(SHA256)
	fmt.Println(MD5)
	fmt.Println(SHA1)
	fmt.Println(SHA512)

	s := "Some string for you"
	fmt.Printf("String hashes ('%s'):\n", s)
	fmt.Println(hasher.HashString(s, hasher.SHA256))
	fmt.Println(hasher.HashString(s, hasher.MD5))
	fmt.Println(hasher.HashString(s, hasher.SHA1))
	fmt.Println(hasher.HashString(s, hasher.SHA512))

	fmt.Println("[]byte hashes:")
	b := []byte(s)
	fmt.Println(hasher.HashBytes(b, hasher.SHA256))
	fmt.Println(hasher.HashBytes(b, hasher.MD5))
	fmt.Println(hasher.HashBytes(b, hasher.SHA1))
	fmt.Println(hasher.HashBytes(b, hasher.SHA512))
}
```
Method (t \*HashType) CheckType() bool: returns 'true' if specified hash type exists in the lib. <br>
Method (t HashType) String() string: returns hash name or "illegal" if type is incorrect.

### No-pointer time

NPT is a struct of two fields: seconds and nanoseconds which can recreate default Go time.Time struct by calling time.Unix(sec, nsec) <br>
NPT holds no pointers (default time does - \*Location) and this makes NPT more memory-friendly due to reduced GC load. <br>
The main problem of pointers is that GC could chase them in memory in some cases and if you have an app that stores lots of time records, it's performance may be reduced due to that effect. So, if you don't need location data and your app uses time only internally, you may want to replace it with NPT.
```
package main

import (
	"fmt"
	"time"

	"github.com/lazybark/go-helpers/npt"
)

func main() {
	//Calling Now() will create a new NPT from current moment in time
	t := npt.Now()
	fmt.Println("Now it's:", t.Time())
	//ToNow will set internals of NPT to current moment
	time.Sleep(2 * time.Second)
	t.ToNow()
	fmt.Println("And now it's:", t.Time())
	//FromTime() will set NPT to specified time value
	t.FromTime(time.Now().Add(time.Hour))
	fmt.Println("And now it's:", t.Time())
	//Add() will add specified duration to NPT
	t.Add(time.Hour)
	fmt.Println("And now it's:", t.Time())
	//Time() will return time.Time object from NPT internals
	gt := t.Time()
	fmt.Println("And now it's:", gt)
}
```
### Text colors for console output
Simple ANSI escape sequences to format CLI-output:
```
package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/cli/clf"
)

func main() {
	fmt.Println("Let's make console pretty!")
	fmt.Println(clf.Red("Red text"))
	fmt.Println(clf.Green("Green text"))
	fmt.Println(clf.Yellow("Yellow text"))
	fmt.Println(clf.Blue("Blue text"))
	fmt.Println(clf.Magenta("Magenta text"))
	fmt.Println(clf.Cyan("Cyan text"))
	fmt.Println(clf.Gray("Gray text"))
	fmt.Println(clf.White("White text (ha-ha)"))
	fmt.Println()
	fmt.Println(clf.BBlack("Text with black background"))
	fmt.Println(clf.BRed("Text with red background"))
	fmt.Println(clf.BYellow("Text with yellow background"))
	fmt.Println(clf.BBlue("Text with blue background"))
	fmt.Println(clf.BMagenta("Text with magenta background"))
	fmt.Println(clf.BCyan("Text with cyan background"))
	fmt.Println(clf.BWhite("Text with white background"))
	fmt.Println()
	fmt.Println(clf.FBold("Text with bold formatting"))
	fmt.Println(clf.FUnderline("Text with underline formatting"))
	fmt.Println(clf.FBlink("Text with blinking formatting"))
	fmt.Println()
	fmt.Println(clf.Yellow(clf.BBlue("Text can " + clf.FReverse("be reversed") + clf.FUnReverse((" and dropped back, ")+"but both FReverse & FUnReverse need to have clf.Reset() at the end."+clf.Reset()))))
	fmt.Println()
	fmt.Println("Combining", clf.FBold(clf.BBlue("formats is even more fun!")))
	fmt.Println()
	fmt.Println("Manually", clf.BgMagenta.String()+"injected"+clf.Reset(), clf.CRed.String()+"formatting"+clf.Reset(), "also possible")
}
```
