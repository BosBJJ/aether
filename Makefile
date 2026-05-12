build:
	go build -o aether .
	sudo setcap cap_net_raw+ep ./aether
	
test:
	go test ./internal/crypto -run .
	go test ./internal/scanner -run TestGrabBanner -v

clean: 
	rm -f aether