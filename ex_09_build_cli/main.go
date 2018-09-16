package main

import (
	"io/ioutil"
	"flag"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"os"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"fmt"
	"github.com/janeczku/go-spinner"
	"os/user"
)

func main() {
	var song string

	user, _ := user.Current()
	defaultPath := user.HomeDir + "/Music/"

    flag.StringVar(&song, "song", "", "song name")
	flag.Parse()

	songPath :=  defaultPath + song + ".mp3"
	//-----------------------------------------------
	if _, err := os.Stat(songPath); err == nil {
		
		if _, fileErr := os.Stat("/tmp/banner.txt"); os.IsNotExist(fileErr) {
			bannerTxtBytes := MustAsset("banner.txt")
			errRewrite := ioutil.WriteFile("/tmp/banner.txt", bannerTxtBytes, 0777)
			if errRewrite != nil {
				fmt.Println(errRewrite)
			}
		}
		in, _ := os.Open("/tmp/banner.txt")
		fmt.Println()
		banner.Init(colorable.NewColorableStdout(), true, true, in)
		fmt.Println()
		s := spinner.StartNew(song + ".mp3 now playing... ♪ ♫ ♬ ")
		var file, _ = ioutil.ReadFile(defaultPath + song + ".mp3")
		dec, data, _ := minimp3.DecodeFull(file)
		player, _ := oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 1024)
		player.Write(data)
		s.Stop()
	} else {
		fmt.Println(song + " song could not be found !")
	}
}