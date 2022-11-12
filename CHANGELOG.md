## 0.0.2-alpha

This is the initial release.

### Supported

* Generate redis store for scripts text
* Generate text manager for script, that can add to text vars
* Generate base structs for scenes

## 0.0.6-alpha

Some modification and addition of new functionality.

### Added

* Converts lib number to words and words to number
* Generate custom matchers
* Generate custom text error
* Support for error descriptions in scenes
* Support for matchers descriptions in scenes
* New positive number matcher with converting words to number
* CI for repository
* New coomand for scene. *StashScene* for stashing current scene
* The application can show its version using the ``--version`` flag

### Changes

* Now the goodbye scene described in the body of the script. Only its name is written in the script settings.
* The generated files are overwritten when the application is re-executed

## 0.1.0-alpha

### Added

* Auto generated of choose next scene
* Auto generated of button payloads
* Button support
* Saving user changes in script structures after its regeneration with the ``--update`` flag

### Changes

* Moved example to another dir
* Without the ``--update`` flag, all files will be regenerated with loss of changes

## 0.1.2-alpha

### Added

* Selectors between expected string in scene react
* Name for regex matched user input

### Changes

* Move creation of matchers to main part of config
* Bugs fix

## 0.1.3-alpha

### Added

* Check the message from the button separately from the matching

### Changes

* Correct name of number matcher in words
* Move creation of matchers to main part of config
* For an information scene, you need to specify only one following scene in the `nextScene` field
* For a non-informational scene, you need to specify the following scenes/scene only in the `nextScenes` field

## 0.1.4-alpha

### Added

* Documentation on github.io
* Create redis client package
* Help flag `-h --help` and flag `--http-server' fot generation base http-server

### Changes

* Generalize marusia handler
* New help text