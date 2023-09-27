curl https://learn.01founders.co/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name'
curl https://learn.01founders.co/assets/superhero/all.json | jq '.[] | select(.id == 170) | .powerstats.power'
curl https://learn.01founders.co/assets/superhero/all.json | jq '.[] | select(.id == 170) | .appearance.gender'