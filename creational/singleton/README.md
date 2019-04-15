# Singleton

## Definition
> Singleton pattern is a software design pattern that restricts the instantiation of a class to one "single" instance. This is useful when exactly one object is needed to coordinate actions across the system.

## Usages
- Connection to database.
- SSH connection to a server.
- Limit access.
- Limit number of calls.
- Permit lazy allocation and initialization.
- Hide the complexity of creating an object.

## Implementation
In this example, we don't allow to instanciate the `singleton` struct directly. Instead, we force to use the function `GetInstance()` to interact with the object.

```go
package singleton

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}   // <--- NOT THREAD SAFE
    }
    return instance
}
```
Note that this example works for a single threaded context, but it could fails on a concurrent context. The problem here is that is not thread safe, multiple go routines could evaluate the first check:
```go
if instance == nil {
``` 
and they would all create an instance of the singleton type and 
override each other.

## Example
The file `singleton.go` includes an implementation of the singleton pattern.

## Test
The file `singleton_test.go` tests the function `GetInstance`. Here is the output of running the testcases:

```bash
$ go test -race -v
=== RUN   TestGetInstance
--- PASS: TestGetInstance (0.00s)
PASS
ok      github.com/daniel-gil/go-design-patterns/creational/singleton       1.018s
```

## References
- [Singleton pattern](https://en.wikipedia.org/wiki/Singleton_pattern) definition from Wikipedia
- [Singleton Pattern in Go](http://marcio.io/2015/07/singleton-pattern-in-go/) by Marcio Castilho.