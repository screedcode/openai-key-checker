#!/bin/bash
for key in $(cat $1) 
do 
  if curl --output /dev/null --silent --fail https://api.openai.com/v1/models  -H "Authorization: Bearer $key"; then
    echo "$key"
  fi 
done
