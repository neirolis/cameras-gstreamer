package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/mattn/go-shellwords"
	"github.com/sg3des/argum"
)

var args struct {
	Addr      string `argum:"pos,req" help:"rtsp stream address"`
	FrameRate int    `argum:"--framerate" help:"target fps" default:"5"`
}

var execName = "gst-launch-1.0"
var execTmpl = "rtspsrc location={{.Addr}} latency=10 ! queue ! rtph264depay ! h264parse ! avdec_h264 ! videoconvert ! videoscale ! videorate ! video/x-raw,framerate={{.FrameRate}}/1 ! jpegenc quality=95 ! filesink location=/dev/stdout buffer-size=0 buffer-mode=2"

func init() {
	argum.MustParse(&args)
}

func main() {
	tmpl, err := template.New("gst").Parse(execTmpl)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(buf, args); err != nil {
		log.Fatal(err)
	}

	execLine := buf.String()
	args, err := shellwords.Parse(execLine)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(execName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
