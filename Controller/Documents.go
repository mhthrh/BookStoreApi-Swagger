package Controller

import "github.com/mhthrh/ApiStore/Model/Book"

// Package classification of book API
//
// Documentation for book API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.3
//
//	Consumes:
//	- application/json
//
//	Controller:
//	- application/json
//
// swagger:meta

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the Controller

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// ValidationUtil errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of Controller
// swagger:response booksResponse
type booksResponseWrapper struct {
	// All current Controller
	// in: body
	Body []Book.Book
}

// Data structure representing a single book
// swagger:response bookResponse
type bookResponseWrapper struct {
	// Newly created book
	// in: body
	Body Book.Book
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateBook createBook
type bookParamsWrapper struct {
	// book data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body Book.Book
}

// swagger:parameters updateBook
type bookIDParamsWrapper struct {
	// The id of the book for which the operation relates
	// in: path
	// required: true
	Id int `json:"id"`
}
