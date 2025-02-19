build:
	go build -0 dirscan

clean:
	rm -rf ./dirscan

commit:
	git add .
	@read -p "Enter commit message: " message; \
	git commit -m "$$message"
