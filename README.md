# Mandelbrot Set

<p align="center">
  <img src="/sample.png?raw=true" alt="Sample output from this program" width=350/>
</p>

Based entirely on [Daniel Shiffman's video][], this program generates an image
of a Mandelbrot set using OCaml instead of JavaScript.

I do plan on improving this program by implementing other image-generating
algorithms in the future, but I'm not sure when.

## Compilation and Usage

```sh
$ oasis setup
$ ocaml setup.ml -configure
$ ocaml setup.ml -build
```

[Daniel Shiffman's video]: https://www.youtube.com/watch?v=6z7GQewK-Ks
[kingpin]: https://github.com/alecthomas/kingpin
