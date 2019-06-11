# A Tiny Pidfile Util for Golang

This package provides structure and helper functions to create and remove PID file. 
PIDFile is a file used to store the process ID of a running process.

## Feature

* Support on muti-system (Linux, macOS, Windows and FreeBSD)
* With all full tested

## Usage

To usage this package is simple, here is an example:

```golang
var pidFilePath = "/var/run/my.pid"
if pid, err := pidfile.New(pidFilePath); err != nil {
  log.Panic(err)
} else {
  fmt.Println(pid)
  defer pid.Remove()
}
```

## Feedback

If you have any suggest, sending me via email to `echo bWluZ2NoZW5nQG91dGxvb2suY29tCg== | base64 -D`, with huge thanks.

`- eof -`
