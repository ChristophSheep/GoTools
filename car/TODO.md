
# TODOS

## 2018-10-30

- vendor directories
- local packages
- model, view, view-model, controller

### Vendor directories

https://golang.org/cmd/go/#hdr-Vendor_Directories

```
/home/user/go/
    src/
        crash/
            bang/              (go code in package bang)
                b.go
        foo/                   (go code in package foo)
            f.go
            bar/               (go code in package bar)
                x.go
            vendor/
                crash/
                    bang/      (go code in package bang)
                        b.go
                baz/           (go code in package baz)
                    z.go
            quux/              (go code in package main)
                y.go
```