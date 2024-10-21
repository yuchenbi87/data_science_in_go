# Recommendation to Management

## Overview

The Go program uses the `github.com/montanaflynn/stats` package to perform statistical analysis on the Anscombe dataset. The primary objective is to ensure that Go's statistical capabilities can provide results comparable to Python and R for linear regression analysis.

## Recommendation

The `github.com/montanaflynn/stats` package in Go provides core statistical functions such as mean, correlation, and linear regression. The Go package is suitable for basic statistical calculations and provides the added benefits of performance, concurrency, and easier integration into production-level systems.

- **Performance and Concurrency**: Go is known for its speed and performance, particularly for concurrent applications. This makes it suitable for handling statistical tasks efficiently in production environments.
- **Simplicity and Efficiency**: Go's simplicity in syntax and efficient execution makes it a compelling choice for teams looking to integrate statistical capabilities without adding complexity to the development environment.

## Instruction

### Step 1:

Download the source code:

`git clone https://github.com/yuchenbi87/data_science_in_go.git -b master`

### Step 2:

Build the executable from source code:

`cd data_science_in_go`

`go build -o data_science_in_go.exe`

### Step 3:

Run the program:

`.\data_science_in_go`

## UT

Run the command to test the code:

`go test -v`