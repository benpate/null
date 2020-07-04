# null 🚫


[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/benpate/null)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/null?style=flat-square)](https://goreportcard.com/report/github.com/benpate/null)
[![Build Status](http://img.shields.io/travis/benpate/null.svg?style=flat-square)](https://travis-ci.org/benpate/null)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/null.svg?style=flat-square)](https://codecov.io/gh/benpate/null)

## Simple library for null values in Go

This library provides simple primitives for nullable variables.  Int, Bool, and Float are supported.

```
	var b null.Bool // This is null

	b.Set(false) // This is false
	b.Set(true) // This is true
	b.Unset() // This is null again.
```

## Pull Requests Welcome

Please use GitHub to make suggestions, pull requests, and enhancements.  We're all in this together! 🚫