// Code generated from Version.g4 by ANTLR 4.13.1. DO NOT EDIT.

package version // Version
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VersionParser struct {
	*antlr.BaseParser
}

var VersionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func versionParserInit() {
	staticData := &VersionParserStaticData
	staticData.LiteralNames = []string{
		"", "'Version'", "'Name'", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "'\\n'", "", "", "'{'", "'}'", "'('", "')'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "Version", "Method", "Route", "PreflightHeader", "Plugin",
		"Bool", "Int", "String", "Float", "Array", "MuOperator", "ArOperator",
		"JudgeOperator", "NewLine", "WS", "LINE_COMMENT", "LeftBrace", "RightBrace",
		"LeftParenthesis", "RightParenthesis", "Equal", "ID", "Var",
	}
	staticData.RuleNames = []string{
		"whole", "version", "name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 18, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0,
		0, 14, 0, 6, 1, 0, 0, 0, 2, 9, 1, 0, 0, 0, 4, 13, 1, 0, 0, 0, 6, 7, 3,
		2, 1, 0, 7, 8, 3, 4, 2, 0, 8, 1, 1, 0, 0, 0, 9, 10, 5, 1, 0, 0, 10, 11,
		5, 3, 0, 0, 11, 12, 5, 16, 0, 0, 12, 3, 1, 0, 0, 0, 13, 14, 5, 2, 0, 0,
		14, 15, 5, 24, 0, 0, 15, 16, 5, 16, 0, 0, 16, 5, 1, 0, 0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VersionParserInit initializes any static state used to implement VersionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVersionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VersionParserInit() {
	staticData := &VersionParserStaticData
	staticData.once.Do(versionParserInit)
}

// NewVersionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVersionParser(input antlr.TokenStream) *VersionParser {
	VersionParserInit()
	this := new(VersionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VersionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Version.g4"

	return this
}

// VersionParser tokens.
const (
	VersionParserEOF              = antlr.TokenEOF
	VersionParserT__0             = 1
	VersionParserT__1             = 2
	VersionParserVersion          = 3
	VersionParserMethod           = 4
	VersionParserRoute            = 5
	VersionParserPreflightHeader  = 6
	VersionParserPlugin           = 7
	VersionParserBool             = 8
	VersionParserInt              = 9
	VersionParserString_          = 10
	VersionParserFloat            = 11
	VersionParserArray            = 12
	VersionParserMuOperator       = 13
	VersionParserArOperator       = 14
	VersionParserJudgeOperator    = 15
	VersionParserNewLine          = 16
	VersionParserWS               = 17
	VersionParserLINE_COMMENT     = 18
	VersionParserLeftBrace        = 19
	VersionParserRightBrace       = 20
	VersionParserLeftParenthesis  = 21
	VersionParserRightParenthesis = 22
	VersionParserEqual            = 23
	VersionParserID               = 24
	VersionParserVar              = 25
)

// VersionParser rules.
const (
	VersionParserRULE_whole   = 0
	VersionParserRULE_version = 1
	VersionParserRULE_name    = 2
)

// IWholeContext is an interface to support dynamic dispatch.
type IWholeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	Name() INameContext

	// IsWholeContext differentiates from other interfaces.
	IsWholeContext()
}

type WholeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWholeContext() *WholeContext {
	var p = new(WholeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_whole
	return p
}

func InitEmptyWholeContext(p *WholeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_whole
}

func (*WholeContext) IsWholeContext() {}

func NewWholeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WholeContext {
	var p = new(WholeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_whole

	return p
}

func (s *WholeContext) GetParser() antlr.Parser { return s.parser }

func (s *WholeContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *WholeContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *WholeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WholeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WholeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterWhole(s)
	}
}

func (s *WholeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitWhole(s)
	}
}

func (p *VersionParser) Whole() (localctx IWholeContext) {
	localctx = NewWholeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VersionParserRULE_whole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.Version()
	}
	{
		p.SetState(7)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() antlr.TerminalNode
	NewLine() antlr.TerminalNode

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) Version() antlr.TerminalNode {
	return s.GetToken(VersionParserVersion, 0)
}

func (s *VersionContext) NewLine() antlr.TerminalNode {
	return s.GetToken(VersionParserNewLine, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *VersionParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VersionParserRULE_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(9)
		p.Match(VersionParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(10)
		p.Match(VersionParserVersion)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(11)
		p.Match(VersionParserNewLine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	NewLine() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VersionParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VersionParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(VersionParserID, 0)
}

func (s *NameContext) NewLine() antlr.TerminalNode {
	return s.GetToken(VersionParserNewLine, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VersionListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *VersionParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VersionParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(VersionParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(14)
		p.Match(VersionParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(15)
		p.Match(VersionParserNewLine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
