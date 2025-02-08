 - Golang is compiled language 
 - Go tool can run file directyly without VM
 - Executables are different for different OS

 - What and where to use? - system apps to web apps - cloud

 - Object-Oriented? - No, but it has some features of OOP.

 - A lot things are missing in Go, but it is very fast and efficient still.
 - No try-catch, 
 - A lexer does a lot of work in Go.

- ----------------------------
-  ```go mod init <module-name>``` - to create a module in Go

-  ```go run <file-name>``` - to run a file in Go

- Go is case-sensitive in a lot of cases.

- ----------------------------
### GOPATH

The `GOPATH` environment variable is used to resolve import statements in Go. It is implemented by and documented in the `go/build` package.

- **Unix**: The value is a colon-separated string.
- **Windows**: The value is a semicolon-separated string.
- **Plan 9**: The value is a list.

If the environment variable is unset, `GOPATH` defaults to a subdirectory named `go` in the user's home directory (`$HOME/go` on Unix, `%USERPROFILE%\go` on Windows), unless that directory holds a Go distribution. Run `go env GOPATH` to see the current `GOPATH`.

For more information on setting a custom `GOPATH`, see [SettingGOPATH](https://golang.org/wiki/SettingGOPATH).

Each directory listed in `GOPATH` must have a prescribed structure:

- **src**: Holds source code. The path below `src` determines the import path or executable name.
- **pkg**: Holds installed package objects. Each target operating system and architecture pair has its own subdirectory of `pkg` (`pkg/GOOS_GOARCH`).
- **bin**: Holds compiled commands. Each command is named for its source directory, but only the final element, not the entire path. For example, the command with source in `DIR/src/foo/quux` is installed into `DIR/bin/quux`, not `DIR/bin/foo/quux`. The `foo/` prefix is stripped so that you can add `DIR/bin` to your `PATH` to access the installed commands. If the `GOBIN` environment variable is set, commands are installed to the directory it names instead of `DIR/bin`. `GOBIN` must be an absolute path.
