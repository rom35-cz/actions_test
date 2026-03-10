# Actions Test

[![PR Tests](https://github.com/rom35-cz/actions_test/actions/workflows/pr_test.yml/badge.svg)](https://github.com/rom35-cz/actions_test/actions/workflows/pr_test.yml)
[![Release Tests](https://github.com/rom35-cz/actions_test/actions/workflows/release.yml/badge.svg)](https://github.com/rom35-cz/actions_test/actions/workflows/release.yml)

A learning project for Go and GitHub Actions CI/CD. This repository demonstrates automated testing workflows triggered on pull requests and semantic version tags.

## Prerequisites

- [Go](https://golang.org/dl/) 1.24.1 or later

## Installation

### Using go install

```bash
go install github.com/rom35-cz/actions_test@latest
```

### From source

```bash
git clone https://github.com/rom35-cz/actions_test.git
cd actions_test
make build
```

## Usage

```bash
./actions_test
```

Output:
```
Yesterday's date is: 2026-03-09
Hello, World!
```

## Development

This project uses a Makefile for common tasks:

| Command | Description |
|---------|-------------|
| `make all` | Run fmt, vet, test, and build |
| `make build` | Compile the binary |
| `make run` | Build and execute the binary |
| `make test` | Run all tests |
| `make fmt` | Format Go code |
| `make vet` | Run go vet for static analysis |
| `make clean` | Remove build artifacts |

## CI/CD

This repository uses GitHub Actions for continuous integration:

- **PR Tests**: Runs on every pull request
- **Release Tests**: Runs when a semantic version tag (`v*.*.*`) is pushed

### Creating a Release

To trigger the release workflow, push a semantic version tag:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The release workflow will run tests to validate the tagged commit.

## License

See [LICENSE](LICENSE) file for details.
