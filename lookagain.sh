#!/bin/bash
find . -type f -name "*.sh"| sort -r |sed 's/\.sh$//'|xargs -n 1 basename