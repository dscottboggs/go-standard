# System Exits

Gives names to the various exit codes that have standard values according to [sysexits.h](https://gist.githubusercontent.com/bojanrajkovic/831993/raw/79d07934534ba03d1b21c78917b9a8b699d8d6fe/sysexits.h) and common usage.

## Usage
##### When it comes time to exit a program

```go
import github.com/dscottboggs/go-standard/exits

exit.With(exit.Success)
// or
exit.With(exit.Fatal(exit.PermissionDenied))
exit.With(exit.NoSuchFile)
```
