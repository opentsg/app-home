//  Copyright ©2017-2025  Mr MXF   info@mrmxf.com
//  BSD-3-Clause License           https://opensource.org/license/bsd-3-clause/

// Package care manages the embedded files of clog
// if you are forking clog, then you probably want to override this package
// to include ONLY the files that you want in your own fork.

package www

import (
	"embed"
)

//variable CoreFs is the exported file system

//go:embed all:r
//go:embed all:templates
var StaticFs embed.FS
