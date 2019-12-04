package main

import (
	"WhatsappProfile/wppprofilebot"
	"bytes"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/nfnt/resize"
	"image/jpeg"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// JPG images
// preview 96x96
// fullsize 640x640

func main() {

	// Timeout to exit
	go func() {
		time.Sleep(time.Second * 30)
		os.Exit(0)
	}()

	if len(os.Args) > 1 {
		switch os.Args[1] {
			case "--help", "-h": {
				fmt.Println("Help: --use-exe-dir to use exe location as working dir")
				os.Exit(0)
			}
			case "--use-exe-dir": {
				err := os.Chdir(filepath.Dir(os.Args[0]))
				if err != nil {
					panic(err)
				}
			}
		}
	}
	rand.Seed(time.Now().Unix())

	img := wppprofilebot.DoNextStep()

	wac, err := whatsapp.NewConn(20 * time.Second)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}
	_ = wac.SetClientName("Profile pic updater", "picupdate")

	err = login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	var big, small bytes.Buffer
	_ = jpeg.Encode(&big, resize.Resize(640, 640, img, resize.Bicubic), nil)
	_ = jpeg.Encode(&small, resize.Resize(96, 96, img, resize.Bicubic), nil)

	uploadReturnChan, err := wac.UploadProfilePic(big.Bytes(), small.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error uploading picture: %v\n", err)
		return
	}

	<-uploadReturnChan

	_, _ = wac.Disconnect()

	os.Exit(0)
}
