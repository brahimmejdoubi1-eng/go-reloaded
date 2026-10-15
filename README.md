# go-reloaded

A small text completion / editing / auto-correction tool written in Go.

It reads a text file, applies a set of formatting and conversion rules, and writes the corrected text to an output file.

## Usage

```bash
go run . <input file> <output file>
```

Example:

```bash
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

The program prints a usage message if it does not receive exactly two arguments. It refuses to write to an output file ending in `.go`, so you cannot overwrite your own source code by mistake.

## Rules

| Rule | Example input | Example output |
|---|---|---|
| `(hex)` converts the previous word from hexadecimal to decimal | `1E (hex) files` | `30 files` |
| `(bin)` converts the previous word from binary to decimal | `10 (bin) years` | `2 years` |
| `(up)` uppercases the previous word | `go (up)` | `GO` |
| `(low)` lowercases the previous word | `SHOUTING (low)` | `shouting` |
| `(cap)` capitalizes the previous word | `bridge (cap)` | `Bridge` |
| `(up, N)`, `(low, N)`, `(cap, N)` apply the change to the previous N words | `so exciting (up, 2)` | `SO EXCITING` |
| Punctuation `. , ! ? : ;` sticks to the previous word and is followed by one space | `there ,and` | `there, and` |
| Groups of punctuation such as `...` or `!?` stay together | `thinking ... You` | `thinking... You` |
| A pair of `'` hugs the words between them | `' awesome '` | `'awesome'` |
| `a` becomes `an` before a vowel or `h` | `A amazing rock` | `An amazing rock` |

## Project structure

```
go-reloaded/
├── main.go       # reads the input file, runs every rule on each line, writes the output
├── Hex.go        # (hex) rule
├── Bin.go        # (bin) rule
├── ToUpper.go    # (up) and (up, N) rules
├── ToLower.go    # (low) and (low, N) rules
├── Cap.go        # (cap) and (cap, N) rules
├── Articles.go   # a -> an rule
├── Quotes.go     # ' ' rule
├── Punct.go      # punctuation spacing rule
├── go.mod
└── README.md
```

### `main.go`

1. Checks that exactly two arguments were given, and that the output file does is .txt.
2. Reads the input file.
4. Splits the text into lines and runs every rule on each line, so line breaks and blank lines are kept.
5. Joins the lines back together and writes the output file.

### Rule files

| File | Function | What it does |
|---|---|---|
| `Hex.go` | `Hex` | Handles `(hex)`: parses the previous word in base 16 and writes it in base 10 |
| `Bin.go` | `Bin` | Handles `(bin)`: same idea in base 2 |
| `ToUpper.go` | `ToUpper` | Handles `(up)` and `(up, N)` |
| `ToLower.go` | `ToLower` | Handles `(low)` and `(low, N)` |
| `Cap.go` | `Cap`, `Capitalize` | Handles `(cap)` and `(cap, N)`; `Capitalize` uppercases the first letter of a word and lowercase teh rest |
| `Articles.go` | `Articles`, `IsVowel` | Changes `a` or `A` to `an` or `An` when the next word starts with a vowel or `h` |
| `Quotes.go` | `Quotes` | Glues each pair of `'` onto the first and last word between them |
| `Punct.go` | `Punct`, `IsPunc` | Scans the text character by character and fixes the spacing around punctuation |

## How it works

**Word-based rules.** `Hex`, `Bin`, `ToUpper`, `ToLower`, `Cap`, `Articles` and `Quotes` split the line into words with `strings.Fields` and work on the resulting slice. When a marker such as `(up)` is found at index `i`, the word at `i-1` is transformed and the marker is removed from the slice with `append(r[:i], r[i+1:]...)`. The `(x, N)` versions read the number from the next word (for example `2)`), remove the `)` with `strings.CutSuffix` and gives a bool to check if the `)` was there , convert it with `strconv.Atoi`, and transform the previous N words. The loop uses an index instead of `range` because the slice shrinks as markers are removed.

**Character-based rule.** `Punct` builds a new string one character at a time. When it reaches a punctuation mark it removes any trailing spaces from the result, adds the mark, and adds a space after it if the next character is not already a space. For groups like `...`, each new mark removes the space the previous one added, so the group stays glued and a single space follows the last mark.

**Order of rules.** `main.go` applies the rules in this order:

1. `Hex`
2. `Bin`
3. `ToUpper`
4. `ToLower`
5. `Cap`
6. `Articles`
7. `Quotes`
8. `Punct`

Markers are processed first so they are removed before word-level and spacing rules look at the text. `Punct` runs last because it only fixes spacing.

## Edge cases handled

- A marker at the very start of a line, with no word before it, is removed instead of crashing.
- `(up, N)` with N larger than the number of words before it is limited to the words that exist.
- A negative `N` is treated as zero.
- A marker with no number after it, such as a trailing `(up,`, does not crash.
- `a` as the last word of a line does not crash.
- An empty pair of quotes, or several pairs of quotes on one line, does not crash.
- Empty strings are checked before reading the first character.

## Known limitations

- The code reads bytes, not full Unicode characters, so it is designed for plain ASCII text. Emoji and accented letters may not be handled correctly.
- Rules work on one line at a time, so a marker or a pair of quotes cannot refer to words on another line.
- `(hex)` and `(bin)` assume the previous word is a valid number, as the spec says.

## Testing

Run the program on small sample files and compare the result with the expected output. The official examples:

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```
becomes
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
becomes
```
Simply add 66 and 2 and you will see the result is 68.
```

```
There is no greater agony than bearing a untold story inside you.
```
becomes
```
There is no greater agony than bearing an untold story inside you.
```

```
Punctuation tests are ... kinda boring ,what do you think ?
```
becomes
```
Punctuation tests are... kinda boring, what do you think?
```

Use `cat -e result.txt` to see line endings and check that blank lines are kept.

## Good practices followed

- `gofmt` formatting
- Small functions with one job each
- Errors checked when reading and writing files
- Bounds checks before reading `r[i-1]` and `r[i+1]`
