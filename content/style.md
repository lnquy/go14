# Go Style

The documents, [Effective Go](https://golang.org/doc/effective_go.html) and [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) should be used as style-guides for Go. Additions, modifications or exceptions to these documents should be listed below.

# Documentation

## General

The repository should contain a “README.md” file at the root of the repository written in markdown, that should explain all the steps necessary for a new developer to:

- Understand the general purpose of the program

- Run the test suite

- Understand the possible configuration options (if any)

- Install required dependencies (if any)

 

Multiple “REAMDE.md”-files is allowed in the case where the point is to document some other folder with non-Go related files.

Example:

```tex
conf/
	debug.conf
	README.md
main.go
README.md
```



## Code

All Go-related code should be documented with the use of comments. The tool “godoc” will generate developer documentation based off these comments and is the main way we should document Go-code.

 

Further reading and examples:

- <https://golang.org/doc/effective_go.html#commentary>

- <https://blog.golang.org/godoc-documenting-go-code>



Currently, the style-guide already defines a few points about how comments should be applied to the code:

- <https://github.com/golang/go/wiki/CodeReviewComments#comment-sentences>

- <https://github.com/golang/go/wiki/CodeReviewComments#doc-comments>

- <https://github.com/golang/go/wiki/CodeReviewComments#package-comments>

### Summary and enforced requirements

- All comments documenting declarations should be full sentences

- Comments should begin with the name of the thing being described and end in a period

- All top-level, exported names should have doc comments, as should non-trivial unexported type or function declarations

- Every package should have a package comment

- Should a package need large amounts of introductory documentation, separate it out into its own file called “doc.go”

# Testing

All tests should use Go’s internal “testing” package:

<https://golang.org/pkg/testing/>

If an assertion library is needed, we have currently opted in for using a lightweight library called “Goblin”:

<https://github.com/franela/goblin> 

Tests should fail with helpful messages saying what was wrong, the input provided to the function being tested, the actual output and the expected output

<https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures>

Example:

```go
if got := fn(input); got != expected {
	t.Errorf(“Func(%q) = %d; expected %d”, input, got, expected)
}
```


Separation of unit- and integration-tests should be done using build tags. All integration tests should be marked with the build tag “integration”:
```go
// +build integration
package notify
```

To run these tests, you will need to pass the tags when invoking the “go test” command: 
```shell
go test -tags=integration
```

# Dependencies

Applications written in Go (go source code that includes a main.go file) should list it’s vendor dependencies. To aid this process, we’ve chosen to use “[Glide](http://glide.sh/)” to manage the dependencies that gets vendor. http://glide.sh , the latest version of glide can be downloaded from <https://github.com/Masterminds/glide/releases>. The glide documentation can be found here: <https://glide.readthedocs.io/en/latest/>.
