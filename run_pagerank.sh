#!/bin/bash 

INPUT1=${1}
DIR=result
INPUT2=$DIR/input2.txt
OUTPUT1=$DIR/preproc.txt
OUTPUT2=$DIR/output.txt
FINAL=$DIR/ranks.txt

rm -rf $DIR; mkdir $DIR

cat $INPUT1 |sort -t$'\t' -k 1,1 |./PrepReducer > $OUTPUT1
TOTAL=$(cat $OUTPUT1|wc -l | sed 's/^ *//')
echo "There are $TOTAL points"

cp $OUTPUT1 $OUTPUT2
counter=0
active=1
while [[ $active -gt 0 ]]; do
    counter=$((counter+1))
    echo "Round $counter"
    cp $OUTPUT2 $INPUT2
    rm -f $OUTPUT2
    cat $INPUT2|./RankMapper|sort -t$'\t' -k 1,1|./RankReducer -total=$TOTAL -thresh=0.001 > $OUTPUT2
    active=$(cat $OUTPUT2| awk -F '\t' '{print $2}' |grep "active"|wc -l )
    echo "There are $active active points"
done

cat $OUTPUT2| awk -F '\t' '{printf "%s\t%s\n",$1,$3}'|sort -t$'\t' -k2,2nr > $FINAL