#!/bin/bash

cmp --silent $1 $2 && echo 'Files are the same' || echo 'Files are different'