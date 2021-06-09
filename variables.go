package main

var minMatchLength = 6
var processSameFile = false
var version = "0.5.0"
var dirFilePaths = []string{}
var allowListExtensions = []string{}
var locationExcludePattern = []string{}
var ignoreIgnoreFile = false
var ignoreGitIgnore = false
var minifiedLineByteLength = 255
var maxReadSizeBytes int64 = 10000000
var verbose = false

const (
	DUPLICATE_DISABLE = "duplicate-disable"
	DUPLICATE_ENABLE  = "duplicate-enable"
)
