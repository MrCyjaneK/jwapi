#!/bin/bash
set -e

# This script should not be verbose.
# Simply telling what it is doing is enough.
GITVERSION="+git"$(date +%Y%m%d%H%M)"."$(git log -n 1 | tr " " "\n" | head -2 | tail -1 | head -c 7)
function ok {
    echo "OK"
}

root=$(dirname $0)
cd "$root"
root=$(pwd)
vcode="1.0.1-"$(cat ../VERSION_CODE | head -1)
echo "Building JW Study - version: $vcode";
cd ..
rm -rf build/
goprodbuilds=$(pwd)"/build/"
~/go/bin/packr2 clean
~/go/bin/packr2
goprod \
    -combo="linux/arm;linux/386;linux/arm64;linux/amd64;windows/amd64;windows/386" \
    -binname="jwstudy-browser" \
    -tags="guibrowser" \
    -version="$vcode"
goprod \
    -combo="linux/arm;linux/386;linux/arm64;linux/amd64;windows/amd64;windows/386" \
    -binname="jwstudy-lorca" \
    -tags="guilorca" \
    -version="$vcode"
goprod \
    -combo="linux/arm;linux/386;linux/arm64;linux/amd64;windows/amd64;windows/386" \
    -binname="jwstudy-nogui" \
    -tags="nogui" \
    -version="$vcode"
    
    
echo "/ ubtouch builds - daemon (custom location)"
echo -n -e "|- bin/jwstudy_ubtouch_arm64......"
goprod \
    -combo="linux/arm;linux/arm64;linux/amd64" \
    -binname="jwstudy-ubtouch" \
    -version="$vcode" \
    -ldflags="-X main.dataDir=/home/phablet/.local/share/jwstudy.anon -X main.Port=8080" \
    -package=false \
    -tags="nogui"

if [[ "X$SKIPANDROID" == "X" ]];
then
    echo "/ android builds - daemon (custom location) + NDK"
    # Android Version
    AV=21
    # NDK downloaded in android studio -> tools? -> sdk manager
    NDKV=$(ls ~/Android/Sdk/ndk/* -d | tr "/" "\n" | tail -1)
    NDK=~/Android/Sdk/ndk/$NDKV/toolchains/llvm/prebuilt/linux-x86_64/bin
    goprod \
        -combo="android/arm;android/386;android/arm64;android/amd64" \
        -binname="jwstudy" \
        -version="$vcode" \
        -ldflags="-X main.dataDir=/data/data/x.x.jwstudy/ -X main.Port=8080" \
        -ndk="$NDK" \
        -tags="nogui" \
        -package=false
fi
echo "===== Packaging"
echo "/ Packaging for Ubuntu Touch"
mkdir -p $goprodbuilds/click
for arch in arm64 arm amd64
do
    echo -n -e "|- bin/jwstudy_$arch.click............" | head -c 34
    cd $goprodbuilds
    cp ../dist/ubtouch ubtouch-$arch -r
    cd $goprodbuilds/ubtouch-$arch/
    clickable clean
    cp "$goprodbuilds/bin/jwstudy-ubtouch_linux_$arch" $(find . -name jwlib.bin)
    chmod +x $(find . -name jwlib.bin)
    sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' manifest.json.in
    archC=$arch
    if [[ "$arch" == "arm" ]];
    then
        archC="armhf"
    fi
    clickable build --arch=$archC
    cp build/*/app/*.click $goprodbuilds/click/jwstudy_$arch.click
    ok
done
echo "\_ DONE"
if [[ "X$SKIPANDROID" == "X" ]];
then
    echo "/ Packaging for android"
    mkdir -p $goprodbuilds/apk/
    for arch in arm64 arm amd64 386 all
    do
        echo -n -e "|- bin/jwstudy.android.$arch.apk.........." | head -c 34
        cd "$goprodbuilds"
        cp ../dist/android android-target-$arch -r
        cd android-target-$arch
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
        case $arch in
        "amd64")
            cp "$goprodbuilds/bin/jwstudy_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
            ;;
        "386")
            cp "$goprodbuilds/bin/jwstudy_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
            ;;
        "arm")
            cp "$goprodbuilds/bin/jwstudy_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
            ;;
        "arm64")
            cp "$goprodbuilds/bin/jwstudy_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
            ;;
        "all")
            cp "$goprodbuilds/bin/jwstudy_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
            cp "$goprodbuilds/bin/jwstudy_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
            cp "$goprodbuilds/bin/jwstudy_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
            cp "$goprodbuilds/bin/jwstudy_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
            ;;
        esac
        chmod +x $(find . -name jwlib.bin)
        sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' app/build.gradle
        ./gradlew build
        cp ./app/build/outputs/apk/debug/app-debug.apk "$goprodbuilds/apk/jwstudy.android.$arch.apk"
        ok
    done
    echo "\_ OK"
fi
echo "DONE! Everything is inside build/"

~/go/bin/packr2 clean