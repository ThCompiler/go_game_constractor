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
* The application can show its version using the flag ``--version``

### Changes

* Now the goodbye scene described in the body of the script. Only its name is written in the script settings.
* The generated files are overwritten when the application is re-executed

## 0.1.0-alpha

### Added

* Auto generated of choose next scene
* Auto generated of button payloads
* Button support

### Changes

* Moved example to another dir