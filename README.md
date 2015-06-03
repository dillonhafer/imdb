# IMDB Tags

A simple and automatic way to set ID3 tags from IMDB, including artwork.

This program gets IMDB info from [The Open Movie Database](http://www.omdbapi.com/) for use with [AtomicParsley](http://atomicparsley.sourceforge.net/).

### Downloads

http://imdb-tags.dillonhafer.com

### Dependencies

You must have [AtomicParsley](http://atomicparsley.sourceforge.net/) in your path.

You can open the download page with `imdb-tags atomic`

### Usage

This assumes there's a file called "song_of_the_sea.m4v" in the current directory.

`imdb-tags -i tt1865505 song_of_the_sea.m4v`

If you don't want to open a browser and look for the IMDB Id manually, you can search for it on the command line.

`imdb-tags search "song of the sea"`

This makes it convient to run commands like:

```bash
#!/bin/bash
movie="song_of_the_sea.m4v"
HandBrakeCLI -i /Volumes/DVD/VIDEO_TS -o $movie && imdb-tags -i tt1865505 $movie
```
