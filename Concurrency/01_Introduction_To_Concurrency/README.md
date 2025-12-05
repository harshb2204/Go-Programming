## Introduction to Concurrency

Modern computing has run into the physical limits predicted by **Moore’s law**: we no longer get dramatically faster single CPUs every couple of years. Instead, we mostly get *more* CPUs (cores). To keep improving performance, we now need to write programs that can **do many things at once** instead of just doing one thing faster.

This is where **concurrency** and **parallelism** come in:
- **Concurrency**: structuring a program so that multiple tasks can make progress independently (e.g., handling many network requests at once).
- **Parallelism**: actually running tasks at the same time on multiple cores to finish work faster.

However, **Amdahl’s law** tells us that the speedup from using more cores is limited by the parts of our program that *cannot* be parallelized. If a significant portion of the work is inherently sequential (like waiting for user input or doing a step that must happen in strict order), adding more cores has diminishing returns.


Cloud computing made it easy to scale applications horizontally just add more machines, run more instances, and embarrassingly parallel problems get solved faster. With on-demand resources and global data centers, developers gained massive computing power but also new challenges: provisioning machines, coordinating distributed work, and handling concurrency safely.
As systems grew to “web scale,” applications were expected to handle huge workloads through parallelism, elasticity, and distribution—but at the cost of greater complexity and fault-tolerance issues.



## Why is concurrency hard?
Concurrency means multiple pieces of code running at the same time.
The problem is: when things run at the same time, you can't easily predict the order in which they happen.
Because of this, bugs appear randomly and often only under heavy load.

Even small timing changes—more users, slower disk, faster CPU can reveal hidden bugs years later.

### Race Condtions
A race condition happens when:
- Two operations must happen in a specific order
- But your program doesn’t enforce that order
So the result becomes unpredictable.


### Data Race 
A data race happens when:
- One thread is reading a variable
- Another thread is writing to the same variable
And the order is unknown


```go
var data int

go func() {
    data++   // writing to data
}()

if data == 0 {  // reading data
    fmt.Println(data)
}


```
Two things happen at the same time:
- A goroutine increases data
- The main code checks if data is 0
Since we don't know which runs first, three outcomes are possible:
- It prints nothing
- It prints "0"
- It prints "1"
This randomness is exactly why concurrency is hard.

## What is Atomicity?

Atomic = happens all at once, cannot be interrupted.

If an operation is atomic, no other thread or process can "peek in the middle" or change anything while it's happening.

### Why does context matter?

Atomicity depends on scope or where you're looking:

Something might be atomic inside a program, but not atomic inside the operating system.

Something atomic at the OS level might not be atomic at the CPU level, etc.

So: Atomicity is not universal—it depends on the environment.

### Example: i++ is NOT atomic

Although i++ looks like a single step, it's really:

1. Read i
2. Add 1
3. Write back

Each step individually might be atomic, but the whole sequence is not.

Meaning: Another goroutine could read or write i between these steps.

So i++ can cause race conditions in concurrent code.

### When can i++ be atomic?

Only when:
- There is no concurrency (single-threaded program), or
- The variable i is not visible to other goroutines.

In that limited context, nothing interrupts it → so it acts atomic.

### Why do we care about atomicity?

Because atomic operations are safe in concurrent code.

If something is atomic:
- It cannot be half-done.
- No one else can interfere during its execution.

This helps build correct, predictable concurrent programs.

## Memory Access Synchronization

When multiple goroutines access the same variable without atomic operations, you get a **data race** and the result becomes nondeterministic. In our earlier example with `data++` and an `if data == 0` check, different interleavings can produce different outputs on each run.

A section of code that needs **exclusive access** to a shared resource is called a **critical section**. In that simple example, all of these are critical sections:
- The goroutine that increments `data`
- The `if` that reads `data`
- The `fmt.Printf` that reads `data` for output

One way to protect critical sections is to use a **mutex** to synchronize memory access:

