package tree_sitter_sagemath_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/Eloitor/tree-sitter-sagemath"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sagemath.Language())
	if language == nil {
		t.Errorf("Error loading Sagemath grammar")
	}
}
