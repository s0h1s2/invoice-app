#!/bin/bash
while ! nc -z db 3306; do sleep 3; done
./app
