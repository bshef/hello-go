# hello-go
Experimentations with GoLang (https://golang.org/)

## Index
- [Setup](#setup)
- [Build](#build)
- [Run](#run)

---

## Setup
### Download and Install the Go Programming Language
https://golang.org/dl/

#### The Go Distribution will be installed by default to:
- Linux and Mac OSX: `/usr/local/go`
- Windows: `C:\Go`

### Set Up Directory Structure
By convention, the Go directory structure should resemble the following:
```
.../Go/                                         # "GOROOT" (deprecated term)
    |
    -- /bin/                                    # Go executables stored here
    |
    -- /src/                                    # Go Source stored here
    |
    -- /pkg/                                    # Third party packages stored here
```

Create a separate working directory, which should resemble:
```
.../gowork/                                     # "GOPATH"
    |
    -- /bin/                                    # Executables stored here
    |
    -- /src/                                    # Projects stored here
    |   |
    |   -- /github.com/bshef/hello-go           # GitHub project
    |                           |
    |                           -- /hello/      # Project source files stored here
    |
    -- /pkg/                                    # Third party packages stored here
```

### Set Environment Variables
#### Linux and Mac OSX
```
export GOPATH=$HOME/gowork
export PATH=$PATH:$GOROOT/bin
```
#### Windows
```
SET GOPATH=C:\gowork
SET PATH=%PATH%;C:\gowork\bin
```
( Note: The MSI installer should have added `C:\Go\bin` to the PATH environment variable )

### For more information setting up GoLang:
https://golang.org/doc/install

## Build
Execute the command:
```
go install github.com/bshef/hello-go/hello
```
The above command will put an executable named `hello` (or `hello.exe`) inside
the `bin` directory under the GOPATH directory.

## Run
### Linux and Mac OSX
`$GOPATH/bin/hello`
### Windows
`%GOPATH%\bin\hello` or `hello` (if `%GOPATH%\bin` was added to the PATH)
