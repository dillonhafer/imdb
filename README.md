# IMDB

A simple and automatic way to set ID3 info tags for movies in iTunes, including artwork.

This program gets IMDB info from [The Open Movie Database](http://www.omdbapi.com/) for use with [AtomicParsley](http://atomicparsley.sourceforge.net/).

### Dependencies

Right now you need to download and have [AtomicParsley](http://atomicparsley.sourceforge.net/) in your path.

### Usage

This assumes there's a file called "DOLPHINE_TALE.m4v" in the current directory.

`imdb -id="tt1564349" -file="DOLPHIN_TALE"`

You can use the `-format="mp4"` flag to change the default format.

**Verbos usage**

```bash
Usage of imdb:
  -file="": Path to video file
  -format="m4v": File format of video file (defaults to m4v)
  -id="": IMDB ID of movie (e.g. tt1564349)
```
