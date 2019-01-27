all:
	go-bindata -o engine/assets.go -pkg engine engine/assets/
	go build -o raytracer
