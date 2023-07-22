#!/bin/sh

./minifab cleanup

sudo rm -rf vars/
cd api-go/
rm -rf wallet/
cd ..
rm -rf wallet/
