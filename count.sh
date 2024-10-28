#!/bin/bash

num_of_dirs=$(ls | wc -l )

for sub in *; do
    if [ -d "$sub" ]; then
       num_of_dirs=$((num_of_dirs + 1))
    fi
done

printf "Total files * 5: %d\n" $((num_of_dirs * 5))
