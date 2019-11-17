# week_parser
`月・火`などの文字列をパースし、曜日として扱える形にするモジュール

## Usage
### Install
```sh
git clone https://github.com/nanachoo0882/week_parser.git
go get golang.org/x/tools/cmd/goyacc
```

### Run
```sh
goyacc -o parser/parser.go parser/parser.go.y
go run parse.go '月曜日~水曜'
```
