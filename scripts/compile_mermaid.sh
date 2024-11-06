#!/bin/bash

for file in $(find docs/ -name "*.mmd"); do
    output=$(head -n 1 "$file" | awk '{print $2}')
    echo "Compiling $file as $output"
    mmdc -i $file -o $output
done
