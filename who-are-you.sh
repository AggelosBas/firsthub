#!/bin/bash
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
