.PHONY: all bench dghubble plar gammazero

all: bench

bench: gammazero dghubble goradix plar

gammazero:
	@echo
	go test -bench=Gammazero -run=xx

dghubble:
	@echo
	go test -bench=Dghubble -run=xx

goradix:
	@echo
	go test -bench=GoRadix -run=xx

plar:
	@echo
	go test -bench=Plar -run=xx
