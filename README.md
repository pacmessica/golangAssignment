# Go Assignment

## Challenges
1) Setup dev environment
2) Write HelloWorld
```
go run helloworld.go
```
3) Write a simple HTTP server, return JSON object with current time, like
{
	"time": "2006-01-02T15:04:05Z07:00"
}

```
go run server.go
```

4) Create an array of random key-value items (structs), like:
[{ key: "c", value: "123" }, { key: "a", value: "456" }, { key: "b", value: "789" }]
and sort the by key and by value
```
go run sortedStrucs.go
```

5) Write a program that creates N items like, where N is specified by a command line argument, where the key has a random value of 6 characters
[{ value: "abcdef",  hex: "", base32: "" }, ...]
enrich the items with hex and base32 based representation by processing each one in a go-routine

6) Write a program that writes the output of exercise 5 to a json file
```
go run UserInputMap/UserInputMap.go UserInputMap/Main.go
```

7) Write a program with 3 go-routines, which does the following:
- Go routine 1 generates continuousstream of items, which are sent to GR2. The time between sending items varies between 10ms and 100ms
- Go routine 2 makes micro batches, either when it receives 25 items are 2 seconds are passed
- Go routine 3 reports the throughput to screen

```
go run Throughput/Main.go
```

8) Create a simple docker container based on alpine 3.5 running the above program

```
sudo docker build -t throughput .
```
