package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/semver"
)

func main() {
	v011 := semver.Ver{
		Major:         0,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "",
		BuildMetadata: "",
		Stable:        true,
	}
	fmt.Printf("Is %s stable? %v\n", v011, v011.IsStable())

	v111 := semver.Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "",
		BuildMetadata: "",
		Stable:        true,
	}
	fmt.Printf("Is %s stable? %v\n", v111, v111.IsStable())
	v111.Stable = false
	fmt.Printf("And now is %s stable? %v\n", v111, v111.IsStable())

	v111_rn := semver.Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "alpha.3",
		BuildMetadata: "",
		Stable:        true,
	}
	fmt.Printf("Is %s stable? %v\n", v111_rn, v111_rn.IsStable())

	v111_rn_bm := semver.Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "alpha.4",
		BuildMetadata: "j32edbshdi3ws",
		Stable:        true,
	}
	fmt.Printf("Is %s stable? %v\n", v111_rn_bm, v111_rn_bm.IsStable())

	fmt.Printf("Is %s > %s? %v\n", v111_rn_bm, v111_rn, v111_rn_bm.IsHigher(v111_rn))
	fmt.Printf("Is %s > %s? %v\n", v111_rn, v111, v111_rn.IsHigher(v111))
	fmt.Printf("Is %s > %s? %v\n", v111, v011, v111.IsHigher(v011))
}
