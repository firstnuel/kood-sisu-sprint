#!/bin/bash

ls -ltF | tr '\n' ', ' | sed 's/.$//'

