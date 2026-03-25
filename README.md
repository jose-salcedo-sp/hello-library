# hello-library

Reusable Go module with greeting and random helper functions.

## Package

- `greetings`

## Exported functions

- `Hello1(name string) string`
- `RandomHello() string`
- `PersonalFunc() int`

## Example usage

```go
package main

import (
	"fmt"

	"github.com/jose-salcedo/hello-library/greetings"
)

func main() {
	fmt.Println(greetings.Hello1("Jose"))
	fmt.Println(greetings.RandomHello())
	fmt.Println(greetings.PersonalFunc())
}
```
