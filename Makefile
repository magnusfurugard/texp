# Convenience file for building and "installing".

build: 
	CGOENABLED=false go build -o texp main.go

# Make sure ~/bin/ is in your $PATH.
install: build
	chmod +x texp && mv texp ~/bin/texp