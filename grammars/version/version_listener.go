// Code generated from Version.g4 by ANTLR 4.13.1. DO NOT EDIT.

package version // Version
import "github.com/antlr4-go/antlr/v4"

// VersionListener is a complete listener for a parse tree produced by VersionParser.
type VersionListener interface {
	antlr.ParseTreeListener

	// EnterWhole is called when entering the whole production.
	EnterWhole(c *WholeContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitWhole is called when exiting the whole production.
	ExitWhole(c *WholeContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
