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
* Help flag `-h --help` and flag `--http-server` fot generation base http-server
* Webhook implementation for gin and base http handler
* `Application` fields to marusia request
* `UserVKId` field to marusia request
* `CardItems` field to marusia request
* Own entity for runner
* Сase-independent comparison button name
* A check that the skill was launched based on the `new` field of the voice assistant request
* New fields for `SessionInfo`(formerly `UserInfo`): `UserVKID` and `IsNewSession`
* `Application` structs for scene `Request`
* Struct `UserInput`
* New field `Request` for `SceneRequest` instead of the [moved](#moved_fields) fields
* New field `UserVKID` for struct `UserInfo` in package `scene`
* New field `UserVKID` for struct `Request` in package `scene`
* New logger based on zap logger
* Http middleware for logger based on gin
* Context for request from marusia (сurrently it is used only for transmitting the logger, but it is possible to
  transmit something else)
* Generating server functional
* Generating app functional
* Generating config functional
* Generating logger functional
* Generating handler functional
* Generating entry point of server
* Loading resources for convert package
* `ARM` and `i386` build architecture for `Windows`, `MacOs`, `Linux`

### Changes

* Generalized marusia handler
* New help text
* Adjusted the default setting of the next scene in the React function
* Moved `UserId` from the session field of the request to the independent `User` structure in the `Session` field of the
  request
* Renamed package `hub` to `runner`
* ScriptDirector moved to package `scriptdirector`
* Renamed `UserInfo` to `SessionInfo`
* <a id="moved_fields"></a>The fields `Command`, `FullUserText`, `WasButton`, `Payload` have been moved
  from `SceneRequest` to struct `UserInput`
* The structs `Text` and `Button` have been moved from package `scene` to package `director`
* Used new Logger into webhook
* Reformatted Directory structure(see in [documentation](https://thcompiler.github.io/go_game_constractor/manual/))
* Some bug fixed in other generation

### Remove

* A check that the skill was launched based on receiving the `debug` message from the user

## 0.2.1-alpha

### Added

* Checking the name of a custom matcher for the presence of the same name in standard matchers
* The ability to specify the next scene for each matcher and appropriate checks. Specified in the `toScene' field
* The ability to specify the next scene for each button and appropriate checks. Specified in the `toScene' field
* Dequeue and Queue structures to package `pkg/structures`
* Graph functional to package `pkg/graph`
* Storing to context value from user input
* Loading data from context as a function
* Getting values for the text of scenes from the context

### Changes

* Removed the remaining parts from the old way of describing the matchers
* Updated some errors
