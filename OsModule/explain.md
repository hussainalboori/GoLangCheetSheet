1. Importing the OS module:
   ```go
   import "os"
   ```

2. Retrieving the current working directory:
   ```go
   wd, err := os.Getwd()
   if err != nil {
       // handle error
   }
   ```

3. Changing the current working directory:
   ```go
   err := os.Chdir("/path/to/directory")
   if err != nil {
       // handle error
   }
   ```

4. Creating a new directory:
   ```go
   err := os.Mkdir("/path/to/directory", os.ModePerm)
   if err != nil {
       // handle error
   }
   ```

5. Creating a new directory (including parent directories):
   ```go
   err := os.MkdirAll("/path/to/directory", os.ModePerm)
   if err != nil {
       // handle error
   }
   ```

6. Removing a directory:
   ```go
   err := os.Remove("/path/to/directory")
   if err != nil {
       // handle error
   }
   ```

7. Removing a directory (including its contents):
   ```go
   err := os.RemoveAll("/path/to/directory")
   if err != nil {
       // handle error
   }
   ```

8. Opening a file for reading:
   ```go
   file, err := os.Open("/path/to/file")
   if err != nil {
       // handle error
   }
   defer file.Close()
   ```

9. Opening a file for writing:
   ```go
   file, err := os.Create("/path/to/file")
   if err != nil {
       // handle error
   }
   defer file.Close()
   ```

10. Reading the contents of a file:
    ```go
    data := make([]byte, 1024)
    n, err := file.Read(data)
    if err != nil {
        // handle error
    }
    content := string(data[:n])
    ```

11. Writing data to a file:
    ```go
    data := []byte("Hello, world!")
    _, err := file.Write(data)
    if err != nil {
        // handle error
    }
    ```

12. Checking if a file or directory exists:
    ```go
    if _, err := os.Stat("/path/to/file_or_directory"); os.IsNotExist(err) {
        // file or directory does not exist
    } else {
        // file or directory exists
    }
    ```

13. Renaming a file or directory:
    ```go
    err := os.Rename("/path/to/old_name", "/path/to/new_name")
    if err != nil {
        // handle error
    }
    ```

14. Getting file information:
    ```go
    fileInfo, err := os.Stat("/path/to/file")
    if err != nil {
        // handle error
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size:", fileInfo.Size())
    fmt.Println("Is directory?", fileInfo.IsDir())
    fmt.Println("Permission mode:", fileInfo.Mode())
    ```

15. Listing files in a directory:
    ```go
    files, err := os.ReadDir("/path/to/directory")
    if err != nil {
        // handle error
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
    ```

Remember to handle errors appropriately in your actual code. This cheat sheet covers some common operations with the OS module in Go, but there are more functionalities available. You can refer to the official Go documentation for a complete list of functions and methods provided by the

 `os` package: https://golang.org/pkg/os/
