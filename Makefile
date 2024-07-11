.PHONY: yrly
yrly:
	go build -tags customcert -o build/yrly ./relayer
