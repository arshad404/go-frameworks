package main

import "errors"

type dataStorage struct {
	blogStorage map[int]*blog
}

func NewDataStorage() dataStorage {
	return dataStorage{
		blogStorage: make(map[int]*blog),
	}
}

func (ds dataStorage) get(id int) *blog {
	if blog, exists := ds.blogStorage[id]; exists {
		return blog
	}
	return nil // Return nil if the blog does not exist
}

func (ds dataStorage) post(blog *blog) error {
	if ds.blogStorage[blog.ID] != nil {
		return errors.New("blog is already present")
	}
	ds.blogStorage[blog.ID] = blog
	return nil
}

func (ds dataStorage) edit(blog blog) error {
	if ds.blogStorage[blog.ID] == nil {
		return errors.New("blog with this id is not present in the database, please create a new blog")
	}
	ds.blogStorage[blog.ID] = &blog
	return nil
}

func (ds dataStorage) delete(id int) error {
	if ds.blogStorage[id] == nil {
		return errors.New("blog with this id is not present in the database")
	}
	delete(ds.blogStorage, id)
	return nil
}
