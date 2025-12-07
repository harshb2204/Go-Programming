## Goroutines

Goroutines are lightweight functions that run **concurrently** (not necessarily in parallel) with other code in the same Go program.  
Every Go program starts with at least one goroutine: the **main goroutine**, which runs the `main` function.

### Starting a goroutine

You start a new goroutine by putting the `go` keyword in front of a function call:

```go
package main

import "fmt"

func main() {
	go sayHello()          // run concurrently with main

	// continue doing other work in main...
	fmt.Println("from main")
}

func sayHello() {
	fmt.Println("hello")
}
```

You can also start goroutines from anonymous functions:

```go
go func() {
	fmt.Println("hello")
}() // note: the anonymous function is invoked immediately
```

Or assign the anonymous function to a variable and start it:

```go
sayHello := func() {
	fmt.Println("hello")
}

go sayHello()
```

### How goroutines work (high level)

- **Not OS threads**: Goroutines are not operating-system threads; they are a higher-level abstraction managed by the Go runtime.
- **Coroutines managed by the runtime**: Conceptually, goroutines are a special kind of coroutine. The Go runtime automatically suspends a goroutine when it blocks (e.g., on I/O or channel operations) and resumes it when it can continue.
- **M:N scheduler**: Go uses an M:N scheduler that maps many goroutines onto a smaller number of OS threads and schedules goroutines across them.
- **Fork-join model**: When you start a goroutine with `go`, you “fork” a new concurrent path of execution. Later, your program logic typically “joins” these paths again (often via synchronization primitives like channels or `sync.WaitGroup`).
![](/diagrams/forkandjoin.png)

### Fork and join with goroutines

- The `go` statement is how Go performs a **fork**: it starts a new goroutine (a new concurrent path of execution).
- A **join point** is where we wait for that goroutine to finish so the program behaves correctly.

If we just start a goroutine and let `main` exit, the goroutine may never run:

```go
sayHello := func() {
	fmt.Println("hello")
}

go sayHello()         // may or may not run before main exits
// main continues and can exit immediately
```

To create a join point, we synchronize `main` with the goroutine. A common tool is `sync.WaitGroup`:

```go
var wg sync.WaitGroup

sayHello := func() {
	defer wg.Done()    // signal that this goroutine is finished
	fmt.Println("hello")
}

wg.Add(1)             // we are waiting for 1 goroutine
go sayHello()         // fork
wg.Wait()             // join: block until sayHello finishes
```

This guarantees that `"hello"` is printed before the program exits.

> **Important thing to remember (samjhun ghe tu adhi): a goroutine is just a function running concurrently with other functions, started with the `go` keyword.**

### Understanding closures and loop variables in goroutines

This section explains how Go handles closures inside goroutines, why a common loop pattern leads to unexpected output, and how to write goroutine loops correctly.

#### Goroutines share the same address space

When you start a goroutine, it runs concurrently but in the same memory space as the main program. This means goroutines can access variables defined outside their function, including loop variables.

A goroutine is essentially a function that is scheduled to run independently at some later time.

#### The problem: closures over loop variables

Consider this loop:

```go
var wg sync.WaitGroup

for _, salutation := range []string{"hello", "greetings", "good day"} {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(salutation)
	}()
}

wg.Wait()
```

Many expect it to print (in any order):

```text
hello
greetings
good day
```

But it often prints:

```text
good day
good day
good day
```

#### Why this happens

- The anonymous function forms a **closure** over the variable `salutation`.
- `salutation` is a **single variable** reused in every loop iteration.
- The goroutines often start after the loop finishes, so all of them reference the same final value of `salutation`, which is `"good day"`.

#### How Go handles memory here

After the loop ends, the variable `salutation` would normally go out of scope. However:

- Go detects that goroutines still reference it.
- The runtime moves it to the heap so it remains accessible.

The memory is safe, but the value is not what you expect because it was updated by later iterations before the goroutines ran.

#### The correct pattern: pass the variable as a parameter

To ensure each goroutine gets the right value for its iteration, pass the loop variable into the function:

```go
var wg sync.WaitGroup

for _, salutation := range []string{"hello", "greetings", "good day"} {
	wg.Add(1)

	go func(salutation string) {
		defer wg.Done()
		fmt.Println(salutation)
	}(salutation)
}

wg.Wait()
```

#### Why this works

- The goroutine function receives its **own copy** of the string.
- The closure no longer depends on the outer loop variable.
- Each goroutine prints the value corresponding to its iteration.

Example output (order may vary, but values are correct):

```text
good day
hello
greetings
```

#### What the syntax means

This pattern:

```go
go func(salutation string) {
	// ...
}(salutation)
```

works like this:

- `func(salutation string) { ... }` defines an anonymous function that expects a parameter.
- `(salutation)` immediately invokes the anonymous function, passing the current loop value as the argument.

This is known as an **immediately invoked function expression (IIFE)**.
