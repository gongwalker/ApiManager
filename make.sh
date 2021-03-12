#!/bin/bash
# @Author: gongwen
# @Date:   2021-03-12 21:52:03
# @Last Modified by:   gongwen
# @Last Modified time: 2021-03-12 21:52:03

build() {
  os=$1
  echo "build ${os} being"
  CGO_ENABLED=0 GOOS=${os} GOARCH=amd64 go build -o ApiManager-${os} main.go &&
    mkdir -p releases/ApiManager-${os}-amd64/app &&
    mv ApiManager-${os} releases/ApiManager-${os}-amd64/ApiManager &&
    cp -rp db.sql config releases/ApiManager-${os}-amd64/ &&
    cp -rp app/static app/views releases/ApiManager-${os}-amd64/app

  if [ $os = linux ] || [ $os = mac ]; then
    cp -rp run.sh config releases/ApiManager-${os}-amd64/
  fi

  echo "build ${os} end"
}
case $1 in
build)
  if [[ $2 = "linux" ]]; then
    build linux
  elif [[ $2 = "windows" ]]; then
    build windows
  elif [[ $2 = "darwin" ]]; then
    build darwin
  else
      build linux
      build windows
      build darwin
  fi
  ;;
clear)
  rm -rf releases
  ;;
*)
  echo "$0 build [linux|windows|darwin|all] | clear"
  exit 4
  ;;
esac
