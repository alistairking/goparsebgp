package goparsebgp

// #cgo CFLAGS: -I./build/include
// #cgo LDFLAGS: ./build/lib/libparsebgp.a
// #include <parsebgp.h>
//
// #define STR(s) _STR(s)
// #define _STR(s) #s
//
// static char* parsebgp_version() {
//   return STR(LIBPARSEBGP_MAJOR_VERSION) "."
//          STR(LIBPARSEBGP_MID_VERSION) "."
//          STR(LIBPARSEBGP_MINOR_VERSION);
// }
import "C"

var (
	// populated by the linker:
	_LibParseBGPCommit = "unknown"

	_GoParseBGPVersion = "0.0.0"
	_GoParseBGPCommit  = "unknown"
)

type VersionDetails struct {
	Version string
	Commit  string
}

type PkgVersions struct {
	LibParseBGP VersionDetails
	GoParseBGP  VersionDetails
}

var (
	Version = PkgVersions{
		LibParseBGP: VersionDetails{
			Version: C.GoString(C.parsebgp_version()),
			Commit:  _LibParseBGPCommit,
		},
		GoParseBGP: VersionDetails{
			Version: _GoParseBGPVersion,
			Commit:  _GoParseBGPCommit,
		},
	}
)
