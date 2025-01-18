# lzcnt.space
source code repository of [lzcnt.space](https://lzcnt.space).

## building
if you modified log2.s, you shoud run build.sh, it builds deps/genhl and runs `go generate`. genhl requires `make`, `gcc` and `lex` to be installed.
otherwise you can just run `go build -o lzcnt.space`
