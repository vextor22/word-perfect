# word-perfect
Tools to generate word lists containing some variation of 7 letters, with one required letter, for a word game.

# Build

go build .


# Run
To run, choose one of the functions: 
* filter
* generate

## Filter
Filter takes an input file by path, or reads from stdin. It will interpret the text and generate a list of unique single words per line.

## Generate
Generate takes a single string as an argument. The first character in the string is the required character, all others are possible characters.

Generate will then return a set of valid words for the input condition and output of Filter.

# Example

cat dictionary.txt | go run . filter | go run . generate cranes
