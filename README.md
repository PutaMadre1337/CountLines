# About
A small Cli-Utility that will calculate for you the number of lines in all files and directories starting from your current directory.
You can use `alias countlines='pwd && echo "Total number of lines: $(cat $(fd -t file) | wc -l)"'`, and it can work even faster, but my program has some svistoperdelki, for example:
- you can only count lines in files that match the desired extension.
- the program can display a table, list or tree of files with the number of lines and their size.
- by default hidden files are not taken into account by the program, but the `-h` flag will fix this.

# Installation
## Install binary file from github release
See: [Releases](https://github.com/PutaMadre1337/tammy/releases)

## Sh script for linux
Build binary and copy it to your /usr/bin (<u>using sudo</u>).
```sh
git clone https://github.com/PutaMadre1337/tammy && cd tammy && chmod +x installation.sh && ./installation.sh
```

## Using Make
```sh
git clone https://github.com/PutaMadre1337/tammy && cd tammy && make build
```

# Configuration
```sh
export DEFAULT_FORM="table"
export ALLWAYS_DISPLAY_SIZE="false"
export ALLWAYS_SHOW_HIDDEN_FILES="false"
export LIST_ENUMERATOR="Roman"
export TREE_ENUMERATOR="RoundedEnumerator"
```

# Usage
Display information about files in the current directory:
```sh
tammy
```

## Flags
- `tammy -f`:
  + Change output format. Available forms: table, list, total, tree (default - table).
- `tammy -h`:
  + Show hidden files.
- `tammy -s`:
  + Show file size.
- `tammy -p`:
  + Specify the path to the directory in which to count lines. It is not necessary to specify the path. The path can also be specified at the very end: `tammy -f list -s -h ~/Documents`.
- `tammy -ft`:
  + Count lines only in files with a certain extension. Example: `tammy -ft md`, or `tammy -ft .md`.
- `tammy -t`:
  + Show execution time.
- `tammy -help`:
  + Show help message.
- `tammy -version`:
  + Check for updates.

# Configuring
The utility is configured using environment variables. Available variables:
- `DEFAULT_FORM`
  + Available values:
    + `export DEFAULT_FORM="table"` - default
    + `export DEFAULT_FORM="list"`
    + `export DEFAULT_FORM="tree"`
    + `export DEFAULT_FORM="total"`
- `ALLWAYS_DISPLAY_SIZE`
  + Available values:
    + `export ALLWAYS_DISPLAY_SIZE="false"` - default
    + `export ALLWAYS_DISPLAY_SIZE="true"`
- `ALLWAYS_SHOW_HIDDEN_FILES`
  + Available values:
    + `export ALLWAYS_SHOW_HIDDEN_FILES="false"` - default
    + `export ALLWAYS_SHOW_HIDDEN_FILES="true"`
- `LIST_ENUMERATOR`
  + Available values:
    + `export LIST_ENUMERATOR="roman"` - default
    + `export LIST_ENUMERATOR="arabic"`
    + `export LIST_ENUMERATOR="dash"`
    + `export LIST_ENUMERATOR="alphabet"`
    + `export LIST_ENUMERATOR="bullet"`
    + `export LIST_ENUMERATOR="asterisk"`
- `TREE_ENUMERATOR`
  + Available values:
    + `export TREE_ENUMERATOR="rounded"` - default
    + `export TREE_ENUMERATOR="default_enumerator"`
    + `export TREE_ENUMERATOR="default_indenter"`

