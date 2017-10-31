# Mandelbrot Set

<p align="center">
  <img src="/sample.png?raw=true" alt="Sample output from this program" width=350/>
</p>

Based entirely on [Daniel Shiffman's video][], this program generates an image
of a Mandelbrot set using OCaml instead of JavaScript.

[Daniel Shiffman's video]: https://www.youtube.com/watch?v=6z7GQewK-Ks

I do plan on improving this program by implementing other image-generating
algorithms in the future, but I'm not sure when.

## Compilation and Usage

To simplify things, I'm using the [oasis][] build system. You can install it
with:

[oasis]: http://oasis.forge.ocamlcore.org/

```sh
$ opam install oasis
```

After that, you'll be able to build this project with:

```sh
$ oasis setup
$ ocaml setup.ml -configure
$ ocaml setup.ml -build
```

And you should have an executable in the root directory.

To quit the program, simply press <kbd>q</kbd>.
