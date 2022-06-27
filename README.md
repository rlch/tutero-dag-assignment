# Tutero Assignment

## Context

In this assignment, you will be responsible for guessing Randy’s favourite Mathematics skill. You are given a knowledge graph (directed acyclic graph), representing all the skills that Randy knows; where the parents of a given skill `s` are pre-requisites to knowing `s`, and the children of `s` require the knowledge of `s` to be learnt.

With each guess `g` that you submit, Randy will only tell you whether `g` is his favourite skill, whether he learnt `g` before his favourite skill, or whether he learnt `g` after his favourite skill. (Randy is annoying) 

---

The code-base you will be working on consists of the following packages:

- `pkg/step`
    - This is where the bulk of your work will be.
    - At a bare minimum, you should implement and test the `Step` function, which is the algorithm for computing your guess for Randy’s favourite skill.
- `pkg/src/graph`
    - Contains a `Graph` struct representing a directed acyclic graph, as well as useful methods for interacting with DAG’s
        - You may choose to iterate on these methods
- `cmd`
    - This command will use your `Step` until it returns the correct guess; printing the number of steps taken.
    - You may run this with `go run ./cmd`
    - You may benchmark your implementation with `go test -bench=. ./cmd` — try to get your `ns/op` and `steps` as low as possible!

Actionable blocks of code are denoted with `//*`

---

## Technical Expectations

- Usage of dependencies other than those included is **not allowed.**
    - This allows you to demonstrate your understanding of core `Go`.
- Clean, idiomatic code.
    - [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)
- Your code should be well-tested.
- Your code will be benchmarked for performance. You should try to minimise the steps taken to achieve the target, as well as the unit-performance of your `Step` function.

---

## Submission

- Version control your assignment on **GitHub**
- Give access to [@rlch](https://github.com/rlch)
- You should send the link to your repository, how much time you took to complete the assignment; as well as a video detailing all the features within your app to the following emails when you are finished: sonny@tutero.com.au, richard@tutero.com.au

---

## Tips + Notes

- Try your best to use your initiative, but if you have any issues please contact  [richard@tutero.com.au](mailto:richard@tutero.com.au)
- We will set up a final interview with you after your assignment is complete to go over your code and ask you various questions about your assignment.
