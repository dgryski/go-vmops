go-vmops
========

Go's `regexp` package is sufficient for many use cases.  However, in cases
where performance in paramount, switching to code generation can offer a
significant speed increase. See for example, this article on [speeding up
regexs][medium] using the [ragel state machine compiler][ragel].

However, large goto-based matchers, while significantly faster on native
platforms, run into trouble on WebAssembly due the limitations of structured
control flow.  Compilers can take an exceedingly long time to turn the large
generated code into something without any gotos.

The VMOPS output format from [libfsm][libfsm] generates a large table of
opcodes for a regular expression matcher that can be evaluated with much
smaller inner loop.  While slower than generating native code, it has the
advantage that it can be used successfully on WebAssebmly.  Due to libfsm's
DFA-based matcher it can still end up being faster than Go's native `regexp`
package.

The opcodes can also be serialized to a binary format to use with
`go:embed` directives or loading from files at runtime with any additional
parsing overhead.

[ragel]: http://www.colm.net/open-source/ragel/
[medium]: https://dgryski.medium.com/speeding-up-regexp-matching-with-ragel-4727f1c16027
[libfsm]: https://github.com/katef/libfsm
