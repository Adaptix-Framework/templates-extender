all: clean
	@ echo "    * Building _SERVICE_ plugin"
	@ mkdir -p dist
	@ cp config.yaml ax_config.axs ./dist/
	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_SERVICE_.so pl_main.go
	@ echo "      done..."

clean:
	@ rm -rf dist
