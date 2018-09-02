# TODO: separate

.PHONY: all

all: clean prepare build-public build-server

clean:
	rm -rf out

prepare:
	mkdir out && \
	(cd public && npm install) && \
	(cd server && dep ensure -vendor-only)

build-public:
	cd public && \
	ng build --prod && \
	mv dist/public ../out/dist && \
	rm -r dist

build-server:
	cd server && \
	go build -o ../out/fws-server

