package main

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("could not find the word you're looking for")

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]

	if !ok {
		return "", errNotFound
	}

	return result, nil
}
