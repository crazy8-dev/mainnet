#!/usr/bin/env sh
# filters coverage stats for code that could be tested, but should not affect coverage metric
#
# * generated code (mock and stringer)
# * command line tools code
grep -v "_mock.go:" | \
    grep -v "_string.go:" | \
    grep -v "_gen.go:" | \
    grep -v 'github.com/insolar/mainnet/cmd/'
