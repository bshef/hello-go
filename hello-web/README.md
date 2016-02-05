# hello-web Go Web Server

## Index
- [Setup](#setup)
- [Build](#build)
- [Run](#run)
- [Test](#test)

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
.../Go/                                             # "GOROOT"
    |
    -- /bin/                                        # Executables stored here
    |
    -- /src/                                        # Projects stored here
    |   |
    |   -- /github.com/golang/hello-go              # Workspace and GitHub project
    |                           |
    |                           -- /hello-web/      # Project source files stored here
    |
    -- /pkg/                                        # Third party packages stored here
```

### Set Environment Variables
#### Linux and Mac OSX
```
export GOROOT=$HOME/go
export GOPATH=$HOME/go/src/github.com/golang/hello-go
export PATH=$PATH:$GOROOT/bin
```
#### Windows
```
SET GOROOT=C:\Go
SET GOPATH=C:\Go\src\github.com\golang\hello-go
```
( Note: The MSI installer should have added `C:\Go\bin` to the PATH environment variable )

### For more information setting up GoLang:
https://golang.org/doc/install

## Build
In a command window, navigate to the `.../Go/src` directory, then execute the command:
```
go install github.com/golang/hello-go/hello-web
```
The above command will put an executable named `hello-web` (or `hello-web.exe`) inside
the `bin` directory under the `GOROOT`.

## Run
### Linux and Mac OSX
`$GOROOT/bin/hello-web`
### Windows
`%GOROOT%\bin\hello-web`

## Test
```
curl -i localhost:8000
curl -i localhost:8000/health
```
