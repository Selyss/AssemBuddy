# AssemBuddy
Command line tool for querying assembly information. Works fully offline.

![demo](https://github.com/Selyss/AssemBuddy/assets/99344963/987e16a2-ce27-4c27-b360-6327232882e2)

## Features

- All query results are printed in a table
- Query results are `$PAGER` friendly
- Query an architecture syscall
- Query a specific syscall over all architectures
- Query a list of all syscalls for a specific architecture
- Pretty print JSON result


## Installation

The recommended installation method is building from source \
You can add the binary to `$PATH` to use it from anywhere on your system

## Examples

### Querying a specific syscall

![demo](https://github.com/Selyss/AssemBuddy/assets/99344963/987e16a2-ce27-4c27-b360-6327232882e2)

### Querying a specific syscall over all architectures

![all](https://github.com/Selyss/AssemBuddy/assets/99344963/4b7f82e4-b2a0-4791-988a-0d23d7b8b314)

### Querying all syscalls for an architecture

![arch](https://github.com/Selyss/AssemBuddy/assets/99344963/d8218e8d-233a-47e0-8ff6-9269b96de6e5)


Big thanks to [syscall.sh](https://syscall.sh/) for such an awesome API.