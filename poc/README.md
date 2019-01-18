# POC

## Who to run it ?
1. Install test program: `go install ./ptest`
1. Install pass POC: `go install .`
1. Run it:
   1. `poc`
   1. `poc unkown_binary`
   1. `poc ptest --echo hello`


## POC Roadmap
* [x] `MethodType_FlagValue`: The password is passed in clear in value of the flag. (eg. `ptest -p {password}`)
* [ ] `MethodType_FlagPath`: The password is in a file and the flag value is the file path. (eg. `ptest -p /etc/var/{password}`)
* [ ] `MethodType_Stdin`: The password given in stdin. (eg. `echo {password} > ptest`)
* [ ] `MethodType_Varenv`: The password given in a varenv. (eg. `PASS={password} ptest`)
