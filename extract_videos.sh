#!/bin/bash

youtube-dl -j --flat-playlist $1 | jq '.id' |  cut -d '"' -f 2 | ts 'https://youtube.com/watch?v=' | tr -d ' '