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


- ----------------------------
- The formal grammer uses semi-colons to terminate statements, but the Go compiler inserts them automatically at the end of each line.

### Lexer
 Every programming language has a lexer. It's job is simply to check the syntax of the code or grammer of the language. 

 Before pre-compilation, the lexer checks the syntax of the code. 

 So, the semicolons are automatically inserted by the lexer in Go. 


### Types in Go

Case-sensitive
Variable type should be known in advance - whether it is public or private, int or string, etc.

Almost everything in Go is a type.

Some types in Go:
- string
- bool
- Integer : uint8, uint64, int8, int64, unintptr
- Floating: float32, float64
- Complex

- ----------------------------

Advance Types in Go:
- Array
- Slice
- Map
- Struct
- Pointer
- Function
- Channel, etc.

- ----------------------------

## Time Package in Go

- Time package is used to work with time in Go.
- It is used to get the current time, format the time, etc.

- Format() method is used to format the time in Go. But the format is different from the other languages.
We use 01-02-2006 15:04:05 Monday to format the time in Go.
Yes, it is weird but it is the way to format the time in Go.

## Builds in Go

- `go env` - to see the environment variables in Go.

`GOOS="windows" go build `-- finds main.go and creates an executable file for windows.

`GOOS="linux" go build` -- finds main.go and creates an executable file for linux. and so on.


## Memory Management in Go

- Memory allocation and deallocation happens automatically in Go.

- We usually use two methods to allocate memory in Go:
    - `new()` - to allocate memory but not initialize it., we get a memory address. A zeroed storage is returned.
    In zeroed storage, you cannot put any data.
    - `make()` - to allocate memory and also initialize it. We get a memory address. Non-zeroed storage is returned.
    In non-zeroed storage, we can put data.

- Automatic garbage collection in Go. Anything out of scope or nil is garbage collected in Go.
- The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely.


## Defer Statement in Go

- `defer` statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
- Deferred functions are executed in LIFO order.


## Web Requests in Go

- `net/http` package is used to make web requests in Go. It is the fastest way to make web requests in Go.


# GO MOD

- ``` go mod init github.com/ayyush08/mymods``` - to initialize a module in Go. It tells that the entire code was dependent on this version of the module. and it creates a go.mod file. For versioning, Go follows the semantic versioning(semver) just like npm.

- go toolchain is the way how you pull all dependencies  of a repo. It is a way to manage dependencies in Go.

- go sum file is created to check the integrity of the dependencies. It is a checksum file to prevent the code from being tampered with or threat attacks.

- ```go mod tidy``` - to clean up the dependencies in Go. It is an expensive operation. after runnning this the '//indirect' is gone from the go.mod file as tidy removes the dependencies that are not used in the code.

- ```gorilla/mux``` - is a package in Go to create a router in Go. It is a third-party package in Go.

- ```go mod verify``` - to verify the dependencies in Go. It checks the integrity of the dependencies by going to go.sum file and checking the checksums of the dependencies.

- ```go list``` - to list the dependencies in present module.
- ```go list -m all``` - to list all the dependencies in current module.

- ```go list all``` - to list all the dependencies in the whole ecosystem.

- ```go mod why <module-name>``` - to check why a particular module is being used in the code.

- ```go mod graph``` - to see the graph of the dependencies in the code.

- ```go mod edit -go 1.14``` - to change the version of Go in the code.

- ```go mod edit -module <module-name>``` - to change the module name in the code.

- ```go mod vendor``` - to create a vendor folder i.e, the cache of the dependencies stored externally. It is used to bring the all files of the dependencies in the project. But we can't say go run main.go, we have to say go run -mod=vendor main.go to run the code. (Generally, it is not recommended to use vendor folder in Go.)


- --------------------------------------------

## Writing APIs in Go

When creating any controller , the parameters for the function are always 2 things: A writer and a reader.
Writer(w http.ResponseWriter) is used to write the response and Reader(r *http.Request) is used to read the request coming in.

### Note 
While designing APIs, it is suggested to have one main.go file and then create a separate folder for the controllers, models, and routes inside the root directory as having all the files in the root directory can be messy or create problems in the future.

- --------------------------------------------

## Concurrency vs Parallelism

- Concurrency is the ability to run multiple tasks at the same time. But it is not necessary that they run at the same time parallelly. You can switch between the tasks and run them at the same time. It is like multitasking.

- Parallelism is the ability to run multiple tasks at the same time parallelly. It is like running multiple tasks at the same time.


## Goroutines in Go

- Goroutines are how we achieve this parallelism in Go. They are lightweight threads managed by the Go runtime. They are not OS threads, they are managed by the Go runtime.
Normal OS threads have fixed stack of 1mb while goroutines have a stack of 2kb. So, we can create a lot of goroutines in Go which makes running task faster.

- "Do not communicate by sharing memory; instead, share memory by communicating." - Rob Pike (Go creator)
- `go` keyword is used to create a goroutine in Go.

- --------------------------------------------

## Question: If the threads are managed by the Go runtime, then who is managing the lock on to a memory location?

- Let's say we have 5 different go routines simultaneously trying to access the same memory location. Then, how does Go manage the lock on to that memory location? We need to have some responsibility to manage the lock on to that memory location.

- Hence, we use Mutex in Go to manage the lock on to a memory location. Mutex is a lock that is used to lock the memory location so that only one goroutine can access that memory location at a time. 

- Mutex allows anybody with read access to the memory location but only one person with write access to the memory location.