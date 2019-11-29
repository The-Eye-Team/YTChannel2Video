#!/bin/bash

youtube-dl --force-ipv4 -j --flat-playlist $1 | jq '.id' |  cut -d '"' -f 2 | ts 'https://youtube.com/watch?v=' | tr -d ' '