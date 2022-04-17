#!/bin/sh
alias fig="figlet -f small"
test_folder="./tmp_test"

setUp() {
  mkdir $test_folder
  cd $test_folder
  touch old.json
  mkdir -p ./data/json
  touch ./data/json/old.json
  cd ..
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
  ./main -path tmp_test
  showStatus "After"
  rm -rf $test_folder
}

runTest "Testing duplicates"
