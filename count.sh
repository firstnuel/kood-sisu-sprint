#!/bin/bash

num_of_items=$(find . -type f -o -type d | wc -l)

result=$((num_of_items * 5))

printf "\t\vTotal files * 5: %d\v\n" "$result"