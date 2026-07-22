# Migration INK

_Versioning with ink + inkjs is complicated._

## inkjs 2.4.0 — evaluated, NOT adopted (2026-07-21)

Latest is inkjs `v2.4.0` (2026-02) / inklecate `v1.2.1` (2026-05). We stay on inkjs `v2.3.0` /
inklecate `v1.2.0`. Rationale (validated, not assumed):

- **Behaviour is identical.** Same story format (`inkVersion:21`), same UMD `inkjs` global + `.Story`
  API. A full real-story playthrough produced byte-identical lines + tags, and cross-version
  `state.LoadJson` round-trips safely both directions.
- **The 2.4.0 "CVE fix" is build-time only.** It's a `rollup` (bundler) bump — it does not ship in the
  embedded runtime bundle we load via v8go, so it buys zero runtime security.
- **2.4.0 changes `state.toJson()` serialization** — emits `isInvisibleDefault` on choices and drops
  empty `tags:[]`. This breaks the golden test `Test_impl_Call_WhenStoryJSON` (inkcallerv8), which pins
  the exact state string. Adopting 2.4.0 means regenerating that fixture for no functional gain.
- inklecate `1.2.1` is Mac universal binaries + Chinese char ranges only — irrelevant to the Linux
  build container.

Revisit when a future inkjs release carries a real runtime fix.

## Ink 1.2.0

- https://github.com/inkle/ink/releases/tag/v.1.2.0
- inkjs engine `v2.3.0` (https://unpkg.com/inkjs@2.3.0/dist/ink.js)

## Ink 1.0, new version inkjs

- inkjs engine `v2.1.0` (https://unpkg.com/inkjs@2.1.0/dist/ink.js)
    - with this version, it would be possible to compile story.ink directly in js.

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
- note the web export in inklecate is using the wrong version of inkjs.

### Design choice on supporting 0.9

The problem is the schema of the ink JSON state changed.\
It's fine, this is internal to ink and they just have to be careful with the API.\
But at the same time, it was very nice getting everything out of the state.\
Version 1.0 is now using the JS API to get its data (but it is still possible to use the JSON state if needed).

In this `inkjs v2.0.0` version, `variablesState` is not showing the unchanged internal values. \
But can still be used to inject data from the outside.

By naming convention, `Ink v1.0` should be more stable than `Ink v0.9` (and then also inkjs). I decided I will not support the ink state struct for `Ink v0.9`.\
If it changes again in the future, I will think of supporting multiple version.
