# klip

[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat-square)](https://github.com/iromli/klip/blob/master/LICENSE)
[![Codeship Status](https://img.shields.io/codeship/17858620-82a6-0132-4836-72ed2dde02fa.svg?style=flat-square&label=codeship)](https://codeship.com/projects/57913)

Text snippets on the command line! Inspired by [boom](https://github.com/holman/boom) and [clip](https://github.com/silent1mezzo/clip). Tested on Ubuntu.

## Crash Course

### Create a list

Usage:

    klip put <list>

Example:

    klip put mylist

### Create a list item

Usage:

    klip put <list> <item> <value>

Example:

    klip put mylist myitem content goes here

### Get a list

Usage:

    klip get <list>

Example:

    klip get mylist

### Get a list item

Usage:

    klip get <list> <name>

Example:

    klip get mylist myitem

### Delete a list

Usage:

    klip delete <list>

Example:

    klip delete mylist

### Delete a list item

Usage:

    klip delete <list> <name>

Example:

    klip delete mylist myitem

## Tips & Tricks

1. Copy a list item to clipboard: `klip get links mylink | xsel -b`.
2. Open a browser with URL from a list item: `klip get links mylink | xargs xdg-open`.
3. Pipe to a pager: `klip get links | less -r`.
