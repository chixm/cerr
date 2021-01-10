# cerr
chixm errors for go

## Very Simple Error with Error line.
cerr holds the line which New method or Wrap method is called.
So, it does not prints out whole print stacks, but enough to know where and what is going wrong.
It helps decreasing error logging lines and saves us from writing a lot of logs to log file. 

### New method only writes the error of file&line where New is declared.
```
err := New(`original error`)
```

This code prints only
```
cerr_test.go[10] original error
``

