## Concurrency vs Parallelism

**Concurrency is about structure.** It means your code is written in a way where multiple tasks can start, run, or pause without waiting for each other.

**Parallelism is about execution.** It means multiple tasks are actually running at the same time on different CPU cores.

## Why they are not the same

You can write code that is concurrent, but it may or may not run in parallel depending on the machine.

**Example:** You write two goroutines.

On a 1-core CPU, they cannot run at the same time. The CPU switches between them very fast, giving the illusion of running together. The code is concurrent, but the execution is not parallel.

On a 2-core CPU, they can run truly at the same time. That is parallelism.

**So:** you write concurrent code. Whether it becomes parallel depends on the runtime and hardware.

## Why this distinction matters

You cannot guarantee parallel execution from your code alone.

Parallelism depends on:
- number of CPU cores
- OS scheduling
- Go runtime
- virtual machines or containers running underneath

Your code only describes how tasks could run together (concurrency). It cannot force them to be parallel.

## Why this abstraction is useful

Because Go lets you write concurrent code without worrying about whether it will actually run in parallel.

This makes your code:
- more portable
- more flexible
- easier to reason about

The runtime decides the best way to run it.

## What "context" means here

Parallelism is relative to context.

Consider two tasks that each take one second:
- If you look at them in a five-second window, you might say they ran in parallel.
- If you only look at one second, you might say they ran sequentially.

Context defines what you consider "parallel."

Context can also be:
- a process
- a thread
- a machine

It is simply the boundary inside which you judge whether something is concurrent, parallel, atomic, or correct.

## The essence of the text

You never write parallel code. You write concurrent code. The system decides whether it runs in parallel.

Concurrency is about how you design and structure code. Parallelism is about how the system executes it in real time.

Whether operations can be considered parallel or atomic depends on the context you choose.

## Why CSP exists

Back in the 1970s, people mostly knew how to write sequential programs:
- Do step 1
- Then step 2
- Then step 3

But nobody had a clean, reliable way to structure concurrent programs—programs where multiple tasks run independently. CSP was created to solve that.

## What CSP actually says 

CSP proposes a model where:
- You write independent units of work called processes.
- Each process runs sequentially (simple, one step at a time).
- Processes do not share memory.
- Processes communicate only by sending and receiving messages.

That’s it.

It is like each process is a small machine, and they talk to each other using messages—not by reading each other's memory. This “communication” is the heart of CSP.


The meaning of input/output in CSP  
CSP introduced:
- `!` to send data
- `?` to receive data  

These are communication operations between processes.

Examples:
- `cardreader?cardimage`  
  Read a card from the `cardreader` process and store it in the variable `cardimage`.
- `lineprinter!lineimage`  
  Send the contents of `lineimage` to the `lineprinter` process.

## Why Go’s concurrency model feels different (with Java context)

In many traditional languages like Java, concurrency is built around **threads, executors, and locks**:
- You decide whether to use raw threads or an `ExecutorService`.
- You choose between cached or fixed thread pools.
- You worry about OS thread limits and per-thread memory cost.
- You coordinate access to shared memory with `synchronized`, `ReentrantLock`, and other locking primitives.

Most of this thinking is about **parallelism and system details** (how many threads run, how they are scheduled, how to avoid contention), not about the actual business logic you are trying to express.

### Java way: system-level thinking

When building something like a web server in Java, your mind often goes first to questions such as:
- How many threads should the pool have?
- What is the optimal pool size on this machine?
- Are threads too heavy on this OS?
- How do I avoid contention and deadlocks around shared state?

All of these are questions about **how the work will be run**, not about **what the work actually is**. The abstraction effectively stops at the thread: you must think in terms of threads, pools, and locks.

### Go way: problem-level thinking

In Go, concurrency is built around **goroutines and channels**:
- Each incoming request or task naturally becomes `go handleRequest(conn)`.
- You think: “A user connects → start a goroutine → handle the request → return a response → the goroutine ends.”
- You do not manage thread pools, tune pool sizes, or worry about per-thread cost.

**Goroutines** are very lightweight, and the Go runtime **multiplexes many goroutines onto a small number of OS threads**. You express your solution in terms of logical tasks (goroutines), and let the runtime decide how to map them onto OS threads and CPU cores.

### Channels vs locks

In Java-style designs:
- You usually share memory between threads.
- You protect that shared memory with locks.
- You manually avoid race conditions and deadlocks.

In Go:
- You are encouraged to **share memory by communicating**, not communicate by sharing memory.
- **Channels** are the main tool for communication between goroutines.
- Instead of many threads competing over shared state, you structure your program so goroutines send values through channels and avoid shared mutable memory whenever possible.

This shift changes how you **model solutions**:
- In Java, you first think about threads, thread pools, and lock strategies.
- In Go, you first think about the problem itself (each user, job, worker, or pipeline stage) and express that directly as goroutines and channels, while the runtime takes care of scheduling and parallelism.

### Fewer concurrency bugs, and runtime-driven improvements

Traditional Java concurrency is powerful but easy to get wrong:
- Race conditions, deadlocks, missed notifications, and thread starvation.
- Incorrect or suboptimal pool sizing that only shows up under load.

By contrast, Go’s goroutines and channels:
- Remove most of the need to hand-manage threads and locks.
- Make it more natural to avoid shared mutable state.
- Let the Go runtime’s scheduler handle mapping many goroutines to a few OS threads.

Because goroutines are multiplexed over OS threads, improvements in the Go runtime scheduler can make your program faster **without changing your code**, whereas Java-style designs often require manual thread-pool tuning and synchronization changes to benefit from similar improvements.