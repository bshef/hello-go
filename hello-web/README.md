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
    |                           -- /hello-web/  # Project source files stored here
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
go install github.com/bshef/hello-go/hello-web
```
The above command will put an executable named `hello-web` (or `hello-web.exe`)
inside the `bin` directory under the GOPATH directory.

## Run
### Linux and Mac OSX
`$GOPATH/bin/hello-web`
### Windows
`%GOPATH%\bin\hello-web` or `hello` (if `%GOPATH%\bin` was added to the PATH)

## Test
Visit `http://localhost:8000` in a modern browser to see static content served.

Visit `http://localhost:8000/view/test` in a modern browser to see templated
content being served.

Enter the following commands to see a simple API in action:
```
curl -i localhost:8000
curl -i localhost:8000/health
```
