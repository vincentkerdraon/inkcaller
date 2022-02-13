# Migration INK

_Versioning with ink + inkjs is complicated._

## Ink 0.9 -> Ink 1.0

First version of this lib was using:
- in story.json: `"inkVersion":19`
- inkjs from https://github.com/y-lohse/inkjs/tree/master/templates (but in fact this is outdated)

New version is for:
- `Ink v1.0` (https://github.com/inkle/ink/releases/tag/v1.0.0)
- `inklecate2`? in the release of `Ink v1.0`
- in input story.json: `"inkVersion":20`
- in output state json: `"inkSaveVersion": 9`
- in output state json: `"inkFormatVersion": 20`
- with inkjs engine `v2.0.0` (https://unpkg.com/inkjs@2.0.0/dist/ink-es2015.js)
- note the web export in inky editor is using the wrong version of inkjs.

### Design choice on supporting 0.9

The problem is the schema of the ink state changed.\
It's fine, this is internal to ink and they just have to be careful with the API.\
But at the same time, it was very nice getting everything out of the state.

In this 1.0 version, `variablesState` is not showing the unchanged internal values. 
But can still be used to inject data from the outside.

By naming convention, `Ink v1.0` should be more stable than `Ink v0.9`. I decided I will not support the ink state struct for `Ink v0.9`.
If it changes again in the future, I will think of supporting multiple version.
