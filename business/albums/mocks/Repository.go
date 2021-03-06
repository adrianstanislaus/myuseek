// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	albums "myuseek/business/albums"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, domain
func (_m *Repository) Add(ctx context.Context, domain albums.Domain) (albums.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 albums.Domain
	if rf, ok := ret.Get(0).(func(context.Context, albums.Domain) albums.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(albums.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, albums.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlbumById provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAlbumById(ctx context.Context, domain albums.Domain) (albums.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 albums.Domain
	if rf, ok := ret.Get(0).(func(context.Context, albums.Domain) albums.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(albums.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, albums.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlbums provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAlbums(ctx context.Context, domain albums.Domain) ([]albums.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []albums.Domain
	if rf, ok := ret.Get(0).(func(context.Context, albums.Domain) []albums.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]albums.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, albums.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
