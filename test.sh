#!/bin/sh
alias fig="figlet -f small"
test_folder="./tmp_test"

setUp() {
  mkdir $test_folder
  touch $test_folder/demo.json
}

showStatus() {
  fig $1
  exa -T --icons tmp_test
}

runTest(){
  fig $1
  go build *.go
  setUp
  showStatus "Before"
  fig "Running Code"
  mkdir $test_folder
  ./main -path $test_folder
  showStatus "After"
  rm -rf $test_folder
}

runTest "Basic Test"
