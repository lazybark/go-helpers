package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemverString(t *testing.T) {
	v011 := Ver{
		Major:         0,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "",
		BuildMetadata: "",
		Stable:        true,
	}
	assert.Equal(t, "0.1.1", v011.String())
	assert.Equal(t, false, v011.IsStable())

	v111 := Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "",
		BuildMetadata: "",
		Stable:        true,
	}
	assert.Equal(t, "1.1.1", v111.String())
	assert.Equal(t, true, v111.IsStable())
	v111.Stable = false
	assert.Equal(t, false, v111.IsStable())

	v111_rn := Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "alpha.3",
		BuildMetadata: "",
		Stable:        true,
	}
	assert.Equal(t, "1.1.1-alpha.3", v111_rn.String())
	assert.Equal(t, false, v111_rn.IsStable())

	v111_rn_bm := Ver{
		Major:         1,
		Minor:         1,
		Patch:         1,
		ReleaseNote:   "alpha.4",
		BuildMetadata: "j32edbshdi3ws",
		Stable:        true,
	}
	assert.Equal(t, "1.1.1-alpha.4+j32edbshdi3ws", v111_rn_bm.String())
	assert.Equal(t, false, v111_rn_bm.IsStable())

	assert.Equal(t, true, v111_rn_bm.IsHigher(v111_rn))
	assert.Equal(t, false, v111_rn.IsHigher(v111))
	assert.Equal(t, true, v111.IsHigher(v011))

	//Versions are comapred by ReleaseNote parts in case Major&Minor&Patch are the same
	assert.Equal(t, false, v111_rn_bm.IsEqual(v111_rn))
	//Versions with same length, but different BuildMetadata are the same
	v111_rn.ReleaseNote = "alpha.4"
	assert.Equal(t, true, v111_rn_bm.IsEqual(v111_rn))
	//But longer versions are higher
	v111_rn.ReleaseNote = "alpha.4.11"
	assert.Equal(t, false, v111_rn_bm.IsEqual(v111_rn))
	//Increasing any number makes version higher
	v111_rn.Minor = 2
	assert.Equal(t, false, v111_rn_bm.IsEqual(v111_rn))
	assert.Equal(t, true, v111_rn.IsHigher(v111_rn_bm))

	//Check base compatibility
	assert.Equal(t, false, v111_rn_bm.IsCompatible(v011))
	assert.Equal(t, true, v111_rn_bm.IsCompatible(v111_rn))
	//Check manually marked as incompatible
	v111_rn_bm.InComp = append(v111_rn_bm.Comp, v111_rn)
	assert.Equal(t, true, v111_rn_bm.IsCompatible(v111_rn))
	//Check manually marked as compatible
	v111_rn_bm.Comp = append(v111_rn_bm.Comp, v011)
	assert.Equal(t, true, v111_rn_bm.IsCompatible(v011))
}
