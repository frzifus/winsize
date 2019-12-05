// Package winsize provides the Get, Set function to read and write the
// tty window size for all supported systems.
//
// Furthermore, no external dependencies are used.
//
// +build amd64,arm
package winsize

// Size is the go implementation of the structure used by ioctl.
// The original struct used by these ioctls is defined as
// struct winsize {
//	unsigned short ws_row;
//	unsigned short ws_col;
//	unsigned short ws_xpixel;   /* unused */
//	unsigned short ws_ypixel;   /* unused */
// };
type Size struct {
	Row uint16
	Col uint16
	X   uint16 // unused
	Y   uint16 // unused
}
