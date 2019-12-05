package winsize

import (
	"syscall"
	"unsafe"
)

// Get size, see strace example:
// $ strace -e write,ioctl stty size
// ioctl(0, TCGETS, {B38400 opost isig icanon echo ...}) = 0
// ioctl(1, TIOCGWINSZ, {ws_row=56, ws_col=239, ws_xpixel=0, ws_ypixel=0}) = 0
// ioctl(0, TIOCGWINSZ, {ws_row=56, ws_col=239, ws_xpixel=0, ws_ypixel=0}) = 0
// write(1, "56 239\n", 756 239) = 7
// +++ exited with 0 +++
func Get() (s *Size, err error) {
	s = &Size{}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ), // TIOCGWINSZ - Get window size
		uintptr(unsafe.Pointer(s)),
	)

	if errno != 0 {
		err = errno
	}
	return s, err
}

// Set is used to define a new winsize.
// When the window size changes, a SIGWINCH signal is sent to the foreground
// process group.
// For more infos see: man ioctl_tty
func Set(s *Size) (err error) {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCSWINSZ), // TIOCSWINSZ - Set window size
		uintptr(unsafe.Pointer(s)),
	)

	if errno != 0 {
		err = errno
	}
	return
}
