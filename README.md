# go_validated_types
Create types that self-validate in debug builds, without runtime overhead in release builds.


Compilation with `go build -tag debug my_own_package` will compile validation logic into the validated types.
Note that below example incurs 20-30x runtime overhead.
```go
package main

import "github.com/JonasMuehlmann/go_validated_types"

func main()  {
    // Create a new range int with a min of 0 and a max of 10 and current value 5
    i := NewRangeInt(5,0,10)
	i.Set(7)  // fine
	i.Set(11) // panic
}
```

Compilation without the `debug` build tag (`go build my_own_package`) will not compile validation logic into the validated types.
The resulting types are wrappers so thin, that the above code becomes binary equivalent to the following:

```go
package main

import "github.com/JonasMuehlmann/go_validated_types"

func main()  {
	i := 5
	i = 7
	i = 11
}
```