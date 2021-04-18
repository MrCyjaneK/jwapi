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

# Installation

You can go directly to [my ci](https://ci.mrcyjanek.net/jobs/build-jwapi) do grab a binary for your system, or use platform-specific way of doing so:

  - debian
    
    Install my APT repo to your system. Make sure to run this command as root
    ```bash
    # wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
      apt install ./mrcyjanek-repo-latest.deb && \
      rm ./mrcyjanek-repo-latest.deb && \
      apt update
    ```
    After that install jwstudy
    ```bash
    # apt install jwstudy
    ```

# Building.

I need to make better docs for that, you can read dist/build.sh to learn how things are being compiled.

For the executable builds (excluting android) `go build` is enough, packaging is a bit more complicated, but in most cases few lines of bash solve that, check dist directory to learn how packaging is done.