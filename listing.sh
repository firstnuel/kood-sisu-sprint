#!/bin/bash

ls -ltF | grep -vE '^..$' | tr '\n' ', ' | sed 's/.$//'

