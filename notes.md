# Notes
## Getting Started
- Create a directory
  ```
  $ mkdir hello-go
  ```
- Create a new file
  ```
  $ touch hello.go
  ```
- We need a mod file to track dependencies
  ```
  $ go mod init example/hello
  ```
  - When the application is in production this module path would point to where the source code is stored. e.g. `github.com/alexmacniven/hello`
- After writing some code we can run it
  ```
  $ go run .
  ```
## Call Code in an External Package
- Visit pkg.go.dev and find a package to use, i.e. rsc.io/quote
  - The package's documentation will tell you what functions the package supports
- Add the package to the our code and tidy
  ```
  $ go mod tidy
  ```
