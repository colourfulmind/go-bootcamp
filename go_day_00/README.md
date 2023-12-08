## Statistical metrics.

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/go_day_00/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/go_day_00/internal/app/run)
- [Internal Package Calculations](http://localhost:6060/pkg/go_day_00/internal/app/calculations)
- [Internal Package Run](http://localhost:6060/pkg/go_day_00/internal/app/run)

There is an app that calculates four statistical metrics from a bunch of integer numbers (strictly between -100000 and 100000) provided by user. The numbers are being read from a standard input, separated by newlines.

There is an example of the output:  
> **Mean:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;323.48**  
> **Median:&nbsp;&nbsp;54.00**  
> **Mode:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;21**  
> **SD:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;881.77**

By default, all of them are being printed, but you can specify which one you need by using flags:

| Flag | Description |
|------|-------|
| mean | an average number of a given sequence |
| median | a middle number of a sorted sequence, if there are several, the smallest one is printed |
| mode | a number which is occurring most frequently |
| sd | a regular standard deviation |
