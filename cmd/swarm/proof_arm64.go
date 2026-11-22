//go:build linux && arm64

package main

// This file fails to build on linux/arm64 only, to prove build-check names that platform.
var proofArm64 = undefinedOnlyOnLinuxArm64
