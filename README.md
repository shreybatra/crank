# Crank CLI

This is the cli tool to connect and use [CrankDB](https://github.com/shreybatra/crankdb)

## Requirements
- Golang 1.16

## Runnning
- Download application - `go get github.com/shreybatra/crank`
- Run server with command - `crank`
- CrankDB server should be running before you can connect.

## Query Language

### SET

```js
// Set a string (quotes are mandatory to declare a string)
> set hello "world" 
// "hello set"

// Set an integer
> set age 30 
// "age set"

// Set a float value
> set marks 98.5 
// marks set

// Set a JSON object (Nested objects are allowed)
> set user_1 {"name" : "shrey", "age": "23" }
// user_1 set

// Set a JSON array object (Any type of JSON serializable object can be stored)
> set arr [ 1, "shrey", {"marks": 90}]
// arr set
```

### GET

```js
> get hello
// "world"

> get marks
// 98.5

> get user_1
// {"name" : "shrey", "age": "23" }
```

### FIND (To run queries on keys storing JSON objects)

```js
> find {}
// returns every JSON storing key

> find {"name": "shrey"}
// finds all keys having a JSON object with name as "shrey"
```

## Steps to build
- Run command - `go get github.com/shreybatra/crank`
- Tidy dependencies using - `go mod tidy`
- Build the application - `go build .`
- Run the server - `./crank`