// skills package holds the necessary code for calculating the skill match of a candidate
package skills

import (
	"fmt"
	"github.com/vocationnations/algorithm/factor/components"
	"math"
	"strconv"
	"strings"
)

// skillComponent specifies the fields required for calculating the skill match for the user
type skillComponent struct {
	percent float64
	payload []*string
}

// skillsStruct specifies the fields for the skills that the user passes in
type skillsStruct struct {
	// required skills from the employer
	required int

	// available skills of the employee
	available int
}

// formatDomains returns a skillsStruct object where each item is int instead of string so that downstream mathematical
// calculations can be performed
func (c skillComponent) formatSkills() ([]skillsStruct, error) {

	var skills []skillsStruct
	_skills := strings.Split(*c.payload[0], ";")
	for s := range _skills {

		var req, avail int
		var err error

		// extract the required value from tuple
		_req := strings.Split(_skills[s], ",")[0][1:]

		// extract the available value from tuple
		__avail := strings.Split(_skills[s], ",")[1]
		_avail := __avail[:len(__avail)-1]

		// check if you can convert required to int
		if req, err = strconv.Atoi(_req); err != nil {
			return nil,fmt.Errorf("ERROR: Cannot convert %s to int",_req)
		}

		// check if you can convert available to int
		if avail, err = strconv.Atoi(_avail); err != nil {
			return nil,fmt.Errorf("ERROR: Cannot covert %s to int", _avail)
		}

		// add newly converted required and available to skillStruct array
		skills = append(skills,
			skillsStruct{
				required:  req,
				available: avail,
			},
		)
	}
	return skills,nil
}

// CalculateFactor performs the calculation for skill match and returns a floating point value representing the difference
// between the employer requirements and the employee available skills.
// NOTE: The lower the difference the better the match
func (c skillComponent) CalculateFactor() (float64, error) {
	// final difference value will be stored here
	var diff float64

	// format skills for mathematical calculations
	skills,err := c.formatSkills()
	if err != nil {
		return 0.0, fmt.Errorf("ERROR: Cannot format skills, err: %s", err)
	}

	// go through all the skills and find the absolute difference between the requirement level and available skill
	// that the employee has
	for s := range skills {
		skill := skills[s]
		diff = diff + math.Abs(float64(skill.required)-float64(skill.available))
	}
	weightedFinalScore := diff * c.percent
	return weightedFinalScore, nil
}

// NewSkillComponent creates a new object which returns a Component object for skill factor
func NewSkillComponent(pcent float64, payload []*string) components.Component {
	return &skillComponent{
		percent: pcent,
		payload: payload,
	}
}
