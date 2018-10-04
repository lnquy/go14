# Go Learning Guide

```
@contact: lnquy.it@gmail.com / lnquy@tma.com.vn
@lastEditted: Oct 4th 2018
@version: 0.0.1
```

Some thoughts:

- Do not try to force yourself like "I have to learn it in X days". The Beginner and Medium level are quite easy to reach but the Advanced level need a lot more time, so keep learning everyday.
  In my experience, it takes around 1-2 days to finish Beginner level, 3-7 days to pass Medium level and >= 2 weeks to comfortably working on a real product.
- Go is pretty simple so just start learning from web is enough. After that, feel free to read books to understand deeply in advanced things.
- Go is not C, Java... so please do not map the knowledge you already learned in other programming languages to Go. Just learn and do it the right (Go) way.
- The fastest way to learn programming is working on actual project(s), so grab an idea and code it in Go.
- Please be active and join other Go groups/events (in TMA and/or Vietnam, world) to get help and familiar with Go community.

### I.  Beginner

At this level, you should learn about Go basic syntaxes and its components. 
The key point of this level is you should learn slowly until becomes familiar with Go syntax and can freely write simple program (as noted at exercises).

#### What you should learn

- Primitive types: byte, int/uint (16, 32 and 64 bits), float (32/64 bits), bool, string, array and map.
  - Know how to cast from a type to another type.
- Variables and constants.
- If-else and switch-case.
- For loop.
- How to write a function.
- Handle errors.
- How to run/build Go program.

#### Documents

*Note: You don't have to read the whole document at once, just read/practice the basics first.*

- Tour of Go (official document): https://tour.golang.org/welcome/1
- Go by example: https://gobyexample.com/

#### Exercises

1. Write a program to calculate the perimeter and square of a rectangle, triangle, round,...

2. Read the marks of Math, Physics and Chemistry from stdin, return the ranking based on average mark.
   - avgMark < 5: Bad
   - avgMark <8: Ok
   - avgMark < 9: Good
   - Otherwise Excellent

### II. Medium

You should learn more about things listed in Beginner level and also learn new things.

#### What you should learn

- Package
- Init()
- Defer, panic, recover
- Slice
- Pointer
- Struct
- Interface
- Methods
- Concurrency primitives: goroutine, channel, select
- Go tools: gofmt, godoc, glide,...

#### Documents

- Continue reading the documents from Beginner level.
- Go tutorial series: https://golangbot.com/learn-golang-series/
- 50 shades of Go: http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
- Go Effective: https://golang.org/doc/effective_go.html (must read, but you can just read the parts you want to know first and come back later to read the whole document).

#### Exercises

1. Write TODO app.
2. 

### III. Advanced

A lot of things you must learn in this level and actually we don't have the time range since you must keep learning everyday.
Basically, at this level, you must understand deeply on Go internals, concurrency programming and must work on Go projects to reinforce your knowledge.

#### What you should learn (deeply understand)

- Array and slice.
- Struct and interface.
- goroutine, channel and select.
- sync package (Mutex, RWMutex, WaitGroup, Pool...).
- Concurrency model:
  - Fan-in
  - Fan-out
  - Worker pool
  - Semaphore
  - ...
- Databases
- net/http package: Write web server/services.
- Take a look on godoc of important Go's packages.

#### Documents

- Go Wiki: https://github.com/golang/go/wiki
- Go Blog: https://blog.golang.org/
- Go source code: https://github.com/golang/go
- Go Academy: https://www.youtube.com/channel/UCx9QVEApa5BKLw9r8cnOFEA

#### Exercises

1. Write a crawler to crawl data from any website you like.
2. Write a web services that serve REST APIs.
3. Write your own blog.
4. Build a microservices system.

### IV. For KSHOP Interview

You must at least comfortably write simple Go application using everything listed on Beginner and Medium level.
Should have better understanding on things listed in Advanced level.
Focus on net/http package and able to write a simple web server/services to serve web contents / REST APIs.
Take a look on these questions: https://goo.gl/azHjqp

Good luck!
