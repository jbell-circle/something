# Something

```
go test ./...
```

If we have an interface, how do we setup a testsuite to validate its many implementations?
This repo demos a `_tester` suite that can be imported to accomplish this task.

- `something.go` defines a simple interface.
- `something_tester.go` defines the testsuite that all implementations must pass to be valid.
- `memory` and `sync` different implementations with `_test.go` files that use the testers.
