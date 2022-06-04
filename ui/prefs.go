package main

import (
	"fmt"
	"image/color"
	"os"

	fyne "fyne.io/fyne/v2"
	canvas "fyne.io/fyne/v2/canvas"
)

// various sizes of our application
var windowSize, inputSize, rowSize fyne.Size
var appKind, appTheme string
var gitImage, docImage, webImage, lockImage, syncImage, passImage, listImage *canvas.Image
var rightArrowImage, leftArrowImage *canvas.Image
var btnColor, editColor, updateColor, redColor color.NRGBA

// helper function to set application preferences/settings
func appSettings(app fyne.App) {
	// set some values for our app preferences
	pref := app.Preferences()

	// default values
	cipher := "aes"
	vdir := fmt.Sprintf("%s/.ecm/Primary", os.Getenv("HOME"))
	fontSize := "Normal"
	appTheme = pref.String("AppTheme")
	if appTheme == "" {
		appTheme = "dark"
	}

	// color for our buttons
	redColor = color.NRGBA{0xff, 0x26, 0x00, 0xff}
	//     redColor = color.NRGBA{0xff, 0x2f, 0x92, 0xff}
	btnColor = color.NRGBA{0x79, 0x79, 0x79, 0xff}
	editColor = color.NRGBA{0x00, 0x8f, 0x00, 0xff}
	updateColor = color.NRGBA{0x04, 0x33, 0xff, 0xff}

	// rowSize represents main row size used in our UI containers
	rowSize = fyne.NewSize(340, 40)
	// input size represents size of the input field which is shorter by row size
	inputSize = fyne.NewSize(340, 40)

	// make changes depending on application kind
	if appKind == "desktop" {
		windowSize = fyne.NewSize(900, 600)
	} else {
		vdir = app.Storage().RootURI().Path()
		windowSize = fyne.NewSize(300, 600)
		// on mobile input size should be short of rowSize
		inputSize = fyne.NewSize(300, 40)
	}

	// save preferences
	pref.SetString("VaultCipher", cipher)
	pref.SetString("VaultDirectory", vdir)
	pref.SetString("FontSize", fontSize)
	pref.SetString("AppTheme", appTheme)
	if autoThreshold != nil {
		thr, err := autoThreshold.Get()
		if err == nil {
			pref.SetString("Autologout", thr)
		}
	}

	// set images
	setCustomImages()

	// write ecmconfig
	WriteSyncConfig(app)
}

// helper function to set custom images for our app theme
func setCustomImages() {
	if appTheme == "light" {
		gitImage = canvas.NewImageFromResource(resourceGithubBlackSvg)
		webImage = canvas.NewImageFromResource(resourceWebBlackSvg)
		docImage = canvas.NewImageFromResource(resourceDocBlackSvg)
		lockImage = canvas.NewImageFromResource(resourceLockBlackSvg)
		syncImage = canvas.NewImageFromResource(resourceSyncBlackSvg)
		passImage = canvas.NewImageFromResource(resourcePassBlackSvg)
		listImage = canvas.NewImageFromResource(resourceListBlackSvg)
		leftArrowImage = canvas.NewImageFromResource(resourceLeftArrowBlackSvg)
		rightArrowImage = canvas.NewImageFromResource(resourceRightArrowBlackSvg)
	} else {
		gitImage = canvas.NewImageFromResource(resourceGithubWhiteSvg)
		webImage = canvas.NewImageFromResource(resourceWebWhiteSvg)
		docImage = canvas.NewImageFromResource(resourceDocWhiteSvg)
		lockImage = canvas.NewImageFromResource(resourceLockWhiteSvg)
		syncImage = canvas.NewImageFromResource(resourceSyncWhiteSvg)
		passImage = canvas.NewImageFromResource(resourcePassWhiteSvg)
		listImage = canvas.NewImageFromResource(resourceListWhiteSvg)
		leftArrowImage = canvas.NewImageFromResource(resourceLeftArrowWhiteSvg)
		rightArrowImage = canvas.NewImageFromResource(resourceRightArrowWhiteSvg)
	}
	//     gitImage.SetMinSize(fyne.NewSize(100, 100))
	//     webImage.SetMinSize(fyne.NewSize(100, 100))
	//     docImage.SetMinSize(fyne.NewSize(100, 100))
	//     lockImage.SetMinSize(fyne.NewSize(100, 100))
	rightArrowImage.SetMinSize(fyne.NewSize(50, 50))
	leftArrowImage.SetMinSize(fyne.NewSize(50, 50))
}