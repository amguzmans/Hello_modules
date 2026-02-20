# Hello_modules
## Installation

Excecute teh following command:

```bash
go get -u github.com/amguzmans/Hello_modules
```

## Usage
```go
import (
	"fmt"

	"github.com/amguzmans/Hello_modules"
)

func main() {
	message := hello_modules.Hello1("ana")
	fmt.Printf(message)
	fmt.Printf(hello_modules.RandomHello(), "Juan") 
}

```