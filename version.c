#include <parsebgp.h>

#define STR(s) _STR(s)
#define _STR(s) #s

char* gpb_parsebgp_version() {
  return STR(LIBPARSEBGP_MAJOR_VERSION) "."
         STR(LIBPARSEBGP_MID_VERSION) "."
         STR(LIBPARSEBGP_MINOR_VERSION);
}
