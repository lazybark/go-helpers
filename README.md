# go-helpers
Simple lib to make life with Go easier. Due to large amount of code cases to show, examples and test modules have been moved to [my site](https://lazybark.dev).

## What's inside

### Hasher

Hasher for files, strings and byte slices. Hash types: MD5, SHA1, SHA256, SHA512<br>
Code examples and test mod [here](https://lazybark.dev/go-helpers/#hasher).


### No-pointer time

NPT is a struct of two fields: seconds and nanoseconds which can recreate default Go time.Time struct by calling time.Unix(sec, nsec) <br>
NPT holds no pointers (default time does - \*Location) and this makes NPT more memory-friendly due to reduced GC load. <br>
The main problem of pointers is that GC could chase them in memory in some cases and if you have an app that stores lots of time records, it's performance may be reduced due to that effect. So, if you don't need location data and your app uses time only internally, you may want to replace it with NPT.
Code examples and test mod [here](https://lazybark.dev/go-helpers/#npt).
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
It's also possible to use `clf.FNone` as default in external methods that accept or pass on clf formatting directives.<br>
It will return just an empty string.<br>
**WARNING**
<br>
Colors will not work in standart Windows console. To get colors on Windows (instead of weird ANSI) use [Windows Terminal](https://docs.microsoft.com/en-us/windows/terminal/install) or any other app that supports ANSI escape codes.

### Logger

LazyEvent - easy to use logger package that can be customized for almost any app. <br>
Main features:
* event-based logging
* events are objects that can be stored and modified
* events have levels and sources
* loggers are objects accessed via one event processor
* colored CLI output
* in-memory events log, N last events are available at any time
<br>
Event logging is quite simple.
Create new processor with desired number of events to keep in chain.

```
p := le.NewV1(8)
```

Add loggers for CLI or console: <br>

```
err := p.NewConsole()
if err != nil {
	log.Fatal(err)
}
err = p.NewFile("log.txt", true)
if err != nil {
	log.Fatal(err)
}
```
You can use event sources that will be added in logs between event type and event text.<br>
To add:

```
srcE := p.Source("EXTRA", "cyan", "[", "]")
srcS := p.Source("SOMESTUFF", "red", "[", "]")
```

Events can be created in different ways:
* from any event source (this source will be set as default)
* from processor (all event parameters will be defaiult and source will be empty)
* from event type (only type will be set, all rest - defaults)




