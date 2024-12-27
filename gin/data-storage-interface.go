package main

type dataLayer interface {
	get(int) *blog
	post(*blog) error
	edit(*blog) error
	delete(int) error
}
