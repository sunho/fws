.PHONY: all

all: clean prepare dep build

clean:
	rm -rf out

prepare:
	mkdir out

dep:
	(cd public && npm install) && \
	(cd server && dep ensure -vendor-only)

build: build-public build-server

build-public:
	cd public && \
	npm run build && \
	rm -rf ../out/dist && \
	mv dist/public ../out/dist && \
	rm -r dist

build-server:
	cd out && \
	go build ../server

dev: dev-public build-server

dev-public:
	cd public && \
	ng build && \
	rm -rf ../out/dist && \
	mv dist/public ../out/dist && \
	rm -r dist
