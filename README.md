# go_game_constractor
Function for constart game based on script for voice assistent. Example Marusia


# SCG
**SCG** - script generator. Generate script structs, functions for store texts of script in redis from yml, or json, or xml file.

Example of use:
```(cmd)
scg --output=./scg/ --script=./scg/example/echo_game.yml
```
Args:
- *output* - path to dir where need generate files
- *script* - path to config file