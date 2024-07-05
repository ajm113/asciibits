# asciibits

*simple cli converts string of decimals to ASCII string or vice versa*

Requires Go 1.22.3 or newer.

## Why?

Tired of manually copy/pasting decimals spat out from Go libraries and then converting em to human readable strings to easily debug. This is a small utility project I don't plan adding much unless I find a really good reason to add more features.

## How to Use

```sh

asciibits [--separator=string]

Description:
    asciibits - used to convert array of decimals in a string to human readable ascii characters or vice versa

Options:
    -s, --separator     Separator to use when parsing string

Sub-commands:
    asciibits ascii     convert string of decimals to ascii characters., shortcut: a
    asciibits decimal   convert string of ascii characters to array of decimals., shortcut: d
```


### Decimals to ASCII


```sh
$ asciibits a "104 101 108 108 111 44 32 119 111 114 108 100"
hello, world
```

### ASCII to Decimals

```sh
$ asciibits d "hello, world"
104 101 108 108 111 44 32 119 111 114 108 100
```

*NOTE:* If you want to change the separator at any time use `-s` flag.

## How to Install

```sh
go install github.com/ajm113/asciibits
```

## How to Build

```sh
git clone https://github.com/ajm113/asciibits.git
cd asciibits
go build .
```