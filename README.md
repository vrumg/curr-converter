# curr-converter
Currency exchange console utility.

# arguments
<amount:float> <src_symbol:string> <dst_symbol:string>

# build
go build -o ./bin/fiatconv main.go

# run
Example: ./bin/fiatconv 123.45 USD RUB
Sample output: 9300.105176 USD RUB