# JWapi
[![Build Status](https://ci.mrcyjanek.net/badge/build-jwapi.svg)](https://ci.mrcyjanek.net/jobs/build-jwapi)

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

# Installation

 - `jwstudy-nogui` - Daemon build, doesn't open gui, just serve it over http.
 - `jwstudy-lorca` - Opens electron-like app - recommended for desktop usage, on mobile have issues with highlighting.
 - `jwstudy-browser` - Opens gui in your default browser (recommended for mobile devices running Linux, due to fractional scaling issues with `-lorca` build)

Ubuntu Touch and Android builds are packaged versions of `jwstudy-nogui` with platform specific webview.
Android builds have issues on android lower than 10, please submit a merge request if you know how to properly fix this issue.

You can go directly to [my ci](https://ci.mrcyjanek.net/jobs/build-jwapi) do grab a binary for your system, or use platform-specific way of doing so:

Install my APT repo to your system. Make sure to run this command as root
```bash
# wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
    apt install ./mrcyjanek-repo-latest.deb && \
    rm ./mrcyjanek-repo-latest.deb && \
    apt update
```
After that install jwstudy. Please note that there are 3 versions of JW Study available. 
```bash
# apt install jwstudy-browser
```

# Building.

I need to make better docs for that, you can read dist/build.sh to learn how things are being compiled.

For the executable builds (excluting android) `go build` is enough, packaging is a bit more complicated, but in most cases few lines of bash solve that, check dist directory to learn how packaging is done.