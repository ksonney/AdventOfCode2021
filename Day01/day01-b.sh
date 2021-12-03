#!/bin/bash

INFILE="../input.txt"
RESULT=0
INCREASE=0
COUNT=0
CURPOS=0
declare -a ALLDATA
declare -a CURARRAY
declare -a PREVARRAY
CURVAL=0
PREVAL=0

for CHECKVAL in $(cat ${INFILE}); do
    ALLDATA[$COUNT]=$CHECKVAL
    COUNT=$(( COUNT + 1 ))
    echo Count is $COUNT
done

while [ $CURPOS -le $COUNT ]; do
    if [ $CURPOS -le 2 ]; then
        VALA=${ALLDATA[$CURPOS]}
        CURVAL=$(( CURVAL + VALA ))
    else
        PREVVAL=$CURVAL
        POSMIN1=$(( CURPOS - 1 ))
        POSMIN2=$(( CURPOS - 2))
        VALA=${ALLDATA[$CURPOS]}
        VALB=${ALLDATA[$POSMIN1]}
        VALC=${ALLDATA[$POSMIN2]}
        CURVAL=$(( VALA + VALB + VALC ))
        echo Current Val is $CURVAL, Previous Val is $PREVVAL, Position is $CURPOS
        if [ $CURVAL -gt $PREVVAL ]; then
            INCREASE=$(( INCREASE + 1 ))
        fi
    fi
    CURPOS=$(( CURPOS + 1 ))
done

echo Processed ${COUNT} entries
echo Found ${INCREASE} increases