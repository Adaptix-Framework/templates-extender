all: clean
	@ echo "    * Building _AGENT_ plugin"
	@ mkdir dist
	@ cp config.json ax_config.axs ./dist/
	@ GOEXPERIMENT=jsonv2,greenteagc go build -buildmode=plugin -ldflags="-s -w" -o ./dist/_AGENT_.so pl_main.go pl_agent.go
	@ echo "      done..."

#	@ echo "    * Building agent"
# 	PRE BUILDING OBJECTS
#	@ echo "      done..."

clean:
	@ rm -rf dist
