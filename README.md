# Batch image resizer
Simple batch image resizer powered by Go.

# Compilation
Go >=1.14 should be installed and ready to compile.

## For windows
```
GOOS=windows go build -ldflags="-s -w"
```
## For macos
```
GOOS=darwin go build -ldflags="-s -w"
```
## For linux
```
GOOS=linux go build -ldflags="-s -w"
```

# Usage
Run binary without parameters to find usage.

## Scale
-scale=<1-99> - scale in percent for new images. Required parameter, no default value.

## Source folder
-src="path" - set source folder for processing. Current folder by default.

## Destination folder
-dst="path" - set destination folder for processing. Source folder by default.

# License
The software released under the terms of the [MIT license](./LICENSE.md). Enjoy.

# Donate
Its free open source software, but you can buy me a coffee:

<a href="https://www.buymeacoffee.com/leopotam" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/yellow_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
