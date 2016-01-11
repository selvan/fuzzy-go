## ENV Setup
* cd fuzzy-go
* export GOPATH=`pwd`
* export PATH=$PATH:$GOPATH/bin

## Running Tests
go test -v ./src/selvan.github.com/...

##  Fuzzy Logic
- [x] score.go : accepts two strings as input and computes score based on how close those two strings are.
- [x] filter.go : accepts array of strings (source) and search string, compute score between each string in the array and search string. Sort results based on score in descending.
- [] match : Need to highlight what chars are matched in final results.

## Usage
See tests
