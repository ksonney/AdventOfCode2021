#!/bin/bash

INFILE="../input.txt"
RESULT=0
INCREASE=0
COUNTER=0
PREVVAL=0

for CHECKVAL in $(cat ${INFILE}); do
    if [ $COUNTER -gt 0 ]; then
        if [ $CHECKVAL -gt $PREVVAL ]; then
            INCREASE=$(( INCREASE + 1 ))
        fi
    fi
    PREVVAL=$CHECKVAL
    COUNTER=$(( COUNTER + 1 ))
done
echo Processed ${COUNTER} entries
echo Found ${INCREASE} increases