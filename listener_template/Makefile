all: clean
	@ echo "    * Building _LISTENER_ plugin"
	@ mkdir -p dist
	@ cp config.yaml ax_config.axs ./dist/
	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_LISTENER_.so pl_main.go
	@ echo "      done..."

# If you split transport into pl_transport.go:
#	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_LISTENER_.so pl_main.go pl_transport.go

clean:
	@ rm -rf dist
