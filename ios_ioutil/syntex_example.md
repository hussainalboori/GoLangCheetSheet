Below is a cheat sheet for the `io/ioutil` package in Go, which provides utility functions for I/O operations.

Import statement:
```go
import "io/ioutil"
```

1. Reading from a file:
```go
func ReadFile(filename string) ([]byte, error)
```
Example:
```go
data, err := ioutil.ReadFile("file.txt")
if err != nil {
    // handle error
}
```

2. Writing to a file:
```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```
Example:
```go
data := []byte("Hello, world!")
err := ioutil.WriteFile("file.txt", data, 0644)
if err != nil {
    // handle error
}
```

3. Reading a directory:
```go
func ReadDir(dirname string) ([]os.FileInfo, error)
```
Example:
```go
files, err := ioutil.ReadDir("dir")
if err != nil {
    // handle error
}
for _, file := range files {
    fmt.Println(file.Name())
}
```

4. Temporarily creating a file:
```go
func TempFile(dir, prefix string) (*os.File, error)
```
Example:
```go
tempFile, err := ioutil.TempFile("", "example")
if err != nil {
    // handle error
}
defer os.Remove(tempFile.Name()) // Clean up the temporary file
```

5. Reading from a `io.Reader` into a byte slice:
```go
func ReadAll(r io.Reader) ([]byte, error)
```
Example:
```go
data, err := ioutil.ReadAll(reader)
if err != nil {
    // handle error
}
```

6. Writing from a byte slice to a `io.Writer`:
```go
func WriteAll(w io.Writer, data []byte) (int, error)
```
Example:
```go
n, err := ioutil.WriteAll(writer, data)
if err != nil {
    // handle error
}
```

Note: The `io/ioutil` package is deprecated as of Go 1.16, and it's recommended to use `io` and `os` packages directly for most cases.
