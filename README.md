# gomips

MIPS interpreter in Go

### Description

When learning some assembly at my University, we used [MARS](https://courses.missouristate.edu/kenvollmar/mars/), a MIPS Assembler and Runtime Simulator.
I've wanted to build some type of interpreter for awhile, so I originally started working on this as a TypeScript package so I could compile it and run it
in the browser.

I've been using Go a little bit for some time now and I wanted to rebuild this package as a Go app that will compile to WASM and still run in the browser.

### Features to be implemented

- Tokenize MIPS
- Parse MIPS
- Interpret MIPS
- Add syscall functionality
- Other things (bug fixes)

### Installation

At some point, this will be able to be installed as a Go package from something like `nstruck.dev/gomips` or something.

### License

I love open source, I always use GPL 3.0. Have at it.
