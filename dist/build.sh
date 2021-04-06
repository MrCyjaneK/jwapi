#!/bin/bash
set -e

# This script should not be verbose.
# Simply telling what it is doing is enough.

function ok {
    echo "OK"
}

root=$(dirname $0)
cd "$root"
vcode=$(cat ../VERSION_CODE | head -1)
echo "Building JW Study - version: $vcode";
rm -rf out || true
mkdir out
cd out
builddir=$(pwd)

cd ../..
~/go/bin/packr2 clean
~/go/bin/packr2

cd "$builddir"
mkdir bin
echo "/ Linux builds - daemon."
echo -n -e "|- bin/jwstudy_linux_386.........."
GOOS=linux GOARCH=386     go build -o bin/jwstudy_linux_386 ../../ && ok
echo -n -e "|- bin/jwstudy_linux_amd64........"
GOOS=linux GOARCH=amd64   go build -o bin/jwstudy_linux_amd64 ../../ && ok
echo -n -e "|- bin/jwstudy_linux_arm.........."
GOOS=linux GOARCH=arm     go build -o bin/jwstudy_linux_arm ../../ && ok
echo -n -e "\_ bin/jwstudy_linux_arm64........"
GOOS=linux GOARCH=arm64   go build -o bin/jwstudy_linux_arm64 ../../ && ok
echo "/ Windows builds - daemon"
echo -n -e "|- bin/jwstudy_windows_386.exe...."
GOOS=windows GOARCH=386   go build -o bin/jwstudy_windows_386.exe ../../ && ok
echo -n -e "|- bin/jwstudy_windows_amd64.exe.."
GOOS=windows GOARCH=amd64 go build -o bin/jwstudy_windows_amd64.exe ../../ && ok
echo -n -e "\_ bin/jwstudy_windows_arm.exe...."
GOOS=windows GOARCH=arm go build -o bin/jwstudy_windows_arm.exe ../../ && ok
echo "/ ubtouch builds - daemon (custom location)"
echo -n -e "|- bin/jwstudy_ubtouch_arm64......"
GOOS=linux GOARCH=arm64   go build --ldflags "-X main.dataDir=/home/phablet/.local/share/jwstudy.anon" -o bin/jwstudy_ubtouch_arm64 ../../ && ok
echo -n -e "|- bin/jwstudy_ubtouch_arm........"
GOOS=linux GOARCH=arm     go build --ldflags "-X main.dataDir=/home/phablet/.local/share/jwstudy.anon" -o bin/jwstudy_ubtouch_arm ../../ && ok
echo -n -e "\_ bin/jwstudy_ubtouch_amd64......"
GOOS=linux GOARCH=amd64   go build --ldflags "-X main.dataDir=/home/phablet/.local/share/jwstudy.anon" -o bin/jwstudy_ubtouch_amd64 ../../ && ok
echo "/ android builds - daemon (custom location) + NDK"
# Android Version
AV=21
# NDK downloaded in android studio -> tools? -> sdk manager
NDKV=$(ls ~/Android/Sdk/ndk/* -d | tr "/" "\n" | tail -1)
export ANDROID_HOME=~/Android/Sdk
export ANDROID_NDK_ROOT=~/Android/Sdk/ndk/$NDKV
NDK=~/Android/Sdk/ndk/$NDKV/toolchains/llvm/prebuilt/linux-x86_64/bin
echo -n -e "|- bin/jwstudy_android_arm64......"
CGO_ENABLED=1 CC=$NDK/aarch64-linux-android$AV-clang CXX=$NDK/aarch64-linux-android$AV-clang GOOS=android GOARCH=arm64     go build --ldflags "-X main.dataDir=/data/data/x.x.jwstudy/" -o bin/jwstudy_android_arm64 ../../ && ok
echo -n -e "|- bin/jwstudy_android_arm........"
CGO_ENABLED=1 CC=$NDK/armv7a-linux-androideabi$AV-clang CXX=$NDK/armv7a-linux-androideabi$AV-clang GOOS=android GOARCH=arm go build --ldflags "-X main.dataDir=/data/data/x.x.jwstudy/" -o bin/jwstudy_android_arm ../../ && ok
echo -n -e "|- bin/jwstudy_android_amd64......"
CGO_ENABLED=1 CC=$NDK/x86_64-linux-android$AV-clang CXX=$NDK/x86_64-linux-android$AV-clang GOOS=android GOARCH=amd64       go build --ldflags "-X main.dataDir=/data/data/x.x.jwstudy/" -o bin/jwstudy_android_amd64 ../../ && ok
echo -n -e "\_ bin/jwstudy_android_386........"
CGO_ENABLED=1 CC=$NDK/i686-linux-android$AV-clang CXX=$NDK/i686-linux-android$AV-clang GOOS=android GOARCH=386             go build --ldflags "-X main.dataDir=/data/data/x.x.jwstudy/" -o bin/jwstudy_android_386 ../../ && ok

echo "===== Packaging"
echo "/ Packaging for Ubuntu Touch"

for arch in arm64 arm amd64
do
    echo -n -e "|- bin/jwstudy_$arch.click............" | head -c 34
    cd $builddir/
    cp ../ubtouch ubtouch-$arch -r
    cd $builddir/ubtouch-$arch/
    clickable clean
    cp "$builddir/bin/jwstudy_ubtouch_$arch" $(find . -name jwlib.bin)
    chmod +x $(find . -name jwlib.bin)
    sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' manifest.json.in
    archC=$arch
    if [[ "$arch" == "arm" ]];
    then
        archC="armhf"
    fi
    clickable build --arch=$archC
    cp build/*/app/*.click $builddir/bin/jwstudy_$arch.click
    ok
done
echo "\_ DONE"
echo "/ Packaging for android"
for arch in arm64 arm amd64 386 all
do
    echo -n -e "|- bin/jwstudy.android.$arch.apk.........." | head -c 34
    cd "$builddir"
    cp ../android android-target-$arch -r
    cd android-target-$arch
    touch "$builddir/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
    touch "$builddir/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
    touch "$builddir/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
    touch "$builddir/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
    case $arch in
    "amd64")
        cp "$builddir/bin/jwstudy_android_amd64" "$builddir/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
        ;;
    "386")
        cp "$builddir/bin/jwstudy_android_386"   "$builddir/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
        ;;
    "arm")
        cp "$builddir/bin/jwstudy_android_arm"   "$builddir/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
        ;;
    "arm64")
        cp "$builddir/bin/jwstudy_android_arm64" "$builddir/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
        ;;
    "all")
        cp "$builddir/bin/jwstudy_android_amd64" "$builddir/android-target-$arch/app/src/main/resources/lib/x86_64/jwlib.bin"
        cp "$builddir/bin/jwstudy_android_386"   "$builddir/android-target-$arch/app/src/main/resources/lib/x86/jwlib.bin"
        cp "$builddir/bin/jwstudy_android_arm"   "$builddir/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/jwlib.bin"
        cp "$builddir/bin/jwstudy_android_arm64" "$builddir/android-target-$arch/app/src/main/resources/lib/arm64-v8a/jwlib.bin"
        ;;
    esac
    chmod +x $(find . -name jwlib.bin)
    sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' app/build.gradle
    ./gradlew build
    cp ./app/build/outputs/apk/debug/app-debug.apk "$builddir/bin/jwstudy.android.$arch.apk"
    ok
done
echo "\_ OK"
echo "DONE! Everything is inside $builddir/bin"