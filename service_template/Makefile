all: clean
	@ echo "    * Building _SERVICE_ plugin"
	@ mkdir dist
	@ cp config.yaml ./dist/
	@ go mod tidy
	@ CGO_ENABLED=1 go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_SERVICE_.so pl_main.go
	@ echo "      done..."

clean:
	@ rm -rf dist
