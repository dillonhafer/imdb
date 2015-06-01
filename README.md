# IMDB

A simple and automatic way to set ID3 info tags for movies in iTunes, including artwork.

This program gets IMDB info from [The Open Movie Database](http://www.omdbapi.com/) for use with [AtomicParsley](http://atomicparsley.sourceforge.net/).

### Dependencies

Right now you need to download and have [AtomicParsley](http://atomicparsley.sourceforge.net/) in your path.

### Usage

This assumes there's a file called "DOLPHINE_TALE.m4v" in the current directory.

`imdb -i tt1564349 DOLPHIN_TALE.m4v`

If you don't want to open a browser and look for the IMDB Id manually, you can search for it on the command line.

`imdb search "Dolphin Tale"`
