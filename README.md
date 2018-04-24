go-finder
=========

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
peco, err := finder.New("peco", "--layout=bottom-up")
if err != nil {
	panic(err)
}
peco.Run()
```

You can set the finder command a little more flexibly like the following script.
By default the data source is Stdin, but you can choose a variety of it.

```go
var opts []string
command := finder.Command()
switch command {
case "fzf":
	opts = []string{
		"--reverse",
		"--height", "40",
	}
case "peco":
	opts = []string{
		"--layout=bottom-up",
	}
}

cli, err := finder.New(command, opts...)
if err != nil {
	panic(err)
}

// You can select the data source to use as filter source
cli.FromFile("some-file.txt")
cli.FromText("sample\ntext\nfoo")
cli.FromCommand("cat", "some-file.txt")
cli.FromStdin() // default

items, err := cli.Run()
if err != nil {
	panic(err)
}
fmt.Printf("%#v\n", items)
```

## Installation

```console
$ go get github.com/b4b4r07/go-finder
```

## License

MIT

## Author

b4b4r07
