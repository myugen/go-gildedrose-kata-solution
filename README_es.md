# Kata de refactorización de La Rosa Dorada (Gilded Rose)

Esta Kata hecha en el lenguaje Go es una version actualizada de la ya existente en
el [repositorio de Emily Bache](https://github.com/emilybache/GildedRose-Refactoring-Kata).

## Como empezar la kata

Realiza un _fork_ de este repositorio y lee el fichero con las
[especificaciones](./SPECIFICATIONS_es.md) del código.

## 🧾 Features

- Go 1.18
- Go Modules

## ⚙️ Ejecutar la aplicación

Para ejecutar la aplicación:

```shell
$ go run texttest_fixtures.go [<number-of-days>; default: 2]
```

## 🧪 Test

Para ejecutar todos los test:

```shell
$ go test ./...
```

Para ejecutar uno o varios tests en particular:

```shell
$ go test <ruta del archivo o carpeta> <ruta de otro archivo o carpeta> ...
```

Para visualizar la cobertura de tests:

```shell
$ go test ./... -coverprofile=cover.out

$ go tool cover -html=cover.out
```

## 📦 Crear binario

Para crear el binario:

```shell
$ go build .
```

Una vez creado, para ejecutar la aplicación:

```shell
$ ./go-gildedrose [<number-of-days>; default: 2]
```
