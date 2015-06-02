# IMDB Tags

A simple and automatic way to set ID3 tags from IMDB, including artwork.

This program gets IMDB info from [The Open Movie Database](http://www.omdbapi.com/) for use with [AtomicParsley](http://atomicparsley.sourceforge.net/).

### Dependencies

Right now you need to download and have [AtomicParsley](http://atomicparsley.sourceforge.net/) in your path.

You can download AtomicParsley with: `imdb-tags atomic`

### Usage

This assumes there's a file called "DOLPHINE_TALE.m4v" in the current directory.

`imdb-tags -i tt1564349 dolphin_tale.m4v`

If you don't want to open a browser and look for the IMDB Id manually, you can search for it on the command line.

`imdb-tags search "Dolphin Tale"`

This makes it convient to run commands like:

```bash
#!/bin/bash
movie="dolphin_tale.m4v"
HandBrakeCLI -i /Volumes/DVD/VIDEO_TS -o $movie && imdb-tags -i tt1564349 $movie
```
