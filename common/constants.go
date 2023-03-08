package common

import "io/fs"

var DateFormat = "2006-01-02 15:04"
var AppName = "takenotesctl"
var StorageFile = "notes.json"
var ExportFile = "notes.csv"
var Extension = "txt"
var Separator = "#\n"
var DirPermission fs.FileMode = 0777
var FilePermission fs.FileMode = 0644
