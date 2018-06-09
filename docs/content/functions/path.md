---
title: path functions
menu:
  main:
    parent: functions
---

The path functions are split into 2 namespaces: 
- `path`, which is useful for manipulating slash-based (`/`) paths, such as in URLs
- `filepath`, which should be used for local filesystem paths, especially when Windows paths may be involved

These functions are wrappers for Go's [`path`](https://golang.org/pkg/path/) and [`path/filepath`](https://golang.org/pkg/path/filepath/) packages.

## `path.Base`

Returns the last element of path. Trailing slashes are removed before extracting the last element. If the path is empty, Base returns `.`. If the path consists entirely of slashes, Base returns `/`.

A wrapper for Go's [`path.Base`](https://golang.org/pkg/path/#Base) function.

### Usage
```go
path.Base path
```

### Examples

```console
$ echo "hello world" > /tmp/foo
$ gomplate -i '{{ $s := file.Stat "/tmp/foo" }}{{ $s.Mode }} {{ $s.Size }} {{ $s.Name }}'
-rw-r--r-- 12 foo
```
