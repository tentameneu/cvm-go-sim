# CVM Go Simulator

Simulation of CVM algorithm implemented in Go based on paper [The CVM Algorithm for Estimating Distinct Elements in Streams](https://cs.stanford.edu/~knuth/papers/cvm-note.pdf).

## Build

Build application using bash script in ```scripts/``` directory with:

```bash
scripts/build.sh
```

This will create ```cmd/cvm-go-sim/cvm-sim``` executable binary file.

## Run

Run application by running compiled binary with:

```bash
cmd/cvm-go-sim/cvm-sim
```

## Testing

Running specific or all tests is easily available through VS Code editor in `Testing` tab.

To run all unit tests with coverage profiling, use script:

```bash
scripts/run_coverage.sh
```

Script above will try to automatically open coverage HTML report in your default browser.

## Parameters

All available parameters that can be passed to binary are visible through `cvm-sim --help` command:

```console
cvm-sim --help

This program runs CVM algorithm simulator.

Algorithm estimates number of distinct elements in stream of elements much bigger than available buffer size.

Usage:

cvm-sim [arguments]

Supported arguments:

  -buffer-size int
        number of elements that can be stored in buffer while processing stream (default 10000)
  -distinct int
        number of distinct elements in generated test stream (default 5000000)
  -file-path string
        path to file containing numbers separated by whitespace. used in stream generated from file
  -log-level string
        logging level. valid values are: [info, debug, deep] (default "info")
  -random-max int
        used in random stream generator - generates values in range [random-min, random-max] (default 10000000)
  -random-min int
        used in random stream generator - generates values in range [random-min, random-max]
  -stream-type string
        how to generate test stream of elements. valid values are: [incremental, random, file] (default "incremental")
  -total int
        total number of elements in generated test stream (default 100000000)
```
