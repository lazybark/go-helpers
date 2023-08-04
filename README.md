# go-helpers
[![Test](https://github.com/lazybark/go-helpers/actions/workflows/test.yml/badge.svg)](https://github.com/lazybark/go-helpers/actions/workflows/test.yml)

go-helpers is a small and simple lib which i use for my everyday projects. It has packages to make life easier with using same solution for similar small tasks.

What's inside:
* No-pointer time (npt)


## What's inside


### No-pointer time

Perfect to use in loggers or other structs that just need to represent the second of some action.


Important: to keep the package simpler and faster, NPT does not provide exact precision up to nano. Max precision is up to a second. It's enough for most tasks, but if your app depends on deeper precision - time.Time is your choice.

Difference will look like that:
```
time.Time: 2023-08-03 19:33:13.0728246 +0000 UTC m=+0.005171401
npt.NPT: 2023-08-03 19:33:13 +0000 UTC
```

In term of nanoseconds difference will be:
```
time.Time: 1691080393072824600
npt.NPT: 1691080393000000000
``

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


