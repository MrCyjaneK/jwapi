# JWapi
[![Build Status](https://ci.mrcyjanek.net/badge/d75ff2b4?branch=master)](https://ci.mrcyjanek.net/repos/169)

JW api is a main project for:

 - libjw - Golang library to access JW.org apis
 - JW Study - OpenSource JW Library app - fully open source and cross platform JW Library alternative.

# Why does this project exist?

Isn't the JW Library enough?

Well, while on some standard devices it _works_, there are several issues:

 - JW Library app register handlers for every possible thing. Seriously. I've stopped using when after flashing pure Lineage OS, and trying to call somebody it just opened JW Library.
 - That app is not lightweight, it require a lot of network usage, storage, and won't run on low-end devices.
 - You can't study on your desktop, and just grab all your notes and highlights with you on the go.
 - It is android-only app. How about those who run Ubuntu Touch, postmarketOS, mobian, droidian, SailfishOS, arch linux or on a halium recovery with telnet only access?
 - It is a proprietary software.

So is jwapi better?

Well. You can make it better, simply submit a PR.


<!--
 ✅ Done
 ❎ Done, but have issues
 ❌ Todo
 -->
# Features (UX)
|    |       |
| -- | ----- |
| ✅ | Downloading publications, to use them offline is made simpler - you just click on the publication. |
| ✅ | No unnecessary permissions required - no need to read files on your device, app uses OS-specific path that doesn't require any permissions |
| ✅ | Study on computer, phone, and have the same highlights on both devices |
| ✅ | Highlighting content |
| ❌ | Video/Audio files support |
| ✅ | Notes |
| ❌ | Searching in noted |
| ❌ | Backups (for now just copy `~/LibJWgo/db/`) |


# Features (Techy)

|    |       |
| -- | ----- |
| ✅ | Same app on every platform - HTML, CSS and JS for the frontend and Golang for the backend |
| ✅ | No special permissions required, app doesn't register any handlers (android - you can open any file with JW Library...) | 
| ✅ | Use same API as JW Library |
| ✅ | Give you better control of what is going on in the background. |
| ❎ | Package `jwstudy` for all systems (check dist/ directory) |
| ❌ | Convert JWPUB into Epub | 

# Binary versions

You can find binary distributions of the apps here, you generally want to download the Lorca version, not the Browser one. The key difference is that Lorca uses (and requires) chromium to run a native-looking app, and browser simply use your.

| Lorca | Browser |
| ----- | ------- |
| [Windows 64bit](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_windows_amd64.exe) | [Windows 64bit](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_windows_amd64.exe) |
| [Windows 32bit](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_windows_386.exe) | [Windows 32bit](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_windows_386.exe) |
| [Linux amd64](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_linux_amd64) | [Linux amd64](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_linux_amd64) |
| [Linux 386](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_linux_386) | [Linux 386](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_linux_386) |
| [Linux arm](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_linux_arm) | [Linux arm](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_linux_arm) |
| [Linux arm64](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy_linux_arm64) | [Linux arm64](https://static.mrcyjanek.net/abstruse/jwapi/jwstudy-browser_linux_arm64) |
# Linux installation

 - `jwstudy` - Opens electron-like app - recommended for desktop usage, on mobile devices (running Linux) have issues with highlighting.
 - `jwstudy-browser` - Opens gui in your default browser (recommended for mobile devices running Linux, due to fractional scaling issues with `-lorca` build)

**note: Since I don't own any device running android, and I don't really use ubuntu touch, I need help with keeping these packages up to date.**
Ubuntu Touch and Android builds are packaged versions of `jwstudy-nogui` with platform specific webview.


Install my APT repo to your system. Make sure to run this command as root
```bash
# wget 'https://static.mrcyjanek.net/abstruse/apt-repository/mrcyjanek-repo/mrcyjanek-repo_2.0-1_all.deb' -O cyjanrepo.deb && \
    apt install ./cyjanrepo.deb && \
    rm ./cyjanrepo.deb && \
    apt update && \
    apt install jwstudy -y
```
After that you will get latest jwstudy, and you will get updated together with other apps on your system.

# Building.

I need to make better docs for that, you can read dist/build.sh to learn how things are being compiled.

For the executable builds (excluting android) `go build` is enough, packaging is a bit more complicated, but in most cases few lines of bash solve that, check dist directory to learn how packaging is done.