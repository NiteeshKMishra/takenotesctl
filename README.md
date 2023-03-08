# takenotesctl

takenotesctl is a cli tool to manage notes

## Overview

takenotesctl is cli app to and view notes from cli. Notes can be filtered by created date,
can be searched by matching fields text and can also be exported in a csv file.

#### Installation

```
brew tap NiteeshKMishra/takenotesctl \
https://github.com/NiteeshKMishra/takenotesctl
brew install takenotesctl
```

#### Usage

```
takenotesctl [command]
```

#### Available Commands

```
add         Add a note
list        List all notes
search      Search notes
export      Export notes to a csv file
help        Help about any command
```

#### Flags

```
-h, --help help for takenotesctl
```

## Local Setup

Install Cobra cli locally in your system, this will help in bootstrap application scaffolding locally

```
go install github.com/spf13/cobra@latest
```

Clone this repo locally

```
https://github.com/NiteeshKMishra/takenotesctl.git
```

cd into this repo and run below command to install locally and use this app

```
go install
takenotesctl --help
```

Run all tests with

```
go test ./...
```