```go
var memoryAccess sync.Mutex
var value int

go func() {
    memoryAccess.Lock()
    value++
    memoryAccess.Unlock()
}()

memoryAccess.Lock()
if value == 0 {
    fmt.Printf("the value is %v.\n", value)
} else {
    fmt.Printf("the value is %v.\n", value)
}
memoryAccess.Unlock()
```

Here, the convention is: any code that touches `value` must call `Lock` before and `Unlock` after. Code inside the `Lock`/`Unlock` pair can assume exclusive access, so the **data race is removed**.

However, this does **not** automatically fix the higher-level **race condition**: the order in which the goroutine and the `if`/`else` run is still nondeterministic. Synchronizing memory access can:
- Remove data races
- Improve safety

## Deadlock

A **deadlock** happens when all concurrent processes are waiting on each other. The program gets stuck and will never recover without outside intervention.

### The Problem

When you use mutexes, you must be careful about the **order** in which you lock them. If two goroutines lock the same mutexes in different orders, they can end up waiting for each other forever.

### Example

```go
type value struct {
    mu    sync.Mutex
    value int
}

printSum := func(v1, v2 *value) {
    v1.mu.Lock()
    defer v1.mu.Unlock()
    
    time.Sleep(2 * time.Second)  // simulate work
    
    v2.mu.Lock()
    defer v2.mu.Unlock()
    
    fmt.Printf("sum=%v\n", v1.value + v2.value)
}

var a, b value
go printSum(&a, &b)  // locks a, then tries to lock b
go printSum(&b, &a)  // locks b, then tries to lock a
```

What happens:
- First goroutine locks `a`, then waits to lock `b`
- Second goroutine locks `b`, then waits to lock `a`
- Both are stuck waiting for each other → **deadlock**

![](/diagrams/deadlock.png)

The Go runtime can detect deadlocks (when all goroutines are blocked), but it's better to prevent them by:
- Always locking mutexes in the **same order** across all goroutines
- Using timeouts or other mechanisms to avoid infinite waits

## Livelock

A **livelock** happens when concurrent processes are actively running, but they keep reacting to each other in a way that prevents any real progress.

Think of two people in a hallway: one moves left to let the other pass, but the other also moves left. Then both move right, then left again—forever stuck in a loop.

Unlike a deadlock (where processes are blocked), a livelock means processes are **active** but not making progress. This makes livelocks harder to spot because:
- The program appears to be working (CPU is busy)
- Processes are executing code
- But no actual work gets done

Livelocks often occur when processes try to prevent deadlocks **without coordinating** with each other. The solution is better coordination: processes need to agree on a strategy (e.g., one person stands still while the other moves).

## Starvation

**Starvation** happens when a concurrent process cannot get the resources it needs to perform work.

Livelock is actually a special case of starvation: all processes are starved equally, so no work gets done. More broadly, starvation usually means one or more **greedy processes** are unfairly preventing others from working efficiently—or at all.

### Example: Greedy vs Polite Worker

A greedy worker might hold a shared lock for its entire work loop, while a polite worker only locks when needed. Even if both do the same amount of work, the greedy worker can get almost twice as much done because it starves the polite worker of access to the shared resource.

The greedy worker unnecessarily expands its hold on the lock beyond its critical section, preventing the polite worker from performing work efficiently.

### Finding a Balance

When using mutexes, you face a trade-off:
- **Coarse-grained locking** (holding locks longer): better performance, but can starve other processes
- **Fine-grained locking** (only locking critical sections): more fair, but more overhead from frequent lock/unlock operations

Start with fine-grained synchronization (only lock critical sections). If synchronization becomes a performance problem, you can broaden the scope. It's much harder to go the other way.

### Beyond Locks

Starvation can apply to any shared resource:
- CPU time
- Memory
- File handles
- Database connections
- Network bandwidth

Starvation can cause your program to behave inefficiently or incorrectly. In severe cases, a greedy process can completely prevent another from accomplishing any work.
