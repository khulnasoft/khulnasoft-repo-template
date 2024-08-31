# Variables
PROTOC_VERSION := 3.15.8
PROTOC_ZIP := protoc-$(PROTOC_VERSION)-linux-x86_64.zip
PROTOC_URL := https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP)

# Install protoc
install_protoc:
	@echo "Installing protoc $(PROTOC_VERSION)..."
	curl -OL $(PROTOC_URL)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o -j $(PROTOC_ZIP) 'include/*' -d /usr/local/include/
	rm -f $(PROTOC_ZIP)
	@echo "protoc $(PROTOC_VERSION) installed successfully."

# Run golangci-lint
check:
	@echo "Running golangci-lint..."
	golangci-lint run

# Clean golangci-lint cache
check-clean-cache:
	@echo "Cleaning golangci-lint cache..."
	golangci-lint cache clean

# Download proto file
protoc-setup:
	@echo "Setting up proto files..."
	wget -P meshes https://raw.githubusercontent.com/meshplay/meshplay/master/server/meshes/meshops.proto

# Generate Go code from proto files
proto:
	@echo "Generating Go code from proto files..."
	/usr/local/bin/protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

# Serve Jekyll site with drafts and livereload
site:
	@echo "Serving Jekyll site with drafts and livereload..."
	$(jekyll) serve --drafts --livereload

# Build Jekyll site with drafts
build:
	@echo "Building Jekyll site with drafts..."
	$(jekyll) build --drafts

# Run Jekyll site in Docker
docker:
	@echo "Running Jekyll site in Docker..."
	docker run --name site -d --rm -p 4000:4000 -v `pwd`:"/srv/jekyll" jekyll/jekyll:4.0.0 \
	bash -c "bundle install; jekyll serve --drafts --livereload"
