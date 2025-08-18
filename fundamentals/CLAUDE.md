# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

As the repository is a collection of individual Go programs, there are no top-level build or test scripts. To run a specific example, navigate to its directory and use `go run`:

```bash
cd <directory_name>
go run .
```

To run tests for a specific package, navigate to its directory and use `go test`:

```bash
cd <directory_name>
go test .
```

## Architecture

The repository is organized as a series of standalone Go programs, each in its own directory. These programs demonstrate various fundamental concepts of the Go language, from basic "Hello, World!" to more advanced topics like goroutines, channels, and reflection. Each directory is self-contained and can be studied and run independently. Some directories also contain a `note.md` file with explanations of the concepts demonstrated in the code.
