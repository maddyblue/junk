#!/bin/bash

cp main.go main/main.go
cd main
rm -rf _third_party
party -c -v
