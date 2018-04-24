go-finder
=========

CLI finder wrapper (fzf, peco, etc) for golang

## Usage

```go
fzf, err := finder.New("fzf", "--reverse", "--height", "40")
if err != nil {
	log.Fatal(err)
}
fzf.Run()
```

```go
peco, err := finder.New("peco", "--layout=bottom-up")
if err != nil {
	log.Fatal(err)
}
peco.Run()
```

## Installation

```console
$ go get github.com/b4b4r07/go-finder
```

## License

MIT

## Author

b4b4r07