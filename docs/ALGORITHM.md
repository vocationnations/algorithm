# Algorithm

# Table of contents

1. [Introduction](#1-introduction)  
    1.1 [Culture Match Calculation](#11-culture-match-calculation)  
        1.1.1 [Domain Representation](#111-domain-representation)  
        1.1.2 [Statement Representation](#112-statement-representation)  
        1.1.3 [Bin Representation](#113-bin-representation)  
        1.1.4 [Final Score Representation](#114-final-score-representation)  
        1.1.5 [Pseudo-code Representation](#115-pseudo-code-representation)      
    1.2 [Skills Match Calculation](#12-skills-match-calculation)  
        1.2.1 [Skill Representation](#121-skill-representation)  
        1.2.2 [Pseudo-code Representation](#122-pseudo-code-representation)  
    1.3 [Percent Match Calculation](13-percent-match-calculation)
        

## Introduction
The following algorithm calculates the percent match between an employer and an employee. In order to figure out the 
over all match, we take into account the culture match as well as the skills match. Each of those are calculated as follows:

### Culture Match Calculation
There are 6 domains that we use to measure the user's culture match. Each domain has 4 statements types `create`, `collaborate`
`control` and `compete`. The statements for each domain are ranked into a bin, whereby each statement type
gets a score. The scores of all the statements of all the domains are accumulated to come up with a final score.

#### Domain Representation

```
type Domain struct {
    name string
    create Statement
    collaborate Statement
    control Statement
    compete Statement   
}
```

#### Statement Representation

```
type Statement struct {
    stype string
    content string
    score float64
}
```

#### Bin Representation

```
type Bin struct {
    name string
    value float64
}
```

#### Final Score Representation
```
type FinalScore struct {
    create, collaborate, compete, control float64 
} 
```
#### Pseudo-code Representation

The implementation will have following functions:
1. `setStatementToBin(statement,domain_index,bin_name)` -- in the domains object, updates score of a particular statement
2. `calculateFinalScore(domains)` -- calculates the cumulative final score
3. `compareTwoFinalScores(fscore1,fscore2)` -- obtains a differential final score between two entities
4. `averageFinalScore(fscore)` -- calculates and returns a single value representing average of the finalscores for each statement.


### Skills Match Calculation

#### Skill Representation
```
type Skill struct {
    name string
    level int
}
```

#### Pseudo-code Representation
The implementation will have following functions:

1. `getSkillDifference(skill1,skill2)` -- obtains the absolute distance in level between two skills
2. `getSkillDifferenceBetweenEntities(entity1_skills, entity2_skills)` -- obtains an array representing absolute difference values between skills of two entities.
3. `getSkillMatchScore(skill_difference_arraY)` -- obtains a single value representing average of skill differences between two entities


### Percent Match Calculation

Percent match between two entities will be calculated through the following steps:

1. `averageCultureFinalScore` -> Calculate Culture Match FinalScore using `averageFinalScore` function
2. `skillMatchScore` -> Calculate skill match score using `getSkillMatchScore` function
3. `percentMatch` -> (`averageCultureFinalScore` * `importanceOfCulture`) + (`skillMatchScore` * `importanceOfSkills`)
4. Use the `percentMatch` in the downstream processes.

