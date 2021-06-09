dev:
	go run httpd/main.go

test:
	cd httpd/test && go test ./...