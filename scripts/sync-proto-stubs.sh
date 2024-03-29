#!/usr/bin/env bash
REPO="https://${GH_TOKEN}@github.com/lukasjarosch/educonn-go-api.git"
REPO_PATH="$(pwd)/educonn-go-api"
COMMIT=$(git rev-parse --short HEAD)
COMMIT_MSG=$(git log -1 --pretty=%B)
GEN_PATH="$(pwd)/api/gen/go"

function clone() {
	rm -rf "$REPO_PATH"
	echo "Cloning into repo: $REPO"
	git clone $REPO "$REPO_PATH"
}

function copyStubs() {
	cp -R $GEN_PATH/* $REPO_PATH/
}

function setupGit() {
	git config --global user.email "travis@travis-ci.org"
    git config --global user.name "Travis CI"
}

function commitAndPush() {
	enterDir $1

	setupGit

	git add -N .

	if ! git diff --no-ext-diff --quiet --exit-code; then
		git add .
		git status
		git commit -m "[lukasjarosch/educonn-monorepo@$COMMIT]: $COMMIT_MSG"
		git push
	else
		echo "No changes detected for $1"
	fi

	leaveDir
}

# Helper for adding a directory to the stack and echoing the result
function enterDir {
  pushd $1 > /dev/null
}

# Helper for popping a directory off the stack and echoing the result
function leaveDir {
  popd > /dev/null
}

clone
copyStubs
commitAndPush $REPO_PATH