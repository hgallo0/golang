# golang

Variable definition
var fileName // Camel Case

var FileName // Pascal Case

# Testing
## Unit Tests
Unit tests cover small parts of a code base and test these in isolation. Often, this might be a single function, and the inputs and outputs of a function are tested. A typical unit test might say, “If I give function x these values, I should expect this value to be returned.” These types of tests are extremely useful to confirm that the smallest building blocks of a program function in the way that is expected. As a program grows and changes, unit tests are an excellent way to catch any regressions. A regression is a bug or fault that has been introduced as a result of a change. Regressions mean that code was working before the change, but not after it. Unit tests can often catch regressions as they test the smallest parts of a program.

## Integration Tests
Integration tests typically test how the various parts of an application work together. If unit tests verify the smallest parts of a program, integration tests look at how the components of an application work together. Integration tests also test things like network calls and database connections, verifying that the system as a whole works as expected. Typically, integration tests are more complex and difficult to construct than unit tests, as the tests need to assess dependent parts of an application.

## Functional Tests
Functional tests are often known as end-to-end tests and outside-in tests. These tests verify that software works as expected from an end-user perspective. They assess how a program works from the outside without being concerned with how the software works internally. For users of software, functional tests are perhaps the most important tests that can be run. Examples of functional tests include:

Image Testing that a command line tool responds to certain inputs with certain outputs.

Image Running automated tests on a web page.

Image Running outside-in tests against an API to check response codes and headers.

## Test-Driven Development
Many developers advocate using test-driven development (TDD). This is the practice of thinking about a new feature in terms of a test. Before any code is written, a test is created that describes the expected functionality of a piece of code. This has many benefits.

Image It can help to inform code design; thinking clearly about how a piece of code works can often improve the design.

Image It can help to provide definition about how a feature should work.

Image There is a test that will verify there are no regressions in the future.

Image There is a test to verify that the code is correctly implemented.

Using TDD software, engineers can improve code design and confirm that the code is functional by ensuring that the test passes.
