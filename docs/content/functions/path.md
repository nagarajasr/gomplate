---
title: path functions
menu:
  main:
    parent: functions
---

## `file.Stat`

Returns a [`os.FileInfo`](https://golang.org/pkg/os/#FileInfo) describing
the named path. 
Essentially a wrapper for Go's [`os.Stat`](https://golang.org/pkg/os/#Stat) function.

### Usage
```go
file.Stat path
```

### Examples

```console
$ echo "hello world" > /tmp/foo
$ gomplate -i '{{ $s := file.Stat "/tmp/foo" }}{{ $s.Mode }} {{ $s.Size }} {{ $s.Name }}'
-rw-r--r-- 12 foo
```
