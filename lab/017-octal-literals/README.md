# #17 Octal literals 

Edit only `excersise.go`. Run `make lab N=017` until it's green.5

1. **Predict**: fill in `Prediction`  before running anything. The failing output shows what Go actually computes.
2. **Fix** `ParseUserID`: a user typed `010` and got account 8. Find the one-argument change, then handle negative IDs.
3. **Fix** `WriteConfig`: the config get written but its owner can't read it. Explain why with `fmt.Printf(" %#o %v", 644, os.FileMode(644))`. Then delete the `/nolint` comment and run `make lint`:staticcheck (SA9002)  flags this exact bug. That's #16, linters paying off.


