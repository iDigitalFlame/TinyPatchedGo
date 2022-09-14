# TinyPatchedGo

Patched Golang source to remove some symbol tables for speed and smaller binary
sizes (~0.9MB saved!)

Replaces and guts the "fmt", "runtime" and "unicode" packages.

Can be used by [JetStream in ThunderStorm](https://github.com/iDigitalFlame/ThunderStorm).

__For now...__

Patches are by me, original code COPYRIGHT/CREDIT is to the Golang authors.
