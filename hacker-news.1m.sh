#!/usr/bin/env bash

# <bitbar.title>Go Hacker News</bitbar.title>
# <bitbar.author>Martins Irbe</bitbar.author>
# <bitbar.author.github>MartinsIrbe</bitbar.author.github>
# <bitbar.image>http://imgur.com/9LjbmH9</bitbar.image>
# <bitbar.dependencies>Go</bitbar.dependencies>
# <bitbar.desc>Lists top stories from Hacker News API</bitbar.desc>

export GOPATH=$HOME/dev/go_workspace # change me!
export PATH=$HOME/bin:/usr/local/bin:$GOPATH/bin

ycombinator-bitbar