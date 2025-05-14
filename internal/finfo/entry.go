package finfo

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type ProviderType int

const (
	Cian ProviderType = iota
	Sber
)

func (pt ProviderType) String() string {
	return [...]string{"Циан", "Сбер"}[pt]
}

func (pt ProviderType) StringLink() string {
	return [...]string{"cian.ru", "<UNK>"}[pt]
}

type Entry struct {
	Provider ProviderType
	ID       int
	Price    string
	Address  string
	Title    string
	IsActive bool
}

func NewEntry(link string) (*Entry, error) {
	var pt ProviderType = -1

	if strings.Contains(link, Cian.StringLink()) {
		pt = Cian
	}

	page, err := getPage(link)
	if err != nil {
		return nil, err
	}

	var entry *Entry
	switch pt {
	case Cian:
		entry, err = parseCian(page)
	case Sber:
		log.Fatal("TODO: Sber not implemented yet")
	default:
		return nil, errors.New("unknown provider type")
	}

	return entry, nil
}

func getPage(link string) (string, error) {
	r, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
