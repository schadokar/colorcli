# colorcli

The `colorcli` is to return the name of the hexcode of the color. If the hexcode doesn't exist then `colorcli` can add the color.

The complete tutorial is published on [codesource.io]().

## Instruction to run the cli

1. Download the `colornames.min.json` from the [color-name](https://github.com/meodai/color-names#cdn-) git repo.

```
curl -O https://github.com/meodai/color-names#cdn-
```
2. go install colorcli
3. get the color name

```
$ colorcli name ffffff
Name: White, Hex: ffffff
```
4. add the color to the colornames file.

```
$ colorcli addcolor fffeee lightcream
Hex to color added successfully!
```

## Author

Hi there, I am Shubham Chadokar. I write articles and tutorials on Blockchain, golang, nodejs and reactjs. You can read my other articles and tutorials on my [blog](https://schadokar.dev).
