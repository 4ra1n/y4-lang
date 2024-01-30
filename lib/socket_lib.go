package lib

import (
	"encoding/hex"
	"net"
	"strconv"
	"time"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
)

const (
	SocketLibName = "骚客特"
)

var (
	SocketLib = base.NewList[*native.NativeFunction]()
)

func init() {
	connectFun, err := native.NewNativeFunction(SocketLibName+SEP+"连接", connect)
	if err != nil {
		return
	}
	closeConnectFun, err := native.NewNativeFunction(SocketLibName+SEP+"关闭", closeConnect)
	if err != nil {
		return
	}
	writeHexFun, err := native.NewNativeFunction(SocketLibName+SEP+"写入海克斯", writeHex)
	if err != nil {
		return
	}
	readHexFun, err := native.NewNativeFunction(SocketLibName+SEP+"读取海克斯", readHex)
	if err != nil {
		return
	}
	writeStrFun, err := native.NewNativeFunction(SocketLibName+SEP+"写入字符串", writeStr)
	if err != nil {
		return
	}
	readStrFun, err := native.NewNativeFunction(SocketLibName+SEP+"读取字符串", readStr)
	if err != nil {
		return
	}

	SocketLib.Add(connectFun)
	SocketLib.Add(closeConnectFun)
	SocketLib.Add(writeHexFun)
	SocketLib.Add(readHexFun)
	SocketLib.Add(writeStrFun)
	SocketLib.Add(readStrFun)
}

func connect(ip string, port int) net.Conn {
	if port < 1 || port > 65535 {
		log.Error("socket lib error: port invalid")
		return nil
	}
	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	if err != nil {
		log.Warnf("socket lib error: %s", err.Error())
		return nil
	}
	return conn
}

func closeConnect(conn net.Conn) {
	if conn == nil {
		log.Error("socket lib error: connection is null")
		return
	}
	err := conn.Close()
	if err != nil {
		log.Errorf("socket lib error: %s", err.Error())
	}
}

func writeHex(conn net.Conn, h string) {
	data, err := hex.DecodeString(h)
	if err != nil {
		log.Errorf("socket lib decode error: %s", err.Error())
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		log.Errorf("socket lib write error: %s", err.Error())
		return
	}
}

func readHex(conn net.Conn, size int) string {
	if size <= 0 {
		log.Errorf("invalid size: %d", size)
		return ""
	}
	result := make([]byte, size)
	n, err := conn.Read(result)
	if err != nil {
		log.Errorf("socket lib read error: %s", err.Error())
		return ""
	}
	return hex.EncodeToString(result[:n])
}

func writeStr(conn net.Conn, data string) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		log.Errorf("socket lib write error: %s", err.Error())
		return
	}
}

func readStr(conn net.Conn, size int) string {
	if size <= 0 {
		log.Errorf("invalid size: %d", size)
		return ""
	}
	result := make([]byte, size)
	n, err := conn.Read(result)
	if err != nil {
		log.Errorf("socket lib read error: %s", err.Error())
		return ""
	}
	return string(result[:n])
}
