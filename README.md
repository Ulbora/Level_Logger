[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/Level_Logger)](https://goreportcard.com/report/github.com/Ulbora/Level_Logger)

Level Logger
============

A three level logger for Golang. It has the following levels: debug, info, and error.

The logging level can be set by setting the LogLevel varible as shown:

### Error
Error logging is alwas on and cannot be turned off.

``` var l Logger
	i, e := strconv.Atoi("1q")
	l.Error(e)
```
#### Output
```
2020/02/02 15:57:51 ERROR:  strconv.Atoi: parsing "1q": invalid syntax

```

### Info
```
    var l Logger
	l.LogLevel = InfoLevel
	i, e := strconv.Atoi("1q")
	l.Info(e)

```

#### Output
```
2020/02/02 15:57:51 INFO:  strconv.Atoi: parsing "1q": invalid syntax

```

### Debug
```
    var l Logger
	l.LogLevel = DebugLevel
	i, e := strconv.Atoi("1q")
	l.Debug(e)
	
```

#### Output
```
2020/02/02 15:57:51 DEBUG:  strconv.Atoi: parsing "1q": invalid syntax

```

### Log All
```
    var l Logger
	l.LogLevel = AllLevel
	i, e := strconv.Atoi("1q")
	l.Info(e)
	
```

#### Output
```
2020/02/02 15:57:51 INFO:  strconv.Atoi: parsing "1q": invalid syntax

```
