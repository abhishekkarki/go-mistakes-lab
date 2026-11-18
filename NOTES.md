# NOTES

One line per mistake: the rule, then why. Fill each line in on its day.
Weekly review: reread this week and last week, then explain 5 random ones from memory.

## Week 1: Data types (Days 2–6)

- [ ] #17 Creating confusion with octal literals (Day 2): 
- [ ] #18 Neglecting integer overflows (Day 2): 
- [ ] #19 Not understanding floating-points (Day 2): 
- [ ] #20 Not understanding slice length and capacity (Day 3): 
- [ ] #21 Inefficient slice initialization (Day 3): 
- [ ] #22 Being confused about nil vs. empty slice (Day 3): 
- [ ] #23 Not properly checking if a slice is empty (Day 3): 
- [ ] #24 Not making slice copies correctly (Day 4): 
- [ ] #25 Unexpected side effects using slice append (Day 4): 
- [ ] #26 Slices and memory leaks (Day 4): 
- [ ] #27 Inefficient map initialization (Day 5): 
- [ ] #28 Maps and memory leaks (Day 5): 
- [ ] #29 Comparing values incorrectly (Day 5): 

## Week 2: Concurrency foundations (Days 7–12)

- [ ] #55 Mixing up concurrency and parallelism (Day 7): 
- [ ] #56 Thinking concurrency is always faster (Day 7): 
- [ ] #57 Being puzzled about when to use channels or mutexes (Day 8): 
- [ ] #58 Not understanding race problems (Day 8): 
- [ ] #59 Not understanding the concurrency impacts of a workload type (Day 9): 
- [ ] #60 Misunderstanding Go contexts (Day 10): 
- [ ] #61 Propagating an inappropriate context (Day 10): 
- [ ] #62 Starting a goroutine without knowing when to stop it (Day 11): 
- [ ] #63 Not being careful with goroutines and loop variables (Day 11): 
- [ ] #64 Expecting deterministic behavior using select and channels (Day 11): 

## Week 3: Concurrency in practice (Days 13–17)

- [ ] #65 Not using notification channels (Day 13): 
- [ ] #66 Not using nil channels (Day 13): 
- [ ] #67 Being puzzled about channel size (Day 13): 
- [ ] #68 Forgetting about possible side effects with string formatting (Day 14): 
- [ ] #69 Creating data races with append (Day 14): 
- [ ] #70 Using mutexes inaccurately with slices and maps (Day 14): 
- [ ] #71 Misusing sync.WaitGroup (Day 15): 
- [ ] #72 Forgetting about sync.Cond (Day 15): 
- [ ] #73 Not using errgroup (Day 16): 
- [ ] #74 Copying a sync type (Day 16): 

## Week 4: Errors and standard library (Days 18–22)

- [ ] #48 Panicking (Day 18): 
- [ ] #49 Ignoring when to wrap an error (Day 18): 
- [ ] #50 Checking an error type inaccurately (Day 18): 
- [ ] #51 Checking an error value inaccurately (Day 18): 
- [ ] #52 Handling an error twice (Day 19): 
- [ ] #53 Not handling an error (Day 19): 
- [ ] #54 Not handling defer errors (Day 19): 
- [ ] #75 Providing a wrong time duration (Day 20): 
- [ ] #76 time.After and memory leaks (Day 20): 
- [ ] #77 Common JSON-handling mistakes (Day 20): 
- [ ] #78 Common SQL mistakes (Day 21): 
- [ ] #79 Not closing transient resources (Day 21): 
- [ ] #80 Forgetting the return statement after replying to an HTTP request (Day 21): 
- [ ] #81 Using the default HTTP client and server (Day 21): 

## Week 5: Testing (Days 23–26)

- [ ] #82 Not categorizing tests (Day 23): 
- [ ] #83 Not enabling the -race flag (Day 23): 
- [ ] #84 Not using test execution modes (Day 23): 
- [ ] #85 Not using table-driven tests (Day 23): 
- [ ] #86 Sleeping in unit tests (Day 24): 
- [ ] #87 Not dealing with the time API efficiently (Day 24): 
- [ ] #88 Not using testing utility packages (Day 24): 
- [ ] #89 Writing inaccurate benchmarks (Day 25): 
- [ ] #90 Not exploring all the Go testing features (Day 25): 

## Week 6: Optimizations (Days 27–31)

- [ ] #91 Not understanding CPU caches (Day 27): 
- [ ] #92 Writing concurrent code that leads to false sharing (Day 27): 
- [ ] #93 Not taking into account instruction-level parallelism (Day 28): 
- [ ] #94 Not being aware of data alignment (Day 28): 
- [ ] #95 Not understanding stack vs. heap (Day 29): 
- [ ] #96 Not knowing how to reduce allocations (Day 29): 
- [ ] #97 Not relying on inlining (Day 29): 
- [ ] #98 Not using Go diagnostics tooling (Day 30): 
- [ ] #99 Not understanding how the GC works (Day 30): 
- [ ] #100 Not understanding the impacts of running Go in Docker and Kubernetes (Day 30): 

## Week 7: Code and project organization (Days 32–35)

- [ ] #1 Unintended variable shadowing (Day 32): 
- [ ] #2 Unnecessary nested code (Day 32): 
- [ ] #3 Misusing init functions (Day 32): 
- [ ] #4 Overusing getters and setters (Day 32): 
- [ ] #5 Interface pollution (Day 33): 
- [ ] #6 Interface on the producer side (Day 33): 
- [ ] #7 Returning interfaces (Day 33): 
- [ ] #8 any says nothing (Day 33): 
- [ ] #9 Being confused about when to use generics (Day 34): 
- [ ] #10 Not being aware of the possible problems with type embedding (Day 34): 
- [ ] #11 Not using the functional options pattern (Day 34): 
- [ ] #12 Project misorganization (Day 35): 
- [ ] #13 Creating utility packages (Day 35): 
- [ ] #14 Ignoring package name collisions (Day 35): 
- [ ] #15 Missing code documentation (Day 35): 
- [ ] #16 Not using linters (Day 35): 

## Week 8: Control flow, strings, functions (Days 36–40)

- [ ] #30 Ignoring the fact that elements are copied in range loops (Day 36): 
- [ ] #31 Ignoring how arguments are evaluated in range loops (Day 36): 
- [ ] #32 Ignoring the impact of using pointer elements in range loops (Day 36): 
- [ ] #33 Making wrong assumptions during map iterations (Day 36): 
- [ ] #34 Ignoring how the break statement works (Day 36): 
- [ ] #35 Using defer inside a loop (Day 36): 
- [ ] #36 Not understanding the concept of a rune (Day 37): 
- [ ] #37 Inaccurate string iteration (Day 37): 
- [ ] #38 Misusing trim functions (Day 37): 
- [ ] #39 Under-optimized string concatenation (Day 37): 
- [ ] #40 Useless string conversions (Day 37): 
- [ ] #41 Substring and memory leaks (Day 37): 
- [ ] #42 Not knowing which type of receiver to use (Day 38): 
- [ ] #43 Never using named result parameters (Day 38): 
- [ ] #44 Unintended side effects with named result parameters (Day 38): 
- [ ] #45 Returning a nil receiver (Day 38): 
- [ ] #46 Using a filename as a function input (Day 38): 
- [ ] #47 Ignoring how defer arguments and receivers are evaluated (Day 38): 

