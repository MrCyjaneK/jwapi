install:
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/jwstudy
	cp dist/debian/logo.png /usr/share/icons/hicolor/scalable/apps/jwstudy.png
	cp dist/debian/jwstudy.desktop /usr/share/applications