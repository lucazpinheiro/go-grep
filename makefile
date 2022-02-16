# Default go compiler
GC := go

# Entry point file
ENTRY := main.go

# Output file
OUTPUT := dist/grop

# Makefile flags
MAKEFLAGS += "--silent"

# Script to generate binary
# output file.
build:
	echo "Generating binary file..."
	cd src && $(GC) build -o ../$(OUTPUT)
	echo "Output file was succesfully generated."

# Script to remove generated
# output files.
clean:
	rm -rf bin