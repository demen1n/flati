package finfo

import (
	"errors"
	"strconv"
	"strings"
)

func parseCian(s string) (*Entry, error) {
	// find cian id
	// example: "cianId":315832471
	ssi := "\"cianId\":"
	i := strings.Index(s, ssi)
	if i == -1 {
		return nil, errors.New("no cian id in input")
	}
	t := s[i+len(ssi):]
	j := strings.Index(t, ",")
	id, err := strconv.Atoi(t[:j])
	if err != nil {
		return nil, err
	}

	// find price
	ssp := "\"priceTotalRur\":"
	i = strings.Index(s, ssp)
	if i == -1 {
		return nil, errors.New("no price in input")
	}
	t = s[i+len(ssp):]
	j = strings.Index(t, ",")
	price := t[:j]

	// find address
	ssa := "<div data-name=\"Geo\">"
	i = strings.Index(s, ssa)
	if i == -1 {
		return nil, errors.New("no address in input")
	}
	t = s[i+len(ssa):]
	ssa = "<span itemprop=\"name\" content=\""
	j = strings.Index(t, ssa)
	t = t[j+len(ssa)+1:]
	j = strings.Index(t, "\">")
	address := t[:j]

	// find is inactive
	ssn := "Объявление снято с публикации"
	isactive := true
	if i = strings.Index(s, ssn); i != -1 {
		isactive = false
	}

	// find title
	sst := "<title>"
	i = strings.Index(s, sst)
	if i == -1 {
		return nil, errors.New("no title in input")
	}
	t = s[i+len(sst):]
	sst = "</title>"
	j = strings.Index(t, sst)
	title := t[:j]

	if len(price) == 0 {
		return nil, errors.New("price length is zero")
	}

	return &Entry{
		Provider: Cian,
		ID:       id,
		Price:    price,
		Address:  address,
		Title:    title,
		IsActive: isactive,
	}, nil
}
