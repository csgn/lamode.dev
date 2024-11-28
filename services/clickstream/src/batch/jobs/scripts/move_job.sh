#!/bin/bash
set -e

function move_raw_events() {
    source_dir="${1}"
    dest_dir="${2}"

    isEmpty=$(hdfs dfs -ls ${source_dir} | grep '\.json$' | wc -l)
    if [[ isEmpty -eq 0 ]];then
        echo "There is no elements to move. Skipping..."
    else
        hdfs dfs -mv ${source_dir}/*.json ${dest_dir}
    fi
}

move_raw_events $1 $2
