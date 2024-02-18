// Code generated from Source.g4 by ANTLR 4.13.1. DO NOT EDIT.

package source // Source
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

type SourceParser struct {
	*antlr.BaseParser
}

var SourceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sourceParserInit() {
	staticData := &SourceParserStaticData
	staticData.LiteralNames = []string{
		"", "'DSN'", "'Connect-Timeout'", "'Try_Times'", "'Version'", "'Name'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'\\n'", "", "",
		"'{'", "'}'", "'('", "')'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "Version", "Method", "Route", "PreflightHeader",
		"Plugin", "Bool", "Int", "String", "Float", "Array", "MuOperator", "ArOperator",
		"JudgeOperator", "NewLine", "WS", "LINE_COMMENT", "LeftBrace", "RightBrace",
		"LeftParenthesis", "RightParenthesis", "Equal", "ID", "Var",
	}
	staticData.RuleNames = []string{
		"whole", "dsn", "timeout", "try", "version", "name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 56, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 0, 3, 0, 20,
		8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9,
		1, 1, 2, 1, 2, 1, 2, 5, 2, 35, 8, 2, 10, 2, 12, 2, 38, 9, 2, 1, 3, 1, 3,
		1, 3, 5, 3, 43, 8, 3, 10, 3, 12, 3, 46, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 0, 54, 0, 12,
		1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8,
		47, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 13, 3, 8, 4, 0, 13, 14, 3, 10,
		5, 0, 14, 16, 3, 2, 1, 0, 15, 17, 3, 4, 2, 0, 16, 15, 1, 0, 0, 0, 16, 17,
		1, 0, 0, 0, 17, 19, 1, 0, 0, 0, 18, 20, 3, 6, 3, 0, 19, 18, 1, 0, 0, 0,
		19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0,
		0, 0, 23, 24, 5, 1, 0, 0, 24, 28, 5, 13, 0, 0, 25, 27, 5, 19, 0, 0, 26,
		25, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0,
		0, 29, 3, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5, 2, 0, 0, 32, 36, 5,
		12, 0, 0, 33, 35, 5, 19, 0, 0, 34, 33, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0,
		36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38, 36, 1, 0,
		0, 0, 39, 40, 5, 3, 0, 0, 40, 44, 5, 12, 0, 0, 41, 43, 5, 19, 0, 0, 42,
		41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0,
		0, 45, 7, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 4, 0, 0, 48, 49, 5,
		6, 0, 0, 49, 50, 5, 19, 0, 0, 50, 9, 1, 0, 0, 0, 51, 52, 5, 5, 0, 0, 52,
		53, 5, 27, 0, 0, 53, 54, 5, 19, 0, 0, 54, 11, 1, 0, 0, 0, 5, 16, 19, 28,
		36, 44,
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

// SourceParserInit initializes any static state used to implement SourceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSourceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SourceParserInit() {
	staticData := &SourceParserStaticData
	staticData.once.Do(sourceParserInit)
}

// NewSourceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSourceParser(input antlr.TokenStream) *SourceParser {
	SourceParserInit()
	this := new(SourceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SourceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Source.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

// SourceParser tokens.
const (
	SourceParserEOF              = antlr.TokenEOF
	SourceParserT__0             = 1
	SourceParserT__1             = 2
	SourceParserT__2             = 3
	SourceParserT__3             = 4
	SourceParserT__4             = 5
	SourceParserVersion          = 6
	SourceParserMethod           = 7
	SourceParserRoute            = 8
	SourceParserPreflightHeader  = 9
	SourceParserPlugin           = 10
	SourceParserBool             = 11
	SourceParserInt              = 12
	SourceParserString_          = 13
	SourceParserFloat            = 14
	SourceParserArray            = 15
	SourceParserMuOperator       = 16
	SourceParserArOperator       = 17
	SourceParserJudgeOperator    = 18
	SourceParserNewLine          = 19
	SourceParserWS               = 20
	SourceParserLINE_COMMENT     = 21
	SourceParserLeftBrace        = 22
	SourceParserRightBrace       = 23
	SourceParserLeftParenthesis  = 24
	SourceParserRightParenthesis = 25
	SourceParserEqual            = 26
	SourceParserID               = 27
	SourceParserVar              = 28
)

// SourceParser rules.
const (
	SourceParserRULE_whole   = 0
	SourceParserRULE_dsn     = 1
	SourceParserRULE_timeout = 2
	SourceParserRULE_try     = 3
	SourceParserRULE_version = 4
	SourceParserRULE_name    = 5
)

// IWholeContext is an interface to support dynamic dispatch.
type IWholeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	Name() INameContext
	Dsn() IDsnContext
	EOF() antlr.TerminalNode
	Timeout() ITimeoutContext
	Try() ITryContext

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
	p.RuleIndex = SourceParserRULE_whole
	return p
}

func InitEmptyWholeContext(p *WholeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_whole
}

func (*WholeContext) IsWholeContext() {}

func NewWholeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WholeContext {
	var p = new(WholeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_whole

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

func (s *WholeContext) Dsn() IDsnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDsnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDsnContext)
}

func (s *WholeContext) EOF() antlr.TerminalNode {
	return s.GetToken(SourceParserEOF, 0)
}

func (s *WholeContext) Timeout() ITimeoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeoutContext)
}

