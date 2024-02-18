// Code generated from Source.g4 by ANTLR 4.13.1. DO NOT EDIT.

package source // Source
import "github.com/antlr4-go/antlr/v4"

// SourceListener is a complete listener for a parse tree produced by SourceParser.
type SourceListener interface {
	antlr.ParseTreeListener

	// EnterWhole is called when entering the whole production.
	EnterWhole(c *WholeContext)

	// EnterDsn is called when entering the dsn production.
	EnterDsn(c *DsnContext)

	// EnterTimeout is called when entering the timeout production.
	EnterTimeout(c *TimeoutContext)

	// EnterTry is called when entering the try production.
	EnterTry(c *TryContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitWhole is called when exiting the whole production.
	ExitWhole(c *WholeContext)

	// ExitDsn is called when exiting the dsn production.
	ExitDsn(c *DsnContext)

	// ExitTimeout is called when exiting the timeout production.
	ExitTimeout(c *TimeoutContext)

	// ExitTry is called when exiting the try production.
	ExitTry(c *TryContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
