// Code generated from Source.g4 by ANTLR 4.13.1. DO NOT EDIT.

package source // Source
import "github.com/antlr4-go/antlr/v4"

// BaseSourceListener is a complete listener for a parse tree produced by SourceParser.
type BaseSourceListener struct{}

var _ SourceListener = &BaseSourceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSourceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSourceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSourceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSourceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWhole is called when production whole is entered.
func (s *BaseSourceListener) EnterWhole(ctx *WholeContext) {}

// ExitWhole is called when production whole is exited.
func (s *BaseSourceListener) ExitWhole(ctx *WholeContext) {}

// EnterDsn is called when production dsn is entered.
func (s *BaseSourceListener) EnterDsn(ctx *DsnContext) {}

// ExitDsn is called when production dsn is exited.
func (s *BaseSourceListener) ExitDsn(ctx *DsnContext) {}

// EnterTimeout is called when production timeout is entered.
func (s *BaseSourceListener) EnterTimeout(ctx *TimeoutContext) {}

// ExitTimeout is called when production timeout is exited.
func (s *BaseSourceListener) ExitTimeout(ctx *TimeoutContext) {}

// EnterTry is called when production try is entered.
func (s *BaseSourceListener) EnterTry(ctx *TryContext) {}

// ExitTry is called when production try is exited.
func (s *BaseSourceListener) ExitTry(ctx *TryContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseSourceListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseSourceListener) ExitVersion(ctx *VersionContext) {}

// EnterName is called when production name is entered.
func (s *BaseSourceListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSourceListener) ExitName(ctx *NameContext) {}
