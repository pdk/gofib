# gofib

A response to 
[this reddit thread](https://www.reddit.com/r/javascript/comments/lg80p2/nodejs_14_is_over_20x_faster_than_python38_for/gmq2s0b/)
about node being faster than python. Go got thrown in the pot, as well. I
thought this was not exactly a far comparison as the strength of go is the
ability to *easily* add concurrency. So I wrote a little thing to compare a bare
recursive version of fibonacci in go to a concurrent version.

This does not do anything fancy to spawn zillions of goroutines. In fact it
should spawn only 8, which is the number of cores on my machine.

    pkelly-macOS:gofib pkelly$ ./gofib 35
    normal  : 14930352 59.803694ms
    parallel: 14930352 17.234335ms
    pkelly-macOS:gofib pkelly$ ./gofib 40
    normal  : 165580141 585.708817ms
    parallel: 165580141 176.313554ms
    pkelly-macOS:gofib pkelly$ ./gofib 45
    normal  : 1836311903 6.203504596s
    parallel: 1836311903 1.935910803s
    pkelly-macOS:gofib pkelly$ ./gofib 46
    normal  : 2971215073 9.958304106s
    parallel: 2971215073 3.26121279s
    pkelly-macOS:gofib pkelly$ ./gofib 47
    normal  : 4807526976 16.909490196s
    parallel: 4807526976 5.58676535s
