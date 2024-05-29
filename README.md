# Jasmine
Jasmine - is a personal income tracker. 

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ex1d3/jasmine-cli/go.yml)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ex1d3/jasmine-cli)
![GitHub last commit](https://img.shields.io/github/last-commit/ex1d3/jasmine-cli)

Currently, Jasmine is unreleased and unstable. There is some goals to achive release:

- [ ] Filtering by fields for "get" and "del" commands
- [ ] Output formatter for table-like outputting of entities
- [ ] Output modificators like "sum()" (e.g - get tx * -> sum(amount);)

## Installing
These intructions will help you get a copy of Jasmine on your local machine

### Prerequisites
Copy a repository
```
git clone https://github.com/ex1d3/jasmine-cli.git
```

### Installing
Running building process
```
go build
```

Starting application itself
```
./jasmine-cli
```

## Running tests
### Classical tests
```
go test ./...
```

## Built With
- [Go](https://go.dev/) - Programming Language
