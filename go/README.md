# hwebview

This Go package implements the host-side of the Flutter [hwebview](https://github.com/Winwisly100/hwebview) plugin.

## Usage

Import as:

```go
import hwebview "github.com/Winwisly100/hwebview/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&hwebview.HwebviewPlugin{}),
```
