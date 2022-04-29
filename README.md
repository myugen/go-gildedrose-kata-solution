# Gilded Rose Refactoring Kata

This kata in Go language is an updated revision of the existing one in
the [Emily Bache's repository](https://github.com/emilybache/GildedRose-Refactoring-Kata).

## How to start this Kata

Just fork the repository and read the file with the [specifications](./SPECIFICATIONS.md) of the
code.

## 🧾 Features

- Go 1.18
- Go Modules

## ⚙️ Usage

Running the application:

```shell
$ go run texttest_fixtures.go [<number-of-days>; default: 2]
```

## 🧪 Test

Running all tests:

```shell
$ go test ./...
```

Running specific tests:

```shell
$ go test <path to one file or directory> <path to another file or directory>
```

Running test coverage:

```shell
$ go test ./... -coverprofile=cover.out

$ go tool cover -html=cover.out
```

## 📦 Build

Building the application:

```shell
$ go build .
```

After that you can run the binary:

```shell
$ ./go-gildedrose [<number-of-days>; default: 2]
```
