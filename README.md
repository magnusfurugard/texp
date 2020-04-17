# texp
Command line utility for token-based string expansion. Because sometimes you need to reliably repeat yourself.

```
Usage:
  texp "<token-string>" [flags]

Flags:
  -h, --help                help for texp
  -o, --output string       output format: json, yaml and raw (default "json")
  -t, --token stringArray   tokens to replace with format key=value
```

# Install
```
# Make sure you have $GOPATH/bin in your $PATH
go get github.com/magnusfurugard/texp
```

Or, if you've cloned the repo, you can use the `Makefile` if you want:
```
# Make sure ~/bin is in your $PATH
make install
```

# Examples
Result strings are automatically generated to cover all combinations.
```
texp "Good @time, how are you @person?" \
    -t @time=day \
    -t @time=evening \
    -t @person=friend \
    -t @person=pal \
    -o raw

# Outputs:
# Good day, how are you friend?
# Good day, how are you pal?
# Good evening, how are you friend?
# Good evening, how are you pal?

texp "select * from @db.schema.@table;" \
    -t @table=tbl1 \
    -t @table=tbl2 \
    -t @table=tbl3 \
    -t @db=dba \
    -t @db=dbb \
    -t @db=dbc \
    -o raw

# Outputs:
# select * from dba.schema.tbl1;
# select * from dba.schema.tbl2;
# select * from dba.schema.tbl3;
# select * from dbb.schema.tbl1;
# select * from dbb.schema.tbl2;
# select * from dbb.schema.tbl3;
# select * from dbc.schema.tbl1;
# select * from dbc.schema.tbl2;
# select * from dbc.schema.tbl3;
```