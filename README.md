# Gotomate Package Template

This is a template of a gotomate's package that you can use to create your own gotomate's packages.

You can choose to create your own functions to do whatever you want with a gotomate process.

## Installation

Download or Clone the repo:
```
git pull https://github.com/pierrelh/gotomate-package-template.git
```

## Setup

Fill the all the files in the Package folder with the right requirements. Fill free to use as an exemple the originals gotomate's packages to help you building this.

Edit the icon.png with an icon corresponding to your package for a better use by users. 

Then rename the "Package" folder to what your package should be named (don't forget to put the folder name on uppercase)

## Installation of the package

Currently, on Windows there is now known way (by me) to use plugin, so you have to add some lines to the gotomate's code in order
to use your package.

### Step 1

Copy & Paste your folder into fiber/packages directory

### Step 2

Go to fiber/fiber.go & add this:
```
case "YourFolderName":
	yourPackageName.Processing(funcName, instructionData, finished)
```
beetween the 105 & 128 lines (at the time I'm writing this), were all the packages are called for the processing.
Be sure that your package is imported at the top of the file.

### Step 3

Go to fiber/packages/packages-dialog.go & add this:
```
case "YourFolderName":
	yourPackageName.Build(funcName)
```
beetween the 96 & 121 lines (at the time I'm writing this), were all the packages are called for the building.
Be sure that your package is imported at the top of the file.


Now the package should be implemented !