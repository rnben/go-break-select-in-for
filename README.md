# go-break-select-in-for

The Go linter go-break-select-in-for checks that break statement inside select statement inside for loop.

For example, in myFunc the break may want to exit the outer for loop, but it doesn't work as expected.

```go
func myFunc() {
    for {
        select {
        case <-ch:
            break // should be careful
        }
    }
}
```
