
## switch

There is no automatic fall through, but cases can be presented in comma-separated lists.

```
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

## The blank identifier

To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the imported package. Similarly, assigning the unused variable fd to the blank identifier will silence the unused variable error. This version of the program does compile.

```
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

## Profiling

Read https://blog.golang.org/profiling-go-programs

## offline doc

If you're ever stuck without internet access, you can get the documentation running locally via:

```
godoc -http=:6060
```

and pointing your browser to http://localhost:6060

## random bytes

ref https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

## range map is random

Since the release of Go 1.0, the runtime has randomized map iteration order. Programmers had begun to rely on the stable iteration order of early versions of Go, which varied between implementations, leading to portability bugs.

1.0版本后，go团队有意将map遍历随机化。

The old language specification did not define the order of iteration for maps, and in practice it differed across hardware platforms. This caused tests that iterated over maps to be fragile and non-portable, with the unpleasant property that a test might always pass on one machine but break on another.

随机化原因：在不同硬件平台上，map遍历的顺序会不一致，可能会导致单测失败。因此刻意使其随机化，让程序员不依赖其顺序。