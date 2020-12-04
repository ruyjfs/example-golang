# example-golang

Golang, GraphQL, Resful, Auth JWT example and much more.

##

- [Step By Step](#Step-By-Step)

## Step By Step

All steps necessary to make this example project.

### Create a mod init

All libs in Golang have a module name.
This is for a package manage for Golang, like a npm/yarm for nodeJS, gradlew/mavem for Java, compose for PHP, mix for Elixir and many more.

```bash
go mod init github.com/ruyjfs/example-golang
```

> See: go.mod

### Create a file main.go - The Hello World

```bash
 echo '' >> main.go
```

```go
package main

import (
	"log"
)

func main() {
	log.Printf("Hello World!")
}
```

Run to see you first hello world

```bash
go run main.go
```

### Install GraphQL

```bash
go run github.com/99designs/gqlgen init
```

To generate graphql files.

```bash
gqlgen generate
```

> On docker

```bash
~/go/bin/gqlgen generate
```
