# go-helpers
[![Test](https://github.com/lazybark/go-helpers/actions/workflows/test.yml/badge.svg)](https://github.com/lazybark/go-helpers/actions/workflows/test.yml)

go-helpers is a small and simple lib which i use for my everyday projects. It has packages to make life easier with using same solution for similar small tasks.

What's inside:
* No-pointer time (npt)
* Security
* Generators (bytes, strings)
* Mock
* CSV worker
* CSV comparer as CLI app
* Google API helper
* Converters
* Hasher (bytes, strings, files)
* CLI formatter
* Semver
* Filesystem worker




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
