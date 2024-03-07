# tree-sitter-sagemath

[![CI][ci]](https://github.com/Eloitor/tree-sitter-sagemath/actions/workflows/ci.yml)

Sagemath grammar for [tree-sitter][].

[tree-sitter]: https://github.com/tree-sitter/tree-sitter

## Differences from the Python grammar

- [ ] Parses exponent `^` and xor `^^`
- [ ] Parses raw numbers  `939393R`. Distinguish between sage and python numbers
- [ ] Methods can be called on sage numbers `16.sqrt()` or `16J.sqrt()` but not `16R.sqrt()`
- [ ] Parses ellipsis `[1..10]`
- [ ] Parses generators `G.0`
- [ ] Parses generators constructor `ZZ.<x> = ZZ['x']`
- [ ] Symbolic functional notation `f(x) = x^2`
- [ ] Parses prompts `sage: 2 + 2` and `>>>   3 + 2`
- [ ] Parses Backslash operator `\`


## Not planned syntax support

- Implicit multiplication

## References

- [Python 3 Grammar](https://docs.python.org/3/reference/grammar.html)
- [Sagemath preparse](https://doc.sagemath.org/html/en/reference/repl/sage/repl/preparse.html)

[ci]: https://img.shields.io/github/actions/workflow/status/Eloitor/tree-sitter-sagemath/ci.yml?logo=github&label=CI
