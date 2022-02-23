## Work in progress

## bulding

In a linux environment run:

```sh
make
```
## running

After building run:

```sh
bin/grop
```

## Usage

| argument       | short | description                                          |
| -------------- | ----- | ---------------------------------------------------- |
| --help         | -h    | Print help information                               |
| --filename     | -f    | Filename to search within the file system (Optional) |
| --substring    | -h    | String to look for                                   |
| --count        | -c    | Count substring occurrences                          |
| --pattern      | -p    | Regex to search within the file                      |


## TODOs

- [x] Implement logic from reading from file
- [x] Implement colored highlight
- [x] Implement line numbers
- [x] Implement better error messages/warnings
- [ ] Implement search with regex
- [ ] Implement logic for piped output (such as `cat $SOURCE | grep ...`)
- [ ] Implement go routines to search big file, necessary refactoring logic for read file by lines.
