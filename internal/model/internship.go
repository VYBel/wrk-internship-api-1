package model

import "strconv"

type Internship struct {
	Id           uint64
	Team_id      uint64
	Description  string
	Period       string
	Compensation bool
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type InternshipEvent struct {
	Id     uint64
	Type   EventType
	Status EventStatus
	Entity *Internship
}

func (internship *Internship) String() string {
	var result string
	result += "id: " + strconv.FormatUint(internship.Id, 10) + ":"
	result += " team_id - " + strconv.FormatUint(internship.Team_id, 10) + ";"
	if internship.Description != "" {
		result += " Description: " + internship.Description + ";"
	}
	if internship.Period != "" {
		result += " Period: " + internship.Period + ";"
	}
	if internship.Compensation {
		result += " compensation: yes."
	} else {
		result += " compensation: no."
	}
	return result
}
