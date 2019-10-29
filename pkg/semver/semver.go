package semver

import "github.com/Masterminds/semver"

type bump = semver.Version

func New(v string) *bump {
	return semver.MustParse(v)
}
