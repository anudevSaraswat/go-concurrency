This program is a demonstration of Sequential, Concurrent and Parallel execution of a Go program. Before diving deep let's understand few terminologies.

## Concurrency vs Parallelism

Concurrency is handling different tasks at once and Parallelism is doing different tasks at once.

## runtime.GOMAXPROCS

`GOMAXPROCS` is a function of `runtime` package which lets you configure the number of CPUs that your program will be utilizing. Default value is set to `runtime.NumCPU()`.

In this program I have mimicked 3 scenarios and you can see the output below :-

<img width="526" alt="Screenshot 2024-09-17 at 00 04 29" src="https://github.com/user-attachments/assets/74fb6360-a54e-438a-9294-7e1fad3b7e63">

As you can see there's a huge difference in execution times if we compare 1st to 2nd and 3rd scenarios.
Reason is that, in 2nd and 3rd case we are leveraging the power of Go routines which is helping us to
concurrently (2nd) or parallely (3rd) do things.
