# POC

## Who to run it ?
1. `go install ./ptest`
1. `ln -s $GOPATH/bin/ptest $GOPATH/bin/ptest_flagvalue`
1. `ln -s $GOPATH/bin/ptest $GOPATH/bin/ptest_flagpath`
1. `ln -s $GOPATH/bin/ptest $GOPATH/bin/ptest_varenv`
1. `./scripts/bundle_plugins.sh`
1. `go install ./pass`
1. Run it:
   1. `pass`
   1. `pass unkown_binary`
   1. `pass ptest_flagvalue --echo hello`
   1. `pass ptest_flagpath --echo hello`
   1. `pass ptest_varenv --echo hello`


## POC Roadmap
* [x] `FlagValue`: The password is passed in clear in value of the flag. (eg. `ptest -p {password}`)
* [x] `FlagPath`: The password is in a file and the flag value is the file path. (eg. `ptest -p /etc/var/{password}`)
* [ ] `Stdin`: The password given in stdin. (eg. `echo {password} > ptest`)
* [x] `Varenv`: The password given in a varenv. (eg. `PASS={password} ptest`)
