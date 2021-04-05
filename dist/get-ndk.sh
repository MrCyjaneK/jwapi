#!/bin/bash
mkdir -p ~/Android/Sdk/ndk/
latest=$(wget --quiet https://developer.android.com/ndk/downloads/ -O - | tr '>' ">\n" | grep "linux-x86_64.zip" | grep href | tr '"' "\n" | head -2 | tail -1)
if [[ "X$(cat ~/Android/Sdk/version)" == "X$latest" ]];
then
    echo "No need to update ndk"
    exit 0
fi
set -e
wget "$latest" -O ndk.zip
rm -rf android-ndk-*
unzip ndk.zip &>/dev/null
rm ndk.zip
echo "$latest" > ~/Android/Sdk/version