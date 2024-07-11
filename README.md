# Go Playground

## Project Description

Welcome to the Go Playground! This repository is a collection of Go language examples and experiments, showcasing a
range of features and techniques in Go programming. It aims to help developers better understand and utilize the Go
language effectively.

## Prerequisites

- **Go Version**: This repository requires [Go 1.22.0](https://golang.org/dl/) or higher. Make sure you have the correct
  version installed to run the examples provided.

## Directory Structure

Here's a breakdown of what you'll find in this repository:

- **Hello World**
    - [Hello World Example](/helloworld/main.go)
- **Go Essentials**

    - **Functions in Depth**
        - [Detailed exploration of functions in Go (functions as value, anonymous functions,
          closures, recursion, variadic functions)](/functions-in-depth/functionsInDepth.go)
        - Projects:
            - [Card Handling Example](/cards)
            - [Basic Banking Application](/bank/bank.go)
            - [Investment Calculator](/investment-calculator/investment_calculator.go)
            - [Price Calculator (includes interface, working with packages, function, goroutines and channel examples)](/price-calculator)
- **Working with Packages**
    - [Examples](/workingWithPackages)
- **Advanced Data Handling**
    - **Pointers**
        - [Pointer Examples](/pointers/pointers.go)
    - **Structs**
        - [Basic Structs Usage](/structs)
        - [More Structs Examples](/struct/main.go)
    - **Interfaces**
        - [Deep Dive into Interfaces](/interfaces-in-depth)
        - [Additional Interface Examples](/interfaces)
- **Collections**
    - **Lists and Maps**
        - [List Operations](/listsOfData/lists.go)
        - [Map Operations](/listsOfData/maps.go)
        - [Maps Examples](/map/main.go)
- **Concurrency**
    - [Goroutines Usage Example](/goroutines/main.go)
    - [Channel Usage Example](/channels/main.go)

## Installation Instructions

To set up this project locally, follow these steps:

```bash
git clone https://github.com/ozgen/go-playground.git
cd go-playground
```

## Usage

To run any of the included Go programs, navigate into the corresponding directory and execute the Go file. For example:

```bash
cd helloworld
go run main.go
```
