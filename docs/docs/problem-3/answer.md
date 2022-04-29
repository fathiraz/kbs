---
sidebar_label: Answer
sidebar_position: 1
---

# Answer

The solution of this problem is get greatest common divisor **using Euclidean Algorithm**.

## What you'll need
- [Golang](https://go.dev/dl//) version 1.17 or above
- Basic knowledge about go mod https://github.com/golang/go/wiki/Modules

## Prerequisites
```bash
go mod download
```

## Default configuration
You can change default configuration on file .env
```dotenv
DEFAULT_CAKES=20
DEFAULT_APPLES=25
DEFAULT_OWNER=Ainun
```

## Running the app
- **command:**
    ```bash
    go run main.go -cakes=20 -apples=25 -owner=Ainun
    ```
- **result:**
    ```bash
    Ainun have 20 cakes and 25 apples.
	Ainun can make 5 box.
	There are 4 cakes and 5 apples in each box.
    ```

## Test
```bash
go test ./... -v
```

## Source code
You can check the source code right here **https://github.com/fathiraz/kbs/tree/main/fpb**
