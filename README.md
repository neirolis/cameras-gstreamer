# Gstremer cameras module

Returns frames from rtsp stream to the stdout.

## Install

```shell
go build -o gstreamer
mkdir -p $RTMIPDIR/cameras/gstreamer
mv gstreamer $RTMIPDIR/cameras/gstreamer/
cp manifest.json $RTMIPDIR/cameras/gstreamer/
```

Or from the package [cameras-gstreamer.yaml](https://github.com/neirolis/repository/blob/master/modules/cameras-gstreamer.yaml)

## Requirements

- gst-launch-1.0
- gst-plugins-good


## Manual Usage

```
usage: gstreamer <addr> [--framerate=<n>]

positional:
  addr                    rtsp stream address

options:
      --framerate=<n>     target fps [default: 5]
  -h, --help              display this help and exit
```