func (s *WholeContext) Try() ITryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryContext)
}

func (s *WholeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WholeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WholeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterWhole(s)
	}
}

func (s *WholeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitWhole(s)
	}
}

func (p *SourceParser) Whole() (localctx IWholeContext) {
	localctx = NewWholeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SourceParserRULE_whole)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Version()
	}
	{
		p.SetState(13)
		p.Name()
	}
	{
		p.SetState(14)
		p.Dsn()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SourceParserT__1 {
		{
			p.SetState(15)
			p.Timeout()
		}

	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SourceParserT__2 {
		{
			p.SetState(18)
			p.Try()
		}

	}
	{
		p.SetState(21)
		p.Match(SourceParserEOF)
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

// IDsnContext is an interface to support dynamic dispatch.
type IDsnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() antlr.TerminalNode
	AllNewLine() []antlr.TerminalNode
	NewLine(i int) antlr.TerminalNode

	// IsDsnContext differentiates from other interfaces.
	IsDsnContext()
}

type DsnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDsnContext() *DsnContext {
	var p = new(DsnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_dsn
	return p
}

func InitEmptyDsnContext(p *DsnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_dsn
}

func (*DsnContext) IsDsnContext() {}

func NewDsnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DsnContext {
	var p = new(DsnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_dsn

	return p
}

func (s *DsnContext) GetParser() antlr.Parser { return s.parser }

func (s *DsnContext) String_() antlr.TerminalNode {
	return s.GetToken(SourceParserString_, 0)
}

func (s *DsnContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(SourceParserNewLine)
}

func (s *DsnContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(SourceParserNewLine, i)
}

func (s *DsnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DsnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DsnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterDsn(s)
	}
}

func (s *DsnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitDsn(s)
	}
}

func (p *SourceParser) Dsn() (localctx IDsnContext) {
	localctx = NewDsnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SourceParserRULE_dsn)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(SourceParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Match(SourceParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SourceParserNewLine {
		{
			p.SetState(25)
			p.Match(SourceParserNewLine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ITimeoutContext is an interface to support dynamic dispatch.
type ITimeoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode
	AllNewLine() []antlr.TerminalNode
	NewLine(i int) antlr.TerminalNode

	// IsTimeoutContext differentiates from other interfaces.
	IsTimeoutContext()
}

type TimeoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeoutContext() *TimeoutContext {
	var p = new(TimeoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_timeout
	return p
}

func InitEmptyTimeoutContext(p *TimeoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_timeout
}

func (*TimeoutContext) IsTimeoutContext() {}

func NewTimeoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeoutContext {
	var p = new(TimeoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_timeout

	return p
}

func (s *TimeoutContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeoutContext) Int() antlr.TerminalNode {
	return s.GetToken(SourceParserInt, 0)
}

func (s *TimeoutContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(SourceParserNewLine)
}

func (s *TimeoutContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(SourceParserNewLine, i)
}

func (s *TimeoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterTimeout(s)
	}
}

func (s *TimeoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitTimeout(s)
	}
}

func (p *SourceParser) Timeout() (localctx ITimeoutContext) {
	localctx = NewTimeoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SourceParserRULE_timeout)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(SourceParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(SourceParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SourceParserNewLine {
		{
			p.SetState(33)
			p.Match(SourceParserNewLine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ITryContext is an interface to support dynamic dispatch.
type ITryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode
	AllNewLine() []antlr.TerminalNode
	NewLine(i int) antlr.TerminalNode

	// IsTryContext differentiates from other interfaces.
	IsTryContext()
}

type TryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryContext() *TryContext {
	var p = new(TryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_try
	return p
}

func InitEmptyTryContext(p *TryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_try
}

func (*TryContext) IsTryContext() {}

func NewTryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryContext {
	var p = new(TryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_try

	return p
}

func (s *TryContext) GetParser() antlr.Parser { return s.parser }

func (s *TryContext) Int() antlr.TerminalNode {
	return s.GetToken(SourceParserInt, 0)
}

func (s *TryContext) AllNewLine() []antlr.TerminalNode {
	return s.GetTokens(SourceParserNewLine)
}

func (s *TryContext) NewLine(i int) antlr.TerminalNode {
	return s.GetToken(SourceParserNewLine, i)
}

func (s *TryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterTry(s)
	}
}

func (s *TryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitTry(s)
	}
}

func (p *SourceParser) Try() (localctx ITryContext) {
	localctx = NewTryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SourceParserRULE_try)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(SourceParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(SourceParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SourceParserNewLine {
		{
			p.SetState(41)
			p.Match(SourceParserNewLine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = SourceParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) Version() antlr.TerminalNode {
	return s.GetToken(SourceParserVersion, 0)
}

func (s *VersionContext) NewLine() antlr.TerminalNode {
	return s.GetToken(SourceParserNewLine, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *SourceParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SourceParserRULE_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(SourceParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(SourceParserVersion)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SourceParserNewLine)
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
	p.RuleIndex = SourceParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SourceParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(SourceParserID, 0)
}

func (s *NameContext) NewLine() antlr.TerminalNode {
	return s.GetToken(SourceParserNewLine, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *SourceParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SourceParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(SourceParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(SourceParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(SourceParserNewLine)
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
