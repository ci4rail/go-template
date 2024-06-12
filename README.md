# go-template

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

a go-template project

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Bootstrap

When moving a go project you have to adjust the module paths.

Search for `github.com/ci4rail/go-template` and replace with your new module url.

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
