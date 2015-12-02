# Bang

**Bang** is a command line tool for quick storage and retriveal of text snippets. It is a go implementation of this thing: https://github.com/jimmycuadra/bang

## Installation

## Get help

    $ bang -h

## Add a snippet

    $ bang jimmy http://jimmycuadra.com/

## Retrieve a snippet

    $ bang jimmy
    http://jimmycuadra.com/

## List all snippets

    $ bang
      jimmy: http://jimmycuadra.com/
       mtnt: http://morethingsneed.to/
    address: 237 Overlook Street

## Storage

Bang's data is serialized to JSON and peristed to a file at `~/.bang`. If you want to share your Bang data across multiple machines, you can move the file to [Dropbox](https://www.dropbox.com/) or some other type of network disk, and then symlink it to `~/.bang` on your local machine.

