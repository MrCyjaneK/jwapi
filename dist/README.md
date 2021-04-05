# Distribution packages

JW Study aim to be usable on every platform, no matter if it is your 512mb/8gb android 4.1 device, modern android 11 device, some Ubuntu Touch running via halium, postmarketOS (without usable calls), or your desktop device running Debian, or that crap m$ made.

On every device app shoud look exactly the same. If you want to port it to new OS, you need to create a simple wrapper, that will:

1. Run binary made for that os/arch combo
2. Open 127.0.0.1:8080 in a full screen mode.
3. Have a writable path, that will not require any specific permission. (for example /data/data/x.x.jwstudy on android or ~/.config on ubtouch)

Then you need to add a build rule to build.sh. Just take look at it, it's simple.

# Supported OS

 - ✅ Android (tested on 11 amd64)
 - ❎ Ubuntu Touch (tested on 16.04 Erfan's GSI v10) - works fine, but suddenly stopped.
 - ❓ Debian package
 - ❓ Flatpak
 - ❓ Sailfish
 - ❓ Windows
 - ❓ MacOS
 - ❓ iOS
<!--
 ✅ working
 ❎ works, but have issues
 ❌ not working
 -->

 