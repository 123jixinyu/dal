// Code generated from Version.g4 by ANTLR 4.13.1. DO NOT EDIT.

package version // Version
import "github.com/antlr4-go/antlr/v4"

// BaseVersionListener is a complete listener for a parse tree produced by VersionParser.
type BaseVersionListener struct{}

var _ VersionListener = &BaseVersionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVersionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVersionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVersionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVersionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWhole is called when production whole is entered.
func (s *BaseVersionListener) EnterWhole(ctx *WholeContext) {}

// ExitWhole is called when production whole is exited.
func (s *BaseVersionListener) ExitWhole(ctx *WholeContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseVersionListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseVersionListener) ExitVersion(ctx *VersionContext) {}

// EnterName is called when production name is entered.
func (s *BaseVersionListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseVersionListener) ExitName(ctx *NameContext) {}
