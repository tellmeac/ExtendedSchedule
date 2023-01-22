// Code generated by ent, DO NOT EDIT.

package ent

import (
	"tellmeac/extended-schedule/adapters/ent/schema"
	entuserconfig "tellmeac/extended-schedule/adapters/ent/userconfig"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entuserconfigFields := schema.UserConfig{}.Fields()
	_ = entuserconfigFields
	// entuserconfigDescEmail is the schema descriptor for Email field.
	entuserconfigDescEmail := entuserconfigFields[1].Descriptor()
	// entuserconfig.EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	entuserconfig.EmailValidator = entuserconfigDescEmail.Validators[0].(func(string) error)
	// entuserconfigDescID is the schema descriptor for id field.
	entuserconfigDescID := entuserconfigFields[0].Descriptor()
	// entuserconfig.DefaultID holds the default value on creation for the id field.
	entuserconfig.DefaultID = entuserconfigDescID.Default.(func() uuid.UUID)
}
