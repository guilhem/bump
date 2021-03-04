package semver

import (
	"fmt"

	"github.com/Masterminds/semver"
)

type bump = semver.Version

func New(v string) *bump {
	return semver.MustParse(v)
}

func Latest(tags []string) (string, error) {
	vs := make([]*semver.Version, len(tags))

	for i, r := range tags {
		v, err := semver.NewVersion(r)
		if err != nil {
			return "", fmt.Errorf("can't parse tag %s semver: %w", r, err)
		}

		vs[i] = v
	}

	latest := vs[0]

	for _, t := range semver.Collection(vs) {
		if t.GreaterThan(latest) {
			latest = t
		}
	}

	return latest.String(), nil
}
