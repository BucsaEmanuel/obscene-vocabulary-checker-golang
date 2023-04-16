# Hyperskill Go - Obscene Vocabulary Checker Project

- [Hyperskill Go - Obscene Vocabulary Checker project](#hyperskill-obscene-vocabulary-checker-project)
    - [About the project](#about-the-project)
        - [Status](#status)
        - [See also](#see-also)
    - [Getting started](#getting-started)

# Hyperskill Go - Obscene Vocabulary Checker project

## About the project

This is the implementation of the Go Obscene Vocabulary Checker project on Hyperskill.

### Status

The project is posted as it currently is, having passed the last step of the project ( 4 / 4 ).

### See also

* [Hyperskill - Obscene Vocabulary Checker (Go)](https://hyperskill.org/projects/201)

## Getting started

Either
`go build main.go`
then `./main`

or
`go run main.go`

## Examples
The greater-than symbol followed by a space (> ) represents the user input. Note that it's not part of the input.

### Example 1:
```
The forbidden_words.txt contents:
disgusting
unpleasant
ugly
bad
```

```
> forbidden_words.txt
> It is a Bad and ugly sentence.
It is a *** and **** sentence.
> It is a Badge.
It is a Badge.
> exit
Bye!
```