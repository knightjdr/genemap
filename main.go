// Package main contains functions for generating a map of gene identifiers and converting between them.
package main

import (
	"github.com/knightjdr/genemap/internal/generate"
	"github.com/knightjdr/genemap/internal/terms"
)

// CreateMapper creates a structure for mapping between geen identifiers.
var CreateMapper = terms.CreateMapper

// Generate mapping file for mapping between gene identifiers.
var Generate = generate.MappingFiles

// Mapper structure for setting options for mapping and retrieving results.
type Mapper terms.Mapper
