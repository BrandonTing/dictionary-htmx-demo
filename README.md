# Project dictionary-htmx-demo

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

## Sharing Topic
### Htmx
- Concept
- Attributes
  - hx-get/hx-post/hx-put/hx-delete
  - hx-target
  - hx-swap
  - hx-swap-oob
  - loading indicator
- When
  - Server Driven Projects -> CMS/Dashboards...
  - Internal tools

### Steps
- Install templ to compile templ files: `go install github.com/a-h/templ/cmd/templ@latest`
- 