// factor is responsible for processing the calculations for all the factors involved in calculating the percent match
package factor

import (
	"fmt"
	"github.com/vocationnations/algorithm/constants"
	"github.com/vocationnations/algorithm/factor/components/culture"
	"github.com/vocationnations/algorithm/factor/components/skills"
)

// factor implements the Factors interface for operations on the factor that are involved in calculating the
// percent match
type (
	factor struct {
		name    string   // a string identifier of the factor component
		percent float64  // the percent worth for this factor
		payload []*string // any payload that is used to calculate this factor
	}
)

func (f *factor) Calculate() (float64, error) {
	switch f.name {
	case constants.Culture:

		cultureComponent := culture.NewCultureComponent(f.percent,f.payload)
		score,err := cultureComponent.CalculateFactor()
		if err != nil {
			return 0.0, fmt.Errorf("ERROR: Unable to calculate culture match, err: %s", err)
		}
		return score,nil

	case constants.Skill:

		skillComponent := skills.NewSkillComponent(f.percent, f.payload)
		score, err := skillComponent.CalculateFactor()
		if err != nil {
			return 0.0, fmt.Errorf("ERROR: Unable to calculate skill match, err: %s", err)
		}
		return score, nil

	default:
		return 0.0, fmt.Errorf("ERROR: Unknown factor %s was provided to calclate function", f.name)
	}
}

// NewFactor creates and returns a new factor struct
func NewFactor(fname string, percent float64, payload []*string) Factor {
	return &factor{
		name:    fname,
		percent: percent,
		payload: payload,
	}
}
