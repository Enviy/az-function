#!/bin/bash

TEST="curl -d @sample-post.json -H \"Content-Type: application/json\"  -X POST http://localhost:7071/api/sentinel"
echo $TEST

RESPONSE=`$TEST`
echo $RESPONSE
