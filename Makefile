.PHONY: all bench dghubble plar gammazero

all: bench

bench: dghubble gammazero plar

dghubble:
	@echo
	go test -bench=Dghubble -run=xx

gammazero:
	@echo
	go test -bench=Gammazero -run=xx

plar:
	@echo
	go test -bench=Plar -run=xx
