[![GitHub Actions][ci-image]][ci-url]

# stringsx

The `stringsx` package is an extensions of Go's standard [`strings`] package
with a few [extra functions].

## Why

Mainly because I feel like there are some functions missing from the standard
`strings` package. Hence, this package provides those functions so that they can
be reused across projects.

This package includes the functions of the standard `strings` package so that
you never have to think about whether you need to use `strings` or `stringsx`.
Just use `stringsx` all the time!

There are [similar projects] out there but those do not seem to provide both
the functions of the standard and (all) the extra functions found here.

## Extra Functions

The following functions are new in the `stringsx` package compared to the
`strings` package.

### `All(elems []string, condition func(string) bool) bool`

Test for every string in _elems_ if the _condition_ holds (i.e. returns `true`).
If the _condition_ holds for **every** string the return value will be `true`,
otherwise the return value will be `false`.

### `Any(elems []string, condition func(string) bool) bool`

Test for every string in _elems_ if the _condition_ holds (i.e. returns `true`).
If the _condition_ holds for **any** string the return value will be `true`,
otherwise the return value will be `false`.

### `IsEmpty(s string) bool`

Test if the string _s_ is [the empty string].

### `MapAll(elems []string, mapping func(string) string) []string`

Create a new slice populated with the results of calling the _mapping_ on every
string in _elems_.

## License

The license for this project is the same as the license for Golang
([link](https://github.com/golang/go/blob/master/LICENSE)).

[extra functions]: #extra-functions
[similar projects]: https://github.com/search?l=Go&q=stringsx&type=Repositories
[`strings`]: https://golang.org/pkg/strings
[the empty string]: https://en.wikipedia.org/wiki/Empty_string

[ci-url]: https://github.com/ericcornelissen/stringsx/actions?query=workflow%3AGo
[ci-image]: https://github.com/ericcornelissen/stringsx/workflows/Go/badge.svg