package all

import (
	// blank import to ensure init() gets called for each package so it can
	// be properly registered.
	_ "github.com/chilons/transporter/function/gojajs"
	_ "github.com/chilons/transporter/function/omit"
	_ "github.com/chilons/transporter/function/opfilter"
	_ "github.com/chilons/transporter/function/ottojs"
	_ "github.com/chilons/transporter/function/pick"
	_ "github.com/chilons/transporter/function/pretty"
	_ "github.com/chilons/transporter/function/remap"
	_ "github.com/chilons/transporter/function/rename"
	_ "github.com/chilons/transporter/function/skip"
)
