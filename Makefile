.PHONY: all
all:
	go build github.com/proxypoke/gotils/cmd/...

.PHONY: doc
doc:
	a2x -f xhtml README.asciidoc

.PHONY: clean
clean:
	rm -rf *.css
	rm -rf *.html
