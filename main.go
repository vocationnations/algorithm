// main package is responsible for runing the VocationNations match-making algorithm
package main

import (
	"flag"
	"fmt"
	"github.com/vocationnations/algorithm/constants"
	"github.com/vocationnations/algorithm/factor"
	"regexp"
)

const (
	Culture = "Culture"
	Skill   = "Skill"
)

// Run function returns the percent match for the employee and employers.
func Run(skills, domains []*string, skillPercent, culturePercent float64) (float64, error) {

	score := 0.0
	for k := range constants.AvailableFactors {
		var fac factor.Factor
		switch k {
		case constants.Skill:
			fac = factor.NewFactor(k, skillPercent, skills)
		case constants.Culture:
			fac = factor.NewFactor(k, culturePercent, domains)
		}

		_score, err := fac.Calculate()
		if err != nil {
			return 0.0, fmt.Errorf("ERROR: Failed to calculate score for %s _factor, err: %s", k, err)
		}
		score = score + _score
	}
	return score, nil
}

const (
	// default value for skills count when argument is not provided
	unsetSkillCount = 0

	// default value for skills when argument is not provided
	unsetSkills = ""

	// default value for domain when argument is not provided
	unsetDomain = ""

	// default value for percent when argument is not provided
	unsetPercent = 0.0
)

// isValidDomain returns true if the domain value that is provided is valid, false otherwise
func isValidDomain(domain string) bool {
	var re = regexp.MustCompile(`^([0-9]{1,},){3}[0-9]{1,}$`)
	if re.MatchString(domain) {
		return true
	}
	return false
}

// isValidSkills returns true if the skills tuple is valid, false otherwise
func isValidSkills(skills string, skillCount int) bool {
	var re = regexp.MustCompile(fmt.Sprintf(`^(\([0-9]{1},[0-9]{1,}\);){%d}\([0-9]{1,},[0-9]{1,}\)`, skillCount-1))
	if re.MatchString(skills) {
		return true
	}
	return false
}

// runs the main program which calculates the percent match for given parameters.
func main() {

	var skills, domains []*string

	// create flags for skill count
	skillCountParam := "skill-count"
	skillCountValue := unsetSkillCount
	skillCountUsage := "number of skills required"
	skillCount := flag.Int(skillCountParam, skillCountValue, skillCountUsage)

	skillsParam := "skills"
	skillsValue := unsetSkills
	skillsUsage := "semi-colon separated tuple of int representing skills required and skills available\n" +
		"e.g., (skill_req,skill_avail);(..,..),..."
	skills = append(skills,flag.String(skillsParam, skillsValue, skillsUsage))

	// create flags for all the 6 domains
	domainParam := []string{"d1", "d2", "d3", "d4", "d5", "d6"}
	domainUsage := []string{
		"scores for domain 1 statements",
		"scores for domain 2 statements",
		"scores for domain 3 statements",
		"scores for domain 4 statements",
		"scores for domain 5 statements",
		"scores for domain 6 statements",
	}

	for i, p := range domainParam {
		flg := flag.String(p, unsetDomain, domainUsage[i])
		domains = append(domains, flg)
	}

	skillPercentParam := "skill-percent"
	skillPercentValue := unsetPercent
	skillPercentUsage := "float value representing the worth of skills"
	skillPercent := flag.Float64(skillPercentParam, skillPercentValue, skillPercentUsage)

	culturePercentParam := "culture-percent"
	culturePercentValue := unsetPercent
	culturePercentUsage := "float value representing the worth of skills"
	culturePercent := flag.Float64(culturePercentParam, culturePercentValue, culturePercentUsage)

	// parsing the flags has to be done after setting all the flags
	flag.Parse()

	if *skillPercent == unsetPercent {
		panic("ERROR: --skill-percent argument is required")
	}

	if *skillCount <= 0 || !isValidSkills(*skills[0], *skillCount) {
		panic("ERROR: -skills argument is required and you need to have at least one skill")
	}

	if *culturePercent == unsetPercent {
		panic("ERROR: --culture-percent argument is required")
	}

	// throw an error if any of the domain argument is invalid
	for d := range domains {

		if *domains[d] == unsetDomain {
			// if domain is not set
			panic(fmt.Sprintf("ERROR: -%s argument is required", domainParam[d]))
		} else if !isValidDomain(*domains[d]) {

			// if domain value is not valid
			panic(fmt.Sprintf(
				"ERROR: %s is not a valid argument for %s\n\tproper usage: -%s=number1,number2,number3,number4",
				*domains[d], domainParam[d], domainParam[d]))
		}
	}

	// calculate the percent match
	percentMatch, err := Run(skills, domains, *skillPercent, *culturePercent)
	if err != nil {
		panic(fmt.Sprintf("ERROR: Failed to run the algorithm: %s", err))
	}

	fmt.Printf(_header()+"Percent Match: %0.2f\n", percentMatch)
}

func _header() string {
	return fmt.Sprint(
		"VocationNations (c) 2020. All Rights Reserved \n" +
			"MATCH-MAKING ALGORITHM\n\n")
}
