//(c) Lorenzo Iannone - lorenzo.iannone@gmail.com
//released under GNU Lesser GPL License

package rpi

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPiSPI.h>
*/
import 	"C"

import (
	"unsafe"
	"errors"
)

//WiringPiSPISetup is the wrapper of 
//int wiringPiSPISetup(int channel, int speed)
func WiringPiSPISetup (channel int, speed int) error {
	if -1 == int(C.wiringPiSPISetup (C.int(channel), C.int(speed))) {
		return errors.New ("WiringPiSPISetup failed")
	}
	return nil
}

//WiringPiSPISetupmode is the wrapper of 
//int wiringPiSPISetupMode (int channel, int speed, int mode) 
func WiringPiSPISetupMode (channel int, speed int, mode int) error {
	if -1 == int(C.wiringPiSPISetupMode (C.int(channel), C.int(speed), C.int(mode))) {
		return errors.New ("wiringPiSPISetupMode failed")
	}
	return nil
}

//WiringPiSPIDataRW is the wrapper of 
//int wiringPiSPIDataRW (int channel, unsigned char *data, int len)
func WiringPiSPIDataRW(channel int, data *[]byte, len int) error{
	if -1 ==  C.wiringPiSPIDataRW(C.int(channel), (*C.uchar)(unsafe.Pointer(data)), C.int(len)) {
		return errors.New("wiringPiSPIDataRW failed")
	}
	return nil


}

//WiringPiSPIGetFd is the wrapper of
//int wiringPiSPIGetFd(int channel) 
func WiringPiSPIGetFd(channel int) error {
	if -1 == C.wiringPiSPIGetFd(C.int(channel)) {
		return errors.New("wiringPiSPIGetFd failed")
	}
	return nil
}





