# check-me-out
A code kata for a basic checkout service

To test the code simply run `go test ./...` as per usual.


The run the cli tool you can compile the code with `go build .`
The executable can then be run in the normal way.

The tool accepts 3 argument:
1. The file path of the JSON file you wish to read in - this is the store you will be using.
Examples of what this file looks like can be found in the testFiles directory.
2. The items you want to scan. This is a string that should look like "AAABB". The program will scan each character individually.
3. The basket total that you think is correct. This will be compared to the tools answer and displayed to you.

An example command that should be successful is `go run . ./testFiles/basic.json AAABB 175`
