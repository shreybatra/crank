# Crank CLI

This is the cli tool to connect and use [CrankDB](https://github.com/shreybatra/crankdb)

## Requirements
- Golang 1.16

## Runnning
- Download and start [CrankDB](https://github.com/shreybatra/crankdb).
- Download crank-cli -
    - `go get github.com/shreybatra/crank` or
    - `go install github.com/shreybatra/crank@latest`
- Run cli and connect to CrankDB using - `crank`
- Optionally, you can write your commands in a `.gsb` file and pass it as an argument to crank cli - `crank commands.gsb`. The gsb file will follow the same commands as interactive shell.
- CrankDB server should be running before you can connect.

## Query Language

### SET

```js
// Set a string (quotes are mandatory to declare a string)
> set hello "world" 
// success:true

// Set an integer
> set age 30 
// success:true

// Set a float value
> set marks 98.5 
// success:true

// Set a JSON object (Nested objects are allowed)
> set user_1 {"name" : "shrey", "age": "23" }
// success:true

// Set a JSON array object (Any type of JSON serializable object can be stored)
> set arr [ 1, "shrey", {"marks": 90}]
// success:true
```

### GET

```js
> get hello
// dataType:STRING stringVal:"world"

> get marks
// dataType:DOUBLE doubleVal:98.5

> get user_1
// dataType:JSON jsonVal:"{\"age\":\"23\",\"name\":\"shrey\"}"
```

### FIND (To run queries on keys storing JSON objects)

```js
> find {}
// returns every JSON (map) storing key
// [ { user_1 map[age:23 name:shrey ] } ]

> find {"name": "shrey"}
// finds all keys having a JSON object with name as "shrey"
```

## Steps to build
- Git clone the repository.
- Tidy dependencies using - `go mod tidy`
- Build the application - `go build .`
- Run the cli application - `./crank`