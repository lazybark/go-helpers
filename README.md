# go-helpers
Simple lib to make life with Go easier. Due to large amount of code cases to show, examples and test modules have been moved to [my site](https://lazybark.dev).

## What's inside


### No-pointer time

NPT is a struct of two fields: seconds and nanoseconds which can recreate default Go time.Time struct by calling time.Unix(sec, nsec) <br>
NPT holds no pointers (default time does - \*Location) and this makes NPT more memory-friendly due to reduced GC load. <br>
The main problem of pointers is that GC could chase them in memory in some cases and if you have an app that stores lots of time records, it's performance may be reduced due to that effect. So, if you don't need location data and your app uses time only internally, you may want to replace it with NPT.<br>
<br>
Code examples and test mod [here](https://lazybark.dev/go-helpers/#npt).

### Logger

LazyEvent - easy to use logger package that can be customized for almost any app. <br>
Main features:
* event-based logging to CLI & file
* events are objects that can be stored and modified
* events have levels and sources
* loggers are objects accessed via event processor
* colored CLI output
* in-memory events log, events are available at any time

Code examples and test mod [here](https://lazybark.dev/go-helpers/#lazy_event).


### Hasher

Hasher for files, strings and byte slices. Hash types: MD5, SHA1, SHA256, SHA512<br>
<br>
Code examples and test mod [here](https://lazybark.dev/go-helpers/#hasher).

### Text colors for console output

Simple ANSI escape sequences to format CLI-output.
**WARNING**
<br>
Colors will not work in standart Windows console. To get colors on Windows (instead of weird ANSI) use [Windows Terminal](https://docs.microsoft.com/en-us/windows/terminal/install) or any other app that supports ANSI escape codes.
<br>
Code examples and test mod [here](https://lazybark.dev/go-helpers/#clf).

### Semver
Semver is a simple package that provides tools to set and compare versions of anything in the world according to [Semantic versioning 2.0.0](https://semver.org/)
<br>
Code examples and test mod [here](https://lazybark.dev/go-helpers/#semver).

### Filesystem worker

Right now fsw is just a simple method to create files along with all dirs in path. Will be more functions here later.<br>
Second parameter, bool, will truncate the file in case true.<br>
```
f, err := fsw.MakePathToFile(path, true)
if err != nil {
    fmt.Println(err)
}
```


