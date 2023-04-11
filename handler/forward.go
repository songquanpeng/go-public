package handler

import (
	"io"
	"net"
)

func forward(src, dst net.Conn) {
	defer src.Close()
	defer dst.Close()
	_, _ = io.Copy(src, dst)
}
