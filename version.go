package goparsebgp

// #include <parsebgp.h>
// #include "bridge.h"
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
			Version: C.GoString(C.gpb_parsebgp_version()),
			Commit:  _LibParseBGPCommit,
		},
		GoParseBGP: VersionDetails{
			Version: _GoParseBGPVersion,
			Commit:  _GoParseBGPCommit,
		},
	}
)
