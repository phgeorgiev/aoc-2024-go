#!/bin/sh

dayFolder=$(printf "Day%02d" "$1")

if [ -d src/"$dayFolder" ]; then
    echo "Solution for day $1 already exists."
    exit
fi

mkdir src/"$dayFolder"
touch src/"$dayFolder"/input.txt
touch src/"$dayFolder"/test_input.txt
cp template/main.go src/"$dayFolder"/main.go
cp template/main_test.go src/"$dayFolder"/main_test.go
