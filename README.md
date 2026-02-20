A repository dedicated to mastering Go through practical code examples and conceptual notes. This project covers essential topics such as goroutines, channels, synchronization primitives, and idiomatic error handling.

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
