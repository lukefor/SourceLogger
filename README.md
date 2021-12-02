# SourceLogger

This is a small program written in Go which redirects
the output of the `srcds_linux` executable to stdout.

It was originally intended for Garrys Mod, but this fork is intended for running TF2 SRCDS via SystemD and Screen.

## How to use

1. [Download](https://github.com/lukefor/SourceLogger/releases) the latest release from GitHub
2. Put the file `source_logger` into the same directory as `srcds_linux`
2. Add `source-logger` before `srcds_run` or `srcds_linux`.

An example start script before:
````sh
#!/bin/bash
./srcds_linux -maxplayers 16 +gamemode terrortown +map ttt_waterworld
````

A start script using SourceLogger
````sh
#!/bin/bash
./source-logger ./srcds_linux -maxplayers 16 +gamemode terrortown +map ttt_waterworld
````

All arguments to `source_logger` will be forwarded to `srcds_linux`, so you can just keep your old arguments.

## How does it work?
This program uses [creack/pty](https://github.com/creack/pty) to create a pseudo console and capture its output.
