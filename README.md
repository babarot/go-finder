go-finder
=========

[![GoDoc](https://godoc.org/github.com/b4b4r07/go-finder?status.svg)](https://godoc.org/github.com/b4b4r07/go-finder)

CLI finder wrapper (fzf, peco, etc) for golang

## Usage

```go
fzf, err := finder.New("fzf", "--reverse", "--height", "40")
if err != nil {
	panic(err)
}
fzf.Run()
```

```go
peco, err := finder.New("peco")
if err != nil {
	panic(err)
}
peco.Run()
```

```go
cli, err := finder.New()
if err != nil {
	panic(err)
}
// If no argument is given to finder.New()
// it scans available finder command (fzf,fzy,peco,etc) from your PATH
```

## Installation

```console
$ go get -d github.com/b4b4r07/go-finder
```

## License

MIT

## Author

b4b4r07
