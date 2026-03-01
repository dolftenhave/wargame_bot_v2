all: 
	export CGO_ENABLED=1
	go run . &

debug:
	export CGO_ENABLED=1
	go run . &
	tail -f logs.txt
