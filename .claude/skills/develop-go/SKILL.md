---
name: develop-go
description: use this skill when an implementation task requires you to write code in Golang.
---

# Guiding principles

- Always use the standard library when possible. Avoid introducing third party dependencies unless absolutely necessary.
- Follow the standard Go formatting and style guidelines.
- Use Makefile's for build entrypoints. There should be standard tasks for generate, lint, test, and build.
- For testing, use the standard `testing` package.
- Write unit tests next to the code they are testing, in a file with the same name but with `_test` suffix. For example, if you have `foo.go`, the tests should be in `foo_test.go`.
- For mocks, we generate mock classes using mockgen. For all interfaces, add a comment with the `go:generate` directive to generate the mock for that interface. Mocks should be generated under a `mocks` directory next to the interface definition. For example:

```go
//go:generate mockgen -source=path/to/interface.go -destination=path/to/mock.go -package=packageName
```

- When testing HTTP calls, prefer using `httptest` to mock out the HTTP server, instead of mocking out the HTTP client.
- For assertions, we use `gomega`, along with its `gomega.NewWithT` function to create a new `Gomega` instance in each test. Always dot-import the `gomega` package in test files to make the assertions more concise. For example:

```go
import (
    . "github.com/onsi/gomega"
)

func TestSomething(t *testing.T) {
    g := NewWithT(t)

    // your test code here

    g.Expect(something).To(Equal(expectedValue))
}
```

- Where possible, use table driven tests to reduce boilerplate and make it easier to add new test cases. For example:

```go
func TestSomething(t *testing.T) {
    g := NewWithT(t)

    tests := []struct {
        name     string
        input    interface{}
        expected interface{}
    }{
        // test cases here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // your test code here
            g.Expect(actual).To(Equal(tt.expected))
        })
    }
}
```

- For linting, we use `golangci-lint`, however this should always be accessed through the Makefile task. You should also prefer using the `--fix` option before trying to fix errors yourself.
- Configuration parameters for an application should be parsed from the CLI flags, using the `flag` package.
- For parameters that shouldn't be provided by flags, such as secrets, read them from environment variables using go-envconfig to safely load them onto a struct. All environment configuration should be stored in a single struct, and there should be a single function that loads all the configuration for the application.