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
| --filename     | -f    | Filename to search within the file system (Optional) |
| --help         | -h    | Print help information                               |
| --line-numbers | -l    | Optional value to show line numbers or not           |
| --pattern      | -p    | Regex to search within the file                      |


## TODOs

- [x] Implement logic from reading from file
- [ ] Implement colored highlight
- [ ] Implement logic for piped output (such as `cat $SOURCE | grep ...`)
- [x] Implement line numbers
- [ ] Implement search with regex
- [x] Implement better error messages/warnings
- [ ] Implement go routines to search big file, necessary refactoring logic for read file by lines.
