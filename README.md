# go-helpers
[![Test](https://github.com/lazybark/go-helpers/actions/workflows/test.yml/badge.svg)](https://github.com/lazybark/go-helpers/actions/workflows/test.yml)

go-helpers is a small and simple lib which i use for my everyday projects. It has packages to make life easier with using same solution for similar small tasks.

What's inside:
* No-pointer time (npt)





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


