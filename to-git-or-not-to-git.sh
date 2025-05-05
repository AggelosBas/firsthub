curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select( .id == 170 ) | .name'
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select( .id == 170 ) | .powerstats.power'
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select( .id == 170 ) | .appearance.gender'