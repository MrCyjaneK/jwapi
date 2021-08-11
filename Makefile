VERSION=2.0.0

install:
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/jwstudy
	cp dist/debian/logo.png /usr/share/icons/hicolor/scalable/apps/jwstudy.png
	cp dist/debian/jwstudy.desktop /usr/share/applications

android:
	goprod -combo="android/386;android/amd64;android/arm;android/arm64;android/all" -tags="nogui" -shouldpkg=true -binname="jwstudy" -version="${VERSION}" -appurl="http://127.0.0.1:4365/"  -ldflags="-X main.dataDir=/data/data/x.x.jwstudy/ -X git.mrcyjanek.net/mrcyjanek/jwapi/webui.SPort=4365"

linux-lorca:
	goprod -combo="linux/amd64;linux/arm;linux/arm64;linux/386" -binname="jwstudy-lorca" -tags="guilorca" -version="${VERSION}"

linux-browser:
	goprod -combo="linux/amd64;linux/arm;linux/arm64;linux/386" -binname="jwstudy-browser" -tags="guibrowser" -version="${VERSION}"

windows-lorca:
	goprod -combo="windows/amd64;windows/386" -binname="jwstudy-lorca" -tags="guilorca" -version="${VERSION}"

windows-browser:
	goprod -combo="windows/amd64;windows/386" -binname="jwstudy-browser" -tags="guibrowser" -version="${VERSION}"
