# genfrisian

A program that generates random Frisian names. Some may be real names, most are just random gibberish that may sound like a real name.

To use it, simply do `genfrisian.exe [options]` on Windows or `genfrisian [options]` on Linux,

### List of options

```
-f, --first <bool>	Whether the program should generate a first name (default: true)
-s, --second <bool>	Whether the program should generate a second name (default: random)
-l, --last <bool>	Whether the program should generate a last/surname (default: true)
-n, --number <number>	How many names should be generated (default: Generates new name when 1 is pressed)
```
Use the --help option to view this at any time (e.g. `genfrisian.exe --help` on Windows)

### Examples

`genfrisian -f true -s false -l false -n 5` will generate 5 names consisting of a first name only.

`genfrisian -s true -n 10` will generate 10 names consisting of a first name, a second name, and a last name.

`genfrisian -l false` will generate a single name consisting of a first name in all cases and a second name in 50% of cases, and will generate more on every press of the 1 key.
