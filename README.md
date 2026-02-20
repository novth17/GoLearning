A repository of "minimalist" Go implementations designed to isolate and study core language mechanics. 
Each program is stripped down to the smallest possible amount of code to observe Go's runtime behavior, concurrency patterns, and memory safety.

## Repository Structure
The project is organized into several sub-directories, each focusing on a specific Go concept.

## 1. Concurrency Patterns
- channel-basic/: Demonstrates the fundamental behavior of buffered and unbuffered channels, including closing channels and receiving from closed ones.
- channel-philo/: An implementation of the "Dining Philosophers" problem using channels to represent forks and manage resource contention.
- worker-pool/: Simulates a pizza factory to demonstrate the worker pool pattern. It uses a shared queue (channel) where limited workers pull tasks to process them concurrently.
- waitgroup/: Shows how to use sync.WaitGroup to synchronize multiple goroutines and ensure the main program waits for them to finish.
- mutex/: Explores thread safety using sync.Mutex and sync/atomic to safely increment shared variables across multiple goroutines.
- context/: Illustrates how to use context.Context to signal cancellation to goroutines, using an "orchestra and singers" analogy where singers stop when the orchestra shuts down.

## 2. Core Concepts & Safety
- error-handling/: A practical ATM machine simulation that demonstrates branching logic based on custom error types and the use of pointer receivers to modify struct state.
- coopVsPreemptive/: Explores the nature of the Go scheduler and how it handles long-running loops versus cooperative yielding.
- main-exit-early/: A demonstration of what happens when the main goroutine finishes before background goroutines have completed their tasks.

## Key Learnings Included
- Unbuffered Channels: Send and receive operations must happen simultaneously (handshake).
- Buffered Channels: Sends only block when the buffer is full; receives only block when the buffer is empty.
- Channel Closure: Closing a channel signals that no more values will be sent. Receivers can still drain remaining values from a closed channel.
- Atomic Operations: Using sync/atomic for low-level memory safety when performing simple arithmetic on shared integers.

## Getting Started
To run any of the examples, navigate to the specific directory and use the Go run command:

``` go run code-practice/worker-pool/pizza.go ```
