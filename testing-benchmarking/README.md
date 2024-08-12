
## Testing

### testing.T:

use for unit test.

the function name should be started by `Test`

```bash
package mypackage

import "testing"

func TestAdd(t *testing.T) {
    expected := 5
    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}
```
### testing.B:

use for benchmarks

the function name should be started by `Benchmark`

```bash
package mypackage

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

### testing.F:

use for fuzz tests

the function name should be started by `Fuzz`

```bash
package mypackage

import "testing"

func FuzzAdd(f *testing.F) {
    f.Add(1, 2) // Add seed inputs
    f.Fuzz(func(t *testing.T, a int, b int) {
        _ = Add(a, b)
    })
}
```

### testing.M:

for manage test suites with setup and tear down

```bash
package mypackage

import (
    "fmt"
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    // Setup code
    fmt.Println("Setup")

    code := m.Run() // Run all tests

    // Teardown code
    fmt.Println("Teardown")

    os.Exit(code)
}
```

### testing.TB:

an interface for T and B

```bash
package mypackage

import "testing"

func assertEqual(tb testing.TB, result, expected int) {
    if result != expected {
        tb.Errorf("expected %d but got %d", expected, result)
    }
}

func TestAdd(t *testing.T) {
    assertEqual(t, Add(2, 3), 5)
}
```

### testing.PB:

for parallel benchmarks

```bash
package mypackage

import "testing"

func BenchmarkAddParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Add(2, 3)
        }
    })
}
``
