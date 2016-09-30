# Mandelbrot Set

<p align="center">
  <img src="/sample.png?raw=true" alt="Sample output from this program" width=350/>
</p>

Based entirely on [Daniel Shiffman's video][], this program generates an image
of a Mandelbrot set using Go instead of JavaScript.

I do plan on improving this program by implementing other image-generating
algorithms in the future, but I'm not sure when.

## Compilation and Usage

For now, the program uses only the standard libraries present in the default go
instalation, so a simple

```sh
$ go build
```

in the project's directory will build the executable for your machine. Of
course, you can also `go run main.go` if you do not feel like having an
executable hanging around.

Keep in mind that I have plans to implement command-line arguments (possibly
using [kingpin][]) for this program and that they may break if this project
evolves.

[Daniel Shiffman's video]: https://www.youtube.com/watch?v=6z7GQewK-Ks
[kingpin]: https://github.com/alecthomas/kingpin
