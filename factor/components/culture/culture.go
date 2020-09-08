// culture package holds the necessary code for calculating the culture match of a candidate
package culture

import (
	"fmt"
	"github.com/vocationnations/algorithm/factor/components"
	"strconv"
	"strings"
)

// cultComponent specifies the fields required for calculating the culture match for the user
type cultComponent struct {
	percent float64
	payload []*string
}

// domainStruct represents the statement scores for a single domain
type domainsStruct struct {
	// the scores for each statement of the domain
	create, collaborate, compete, control float64
}

// formatDomains returns a domainStruct object where each item is float instead of string so that downstream mathematical
// calculations can be performed
func (c cultComponent) formatDomains() ([]domainsStruct, error) {
	var domains []domainsStruct

	for d := range c.payload {
		_constituents := strings.Split(*c.payload[d], ",")

		var _create, _collaborate, _compete, _control float64
		var err error

		if _create, err = strconv.ParseFloat(_constituents[0], 64); err != nil {
			return nil, fmt.Errorf("ERROR: Failed to convert %s to float64", _constituents[0])
		}

		if _collaborate, err = strconv.ParseFloat(_constituents[1], 64); err != nil {
			return nil, fmt.Errorf("ERROR: Failed to convert %s to float64", _constituents[1])
		}
		if _compete, err = strconv.ParseFloat(_constituents[2], 64); err != nil {
			return nil, fmt.Errorf("ERROR: Failed to convert %s to float64", _constituents[2])
		}
		if _control, err = strconv.ParseFloat(_constituents[3], 64); err != nil {
			return nil, fmt.Errorf("ERROR: Failed to convert %s to float64", _constituents[3])
		}

		domains = append(domains,
			domainsStruct{
				create:      _create,
				collaborate: _collaborate,
				compete:     _compete,
				control:     _control,
			})
	}
	return domains, nil
}

// CalculateFactor performs the calculation for culture match and returns a floating point value representing the score
// for the culture match factor
func (c cultComponent) CalculateFactor() (float64, error) {
	var fCreate, fCollaborate, fCompete, fControl float64

	// format the domains for mathematical calculations, throw error if cannot format
	domains,err := c.formatDomains()
	if err != nil {
		return 0.0,fmt.Errorf("ERROR: Cannot calculate factor, err: %s", err)
	}

	// float64 version of domains length
	domainLen := float64(len(domains))

	// accumulate the statement scores for all the domains
	for d := range domains {
		domain := domains[d]
		fCreate = fCreate + domain.create
		fCollaborate = fCollaborate + domain.collaborate
		fCompete = fCompete + domain.compete
		fControl = fControl + domain.control
	}

	finalScore := (fCreate/domainLen) + (fCollaborate/domainLen) + (fCompete/domainLen) + (fControl/domainLen)
	weightedFinalScore := finalScore * c.percent

	return weightedFinalScore, nil
}

// NewCultureComponent creates a new object which returns a Component object for culture factor
func NewCultureComponent(pcent float64, payload []*string) components.Component {
	return &cultComponent{
		percent: pcent,
		payload: payload,
	}
}
