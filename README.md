# GO-NEW Tool

This is a small Golang tool, to generate a Golang Starter Project from commandline.
(Only tested on macOS and Ubuntu)

## Install

To install the tool simply clone the github project and install with go

```bash
git clone https://github.com/SeSc838/go-new.git
cd go-new
go install
```

## Usage

At the moment there are three different options for calling the tool.

### 1. Call with no Arguments

```bash
go-new
```

This generates a new Golang Starter Project in current directory with default name `NewGoProject`

### 2. Call with Argument

```bash
go-new projectName
```

This generates a new Golang Starter Project in current directory with name given by the user as an argument (here `projectName`).

### 2. Call with Argument and flag(s)

```bash
go-new projectName -m
```

This generates a new Golang Starter Project in current directory with name given by the user as an argument (here `projectName`), and initialized as a go module. The project name always has to be the first argument after `go-new`. Right now there is only one flag (`-m`) others may follow.
