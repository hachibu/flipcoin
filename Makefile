# Clean

.PHONY: clean
clean: bazel-clean docker-clean git-clean

.PHONY: git-clean
git-clean:
	@git clean -dfx
	@git gc --aggressive --prune=now

# Gazelle

.PHONY: bazel-clean
bazel-clean:
	@rm -rf bazel-*
	@bazel clean --expunge

.PHONY: gazelle
gazelle:
	@go mod tidy
	@bazel run //:gazelle-update-repos
	@bazel run //:gazelle

# Bazel

.PHONY: build
build:
	@bazel build //...

.PHONY: start
start: build
	@bazel run //cmd/web:web 

.PHONY: test
test:
	@bazel test //...

# Docker

.PHONY: docker-clean
docker-clean:
	@docker-clean

.PHONY: docker-build
docker-build:
	@bazel build //cmd/web:image
	@bazel run //cmd/web:image

.PHONY: docker-start
docker-start:
	@docker run --rm -it --platform linux/amd64 -p8080:8080 bazel/cmd/web:image

.PHONY: docker-push
docker-push:
	@docker login
	@bazel run //cmd/web:image_push

.PHONY: docker-push-github
docker-push-github:
	@bazel run //cmd/web:image_push_github

# Fly

.PHONY: fly-login
fly-login:
	@flyctl auth login

.PHONY: fly-launch
fly-launch:
	@flyctl launch --no-deploy --generate-name --region sjc --image hachibu/flipcoin:latest

.PHONY: fly-deploy
fly-deploy:
	@flyctl deploy
