![imgres](https://cloud.githubusercontent.com/assets/106776/11532048/0e987ea2-98f9-11e5-9608-1439a89a8607.png)

**Bang** is a command line tool for quick storage and retriveal of text snippets. It is a go implementation of this thing: https://github.com/jimmycuadra/bang


## Installation
Todo...

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

