# go_game_constractor
Function for constart game based on script for voice assistent. Example Marusia


# SCG
**SCG** - script generator. Generate script structs, functions for store texts of script in redis from yml, or json, or xml file.

### Example of use:
```(cmd)
scg --output=./scg/ --script=./scg/example/echo_game.yml
```
#### Args:
- `--output` - path to dir where need generate files
- `--script` - path to config file
- `--version` - show program version
- `--update` - save user changes in files

#### Note:
With the `--update` flag, user changes are saved unchanged. 
Comments are embedded in the code with the code that was generated based on the new initializing file.
These comments are limited to the lines ```// >>>>>>> Generated```.
The decision to apply the changes remains with you, as well as the decision to remove unnecessary functionality.

Example of comments:
```go
package example

func DoNothing() error {
	// >>>>>>> Generated 
	// 	return errors.New("Hello")
	// >>>>>>> Generated 
	return nil /// User changes
}
```