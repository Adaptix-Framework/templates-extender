all: clean
	@ echo "    * Building _AGENT_ plugin"
	@ mkdir -p dist
	@ cp config.yaml ax_config.axs ./dist/
	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_AGENT_.so pl_main.go
	@ echo "      done..."

# Optional: build agent payload objects / binaries
#	@ echo "    * Building agent payload"
#	@ (cd src_agent && $(MAKE) --no-print-directory)
#	@ echo "      done..."

clean:
	@ rm -rf dist
