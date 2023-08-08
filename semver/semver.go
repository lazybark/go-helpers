// semver is a package that implements Semantic Versioning.
// https://semver.org/
//
// It provides tools to set and compare versions.
package semver

import (
	"fmt"
	"strconv"
	"strings"
)

// SemVer represents Semantic Versioning document version
var SemVer = Ver{
	Major: 2,
	Minor: 0,
	Patch: 0,
}

// Ver implements Semantic Versioning https://semver.org/
type Ver struct {
	Major uint
	Minor uint
	Patch uint

	//ReleaseNote defines comment to release, e.g. "alpha", "beta.2"
	ReleaseNote string

	//BuildMetadata represents commit hash or any comment to current build
	BuildMetadata string

	//Stable is true in case current app ver works exactly as expecterd.
	//Correct way to ckeck Stable is via Ver.IsStable()
	Stable bool

	//Comp holds list of versions that are totally compatible with current
	Comp []Ver

	//InComp holds list of versions that are totally incompatible with current
	InComp []Ver
}

// String returns string-formatted semver
func (v Ver) String() string {
	rn := ""
	if v.ReleaseNote != "" {
		rn = "-" + v.ReleaseNote
	}
	bm := ""
	if v.BuildMetadata != "" {
		bm = "+" + v.BuildMetadata
	}
	return fmt.Sprintf("%v.%v.%v%s%s", v.Major, v.Minor, v.Patch, rn, bm)
}

// IsStable returns true in case Ver is stable. False is always returned
// in case v.Major is 0 or v.ReleaseNote is not empty. In all other cases
// value of v.Stable is returned.
//
// https://semver.org/#spec-item-9
//
// https://semver.org/#spec-item-4
func (v Ver) IsStable() bool {
	if v.Major == 0 {
		return false
	}
	if v.ReleaseNote != "" {
		return false
	}

	return v.Stable
}

// IsHigher compares versions by rules of Semantic versioning and returns
// true in case v > c
func (v Ver) IsHigher(c Ver) bool {
	//Compare major
	if v.Major > c.Major {
		return true
	} else if v.Major < c.Major {
		return false
	}
	//Compare minor
	if v.Minor > c.Minor {
		return true
	} else if v.Minor < c.Minor {
		return false
	}
	//Compare patches
	if v.Patch > c.Patch {
		return true
	} else if v.Patch < c.Patch {
		return false
	}

	//Ver with non-empty ReleaseNote is considered lower
	if v.ReleaseNote != "" && c.ReleaseNote == "" {
		return false
	} else if c.ReleaseNote != "" && v.ReleaseNote == "" {
		return true
	}
	//If both versions have non-empty ReleaseNote, it should be
	//compared by ASCII sorting.
	a := strings.Split(v.ReleaseNote, ".")
	b := strings.Split(c.ReleaseNote, ".")

	//Longer version is considered higher
	if len(a) > len(b) {
		return true
	} else if len(b) > len(a) {
		return false
	}

	//If ReleaseNote fields have same length
	for n, val := range a {
		i1, err1 := strconv.Atoi(val)
		i2, err2 := strconv.Atoi(b[n])
		//If both ReleaseNote parts are ints - compare numbers
		if err1 == nil || err2 == nil {
			if i1 > i2 {
				return true
			} else if i2 > i1 {
				return false
			}
		} else {
			//Else - compare alphabetically
			if val > b[n] {
				return true
			} else if b[n] > val {
				return false
			}
		}
	}

	return false
}

// IsEqual returns true in case versions are equal
func (v Ver) IsEqual(c Ver) bool {
	if v.IsHigher(c) {
		return false
	}
	if c.IsHigher(v) {
		return false
	}

	return true
}

// IsCompatible returns true in case versions should be compatible
func (v Ver) IsCompatible(c Ver) bool {
	if v.IsEqual(c) {
		return true
	}
	//Even if Major versions are the same, there is still possibility
	//of some weird situations. So it's better to check InComp versions
	for _, iv := range v.InComp {
		if iv.IsEqual(c) {
			return true
		}
	}
	//If c is not in InComp then we can just compare Major versions
	if v.Major == c.Major {
		return true
	}

	//Last try - c can be manually marked as compatible to v
	for _, cv := range v.Comp {
		if cv.IsEqual(c) {
			return true
		}
	}

	return false
}
