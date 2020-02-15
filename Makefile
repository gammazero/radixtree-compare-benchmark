.PHONY: all bench dghubble plar gammazero

all: bench

bench: dghubble gammazero plar

dghubble:
	@echo
	go test -v -bench=Dghubble -run=xx

gammazero:
	@echo
	go test -v -bench=Gammazero -run=xx

plar:
	@echo
	go test -v -bench=Plar -run=xx
