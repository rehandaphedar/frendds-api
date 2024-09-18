# Introduction

frendds-api is a simple wrapper API to the CLI application [frendds](https://sr.ht/~rehandaphedar/frendds).

- [Demo](https://frendds-api.demos.rehandaphedar.com)

# Installation

```shell
go install git.sr.ht/~rehandaphedar/frendds-api@latest
```

# Usage

```shell
frendds-api [port]
```

Example:

```shell
frendds-api 3005
```

# Routes

- `/[domain]`: Get the D2 file of the domain.
