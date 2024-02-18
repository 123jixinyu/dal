// Code generated from Source.g4 by ANTLR 4.13.1. DO NOT EDIT.

package source

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SourceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SourceLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sourcelexerLexerInit() {
	staticData := &SourceLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "Version", "Method", "Route",
		"PreflightHeader", "Plugin", "Kind", "Bool", "Int", "String", "Float",
		"Array", "MuOperator", "ArOperator", "JudgeOperator", "NewLine", "WS",
		"LINE_COMMENT", "LeftBrace", "RightBrace", "LeftParenthesis", "RightParenthesis",
		"Equal", "ID", "Var",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 539, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 4, 5, 106, 8, 5, 11, 5, 12, 5, 107, 1, 5,
		1, 5, 4, 5, 112, 8, 5, 11, 5, 12, 5, 113, 3, 5, 116, 8, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 139, 8, 6, 1, 7, 1, 7,
		5, 7, 143, 8, 7, 10, 7, 12, 7, 146, 9, 7, 4, 7, 148, 8, 7, 11, 7, 12, 7,
		149, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 384, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 399, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 410,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 421, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 426, 8, 12, 10, 12, 12, 12,
		429, 9, 12, 3, 12, 431, 8, 12, 1, 13, 1, 13, 5, 13, 435, 8, 13, 10, 13,
		12, 13, 438, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 446,
		8, 14, 10, 14, 12, 14, 449, 9, 14, 1, 14, 4, 14, 452, 8, 14, 11, 14, 12,
		14, 453, 1, 14, 1, 14, 5, 14, 458, 8, 14, 10, 14, 12, 14, 461, 9, 14, 3,
		14, 463, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 470, 8, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 477, 8, 15, 5, 15, 479, 8, 15, 10,
		15, 12, 15, 482, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 494, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 506, 8, 21, 10, 21, 12, 21,
		509, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 4, 27, 524, 8, 27, 11, 27, 12, 27, 525,
		1, 28, 1, 28, 4, 28, 530, 8, 28, 11, 28, 12, 28, 531, 1, 28, 5, 28, 535,
		8, 28, 10, 28, 12, 28, 538, 9, 28, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 11, 25, 12, 27, 13, 29,
		14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47,
		23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 1, 0, 9, 1, 0, 48, 57, 2, 0,
		65, 90, 97, 122, 1, 0, 49, 57, 1, 0, 39, 39, 2, 0, 42, 42, 47, 47, 2, 0,
		43, 43, 45, 45, 3, 0, 9, 9, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 3, 0,
		46, 46, 65, 90, 97, 122, 578, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 63,
		1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 97, 1, 0, 0, 0, 11,
		102, 1, 0, 0, 0, 13, 138, 1, 0, 0, 0, 15, 147, 1, 0, 0, 0, 17, 151, 1,
		0, 0, 0, 19, 398, 1, 0, 0, 0, 21, 409, 1, 0, 0, 0, 23, 420, 1, 0, 0, 0,
		25, 430, 1, 0, 0, 0, 27, 432, 1, 0, 0, 0, 29, 462, 1, 0, 0, 0, 31, 464,
		1, 0, 0, 0, 33, 485, 1, 0, 0, 0, 35, 487, 1, 0, 0, 0, 37, 493, 1, 0, 0,
		0, 39, 495, 1, 0, 0, 0, 41, 497, 1, 0, 0, 0, 43, 501, 1, 0, 0, 0, 45, 512,
		1, 0, 0, 0, 47, 514, 1, 0, 0, 0, 49, 516, 1, 0, 0, 0, 51, 518, 1, 0, 0,
		0, 53, 520, 1, 0, 0, 0, 55, 523, 1, 0, 0, 0, 57, 527, 1, 0, 0, 0, 59, 60,
		5, 68, 0, 0, 60, 61, 5, 83, 0, 0, 61, 62, 5, 78, 0, 0, 62, 2, 1, 0, 0,
		0, 63, 64, 5, 67, 0, 0, 64, 65, 5, 111, 0, 0, 65, 66, 5, 110, 0, 0, 66,
		67, 5, 110, 0, 0, 67, 68, 5, 101, 0, 0, 68, 69, 5, 99, 0, 0, 69, 70, 5,
		116, 0, 0, 70, 71, 5, 45, 0, 0, 71, 72, 5, 84, 0, 0, 72, 73, 5, 105, 0,
		0, 73, 74, 5, 109, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 111, 0, 0, 76,
		77, 5, 117, 0, 0, 77, 78, 5, 116, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 84,
		0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 121, 0, 0, 82, 83, 5, 95, 0, 0,
		83, 84, 5, 84, 0, 0, 84, 85, 5, 105, 0, 0, 85, 86, 5, 109, 0, 0, 86, 87,
		5, 101, 0, 0, 87, 88, 5, 115, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 86, 0,
		0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 114, 0, 0, 92, 93, 5, 115, 0, 0, 93,
		94, 5, 105, 0, 0, 94, 95, 5, 111, 0, 0, 95, 96, 5, 110, 0, 0, 96, 8, 1,
		0, 0, 0, 97, 98, 5, 78, 0, 0, 98, 99, 5, 97, 0, 0, 99, 100, 5, 109, 0,
		0, 100, 101, 5, 101, 0, 0, 101, 10, 1, 0, 0, 0, 102, 103, 3, 21, 10, 0,
		103, 105, 5, 47, 0, 0, 104, 106, 7, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 115,
		1, 0, 0, 0, 109, 111, 9, 0, 0, 0, 110, 112, 7, 0, 0, 0, 111, 110, 1, 0,
		0, 0, 112, 113, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 116, 1, 0, 0, 0, 115, 109, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		12, 1, 0, 0, 0, 117, 118, 5, 71, 0, 0, 118, 119, 5, 69, 0, 0, 119, 139,
		5, 84, 0, 0, 120, 121, 5, 80, 0, 0, 121, 122, 5, 79, 0, 0, 122, 123, 5,
		83, 0, 0, 123, 139, 5, 84, 0, 0, 124, 125, 5, 80, 0, 0, 125, 126, 5, 85,
		0, 0, 126, 139, 5, 84, 0, 0, 127, 128, 5, 68, 0, 0, 128, 129, 5, 69, 0,
		0, 129, 130, 5, 76, 0, 0, 130, 131, 5, 69, 0, 0, 131, 132, 5, 84, 0, 0,
		132, 139, 5, 69, 0, 0, 133, 134, 5, 80, 0, 0, 134, 135, 5, 65, 0, 0, 135,
		136, 5, 84, 0, 0, 136, 137, 5, 67, 0, 0, 137, 139, 5, 72, 0, 0, 138, 117,
		1, 0, 0, 0, 138, 120, 1, 0, 0, 0, 138, 124, 1, 0, 0, 0, 138, 127, 1, 0,
		0, 0, 138, 133, 1, 0, 0, 0, 139, 14, 1, 0, 0, 0, 140, 144, 5, 47, 0, 0,
		141, 143, 7, 1, 0, 0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144,
		1, 0, 0, 0, 147, 140, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0,
		0, 0, 149, 150, 1, 0, 0, 0, 150, 16, 1, 0, 0, 0, 151, 383, 5, 39, 0, 0,
		152, 153, 5, 97, 0, 0, 153, 154, 5, 99, 0, 0, 154, 155, 5, 99, 0, 0, 155,
		156, 5, 101, 0, 0, 156, 157, 5, 115, 0, 0, 157, 158, 5, 115, 0, 0, 158,
		159, 5, 45, 0, 0, 159, 160, 5, 99, 0, 0, 160, 161, 5, 111, 0, 0, 161, 162,
		5, 110, 0, 0, 162, 163, 5, 116, 0, 0, 163, 164, 5, 114, 0, 0, 164, 165,
		5, 111, 0, 0, 165, 166, 5, 108, 0, 0, 166, 167, 5, 45, 0, 0, 167, 168,
		5, 97, 0, 0, 168, 169, 5, 108, 0, 0, 169, 170, 5, 108, 0, 0, 170, 171,
		5, 111, 0, 0, 171, 172, 5, 119, 0, 0, 172, 173, 5, 45, 0, 0, 173, 174,
		5, 99, 0, 0, 174, 175, 5, 114, 0, 0, 175, 176, 5, 101, 0, 0, 176, 177,
		5, 100, 0, 0, 177, 178, 5, 101, 0, 0, 178, 179, 5, 110, 0, 0, 179, 180,
		5, 116, 0, 0, 180, 181, 5, 105, 0, 0, 181, 182, 5, 97, 0, 0, 182, 183,
		5, 108, 0, 0, 183, 384, 5, 115, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186,
		5, 99, 0, 0, 186, 187, 5, 99, 0, 0, 187, 188, 5, 101, 0, 0, 188, 189, 5,
		115, 0, 0, 189, 190, 5, 115, 0, 0, 190, 191, 5, 45, 0, 0, 191, 192, 5,
		99, 0, 0, 192, 193, 5, 111, 0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5,
		116, 0, 0, 195, 196, 5, 114, 0, 0, 196, 197, 5, 111, 0, 0, 197, 198, 5,
		108, 0, 0, 198, 199, 5, 45, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201, 5, 108,
		0, 0, 201, 202, 5, 108, 0, 0, 202, 203, 5, 111, 0, 0, 203, 204, 5, 119,
		0, 0, 204, 205, 5, 45, 0, 0, 205, 206, 5, 104, 0, 0, 206, 207, 5, 101,
		0, 0, 207, 208, 5, 97, 0, 0, 208, 209, 5, 100, 0, 0, 209, 210, 5, 101,
		0, 0, 210, 211, 5, 114, 0, 0, 211, 384, 5, 115, 0, 0, 212, 213, 5, 97,
		0, 0, 213, 214, 5, 99, 0, 0, 214, 215, 5, 99, 0, 0, 215, 216, 5, 101, 0,
		0, 216, 217, 5, 115, 0, 0, 217, 218, 5, 115, 0, 0, 218, 219, 5, 45, 0,
		0, 219, 220, 5, 99, 0, 0, 220, 221, 5, 111, 0, 0, 221, 222, 5, 110, 0,
		0, 222, 223, 5, 116, 0, 0, 223, 224, 5, 114, 0, 0, 224, 225, 5, 111, 0,
		0, 225, 226, 5, 108, 0, 0, 226, 227, 5, 45, 0, 0, 227, 228, 5, 97, 0, 0,
		228, 229, 5, 108, 0, 0, 229, 230, 5, 108, 0, 0, 230, 231, 5, 111, 0, 0,
		231, 232, 5, 119, 0, 0, 232, 233, 5, 45, 0, 0, 233, 234, 5, 109, 0, 0,
		234, 235, 5, 101, 0, 0, 235, 236, 5, 116, 0, 0, 236, 237, 5, 104, 0, 0,
		237, 238, 5, 111, 0, 0, 238, 239, 5, 100, 0, 0, 239, 384, 5, 115, 0, 0,
		240, 241, 5, 97, 0, 0, 241, 242, 5, 99, 0, 0, 242, 243, 5, 99, 0, 0, 243,
		244, 5, 101, 0, 0, 244, 245, 5, 115, 0, 0, 245, 246, 5, 115, 0, 0, 246,
		247, 5, 45, 0, 0, 247, 248, 5, 99, 0, 0, 248, 249, 5, 111, 0, 0, 249, 250,
		5, 110, 0, 0, 250, 251, 5, 116, 0, 0, 251, 252, 5, 114, 0, 0, 252, 253,
		5, 111, 0, 0, 253, 254, 5, 108, 0, 0, 254, 255, 5, 45, 0, 0, 255, 256,
		5, 97, 0, 0, 256, 257, 5, 108, 0, 0, 257, 258, 5, 108, 0, 0, 258, 259,
		5, 111, 0, 0, 259, 260, 5, 119, 0, 0, 260, 261, 5, 45, 0, 0, 261, 262,
		5, 111, 0, 0, 262, 263, 5, 114, 0, 0, 263, 264, 5, 105, 0, 0, 264, 265,
		5, 103, 0, 0, 265, 266, 5, 105, 0, 0, 266, 384, 5, 110, 0, 0, 267, 268,
		5, 97, 0, 0, 268, 269, 5, 99, 0, 0, 269, 270, 5, 99, 0, 0, 270, 271, 5,
		101, 0, 0, 271, 272, 5, 115, 0, 0, 272, 273, 5, 115, 0, 0, 273, 274, 5,
		45, 0, 0, 274, 275, 5, 99, 0, 0, 275, 276, 5, 111, 0, 0, 276, 277, 5, 110,
		0, 0, 277, 278, 5, 116, 0, 0, 278, 279, 5, 114, 0, 0, 279, 280, 5, 111,
		0, 0, 280, 281, 5, 108, 0, 0, 281, 282, 5, 45, 0, 0, 282, 283, 5, 101,
		0, 0, 283, 284, 5, 120, 0, 0, 284, 285, 5, 112, 0, 0, 285, 286, 5, 111,
		0, 0, 286, 287, 5, 115, 0, 0, 287, 288, 5, 101, 0, 0, 288, 289, 5, 45,
		0, 0, 289, 290, 5, 104, 0, 0, 290, 291, 5, 101, 0, 0, 291, 292, 5, 97,
		0, 0, 292, 293, 5, 100, 0, 0, 293, 294, 5, 101, 0, 0, 294, 295, 5, 114,
		0, 0, 295, 384, 5, 115, 0, 0, 296, 297, 5, 97, 0, 0, 297, 298, 5, 99, 0,
		0, 298, 299, 5, 99, 0, 0, 299, 300, 5, 101, 0, 0, 300, 301, 5, 115, 0,
		0, 301, 302, 5, 115, 0, 0, 302, 303, 5, 45, 0, 0, 303, 304, 5, 99, 0, 0,
		304, 305, 5, 111, 0, 0, 305, 306, 5, 110, 0, 0, 306, 307, 5, 116, 0, 0,
		307, 308, 5, 114, 0, 0, 308, 309, 5, 111, 0, 0, 309, 310, 5, 108, 0, 0,
		310, 311, 5, 45, 0, 0, 311, 312, 5, 109, 0, 0, 312, 313, 5, 97, 0, 0, 313,
		314, 5, 120, 0, 0, 314, 315, 5, 45, 0, 0, 315, 316, 5, 97, 0, 0, 316, 317,
		5, 103, 0, 0, 317, 384, 5, 101, 0, 0, 318, 319, 5, 97, 0, 0, 319, 320,
		5, 99, 0, 0, 320, 321, 5, 99, 0, 0, 321, 322, 5, 101, 0, 0, 322, 323, 5,
		115, 0, 0, 323, 324, 5, 115, 0, 0, 324, 325, 5, 45, 0, 0, 325, 326, 5,
		99, 0, 0, 326, 327, 5, 111, 0, 0, 327, 328, 5, 110, 0, 0, 328, 329, 5,
		116, 0, 0, 329, 330, 5, 114, 0, 0, 330, 331, 5, 111, 0, 0, 331, 332, 5,
		108, 0, 0, 332, 333, 5, 45, 0, 0, 333, 334, 5, 114, 0, 0, 334, 335, 5,
		101, 0, 0, 335, 336, 5, 113, 0, 0, 336, 337, 5, 117, 0, 0, 337, 338, 5,
		101, 0, 0, 338, 339, 5, 115, 0, 0, 339, 340, 5, 116, 0, 0, 340, 341, 5,
		45, 0, 0, 341, 342, 5, 104, 0, 0, 342, 343, 5, 101, 0, 0, 343, 344, 5,
		97, 0, 0, 344, 345, 5, 100, 0, 0, 345, 346, 5, 101, 0, 0, 346, 347, 5,
		114, 0, 0, 347, 384, 5, 115, 0, 0, 348, 349, 5, 97, 0, 0, 349, 350, 5,
		99, 0, 0, 350, 351, 5, 99, 0, 0, 351, 352, 5, 101, 0, 0, 352, 353, 5, 115,
		0, 0, 353, 354, 5, 115, 0, 0, 354, 355, 5, 45, 0, 0, 355, 356, 5, 99, 0,
		0, 356, 357, 5, 111, 0, 0, 357, 358, 5, 110, 0, 0, 358, 359, 5, 116, 0,
		0, 359, 360, 5, 114, 0, 0, 360, 361, 5, 111, 0, 0, 361, 362, 5, 108, 0,
		0, 362, 363, 5, 45, 0, 0, 363, 364, 5, 114, 0, 0, 364, 365, 5, 101, 0,
		0, 365, 366, 5, 113, 0, 0, 366, 367, 5, 117, 0, 0, 367, 368, 5, 101, 0,
		0, 368, 369, 5, 115, 0, 0, 369, 370, 5, 116, 0, 0, 370, 371, 5, 45, 0,
		0, 371, 372, 5, 109, 0, 0, 372, 373, 5, 101, 0, 0, 373, 374, 5, 116, 0,
		0, 374, 375, 5, 104, 0, 0, 375, 376, 5, 111, 0, 0, 376, 384, 5, 100, 0,
		0, 377, 378, 5, 111, 0, 0, 378, 379, 5, 114, 0, 0, 379, 380, 5, 105, 0,
		0, 380, 381, 5, 103, 0, 0, 381, 382, 5, 105, 0, 0, 382, 384, 5, 110, 0,
		0, 383, 152, 1, 0, 0, 0, 383, 184, 1, 0, 0, 0, 383, 212, 1, 0, 0, 0, 383,
		240, 1, 0, 0, 0, 383, 267, 1, 0, 0, 0, 383, 296, 1, 0, 0, 0, 383, 318,
		1, 0, 0, 0, 383, 348, 1, 0, 0, 0, 383, 377, 1, 0, 0, 0, 384, 385, 1, 0,
		0, 0, 385, 386, 5, 39, 0, 0, 386, 18, 1, 0, 0, 0, 387, 388, 5, 98, 0, 0,
		388, 389, 5, 97, 0, 0, 389, 390, 5, 115, 0, 0, 390, 391, 5, 105, 0, 0,
		391, 399, 5, 99, 0, 0, 392, 393, 5, 106, 0, 0, 393, 394, 5, 119, 0, 0,
		394, 399, 5, 116, 0, 0, 395, 396, 5, 108, 0, 0, 396, 397, 5, 111, 0, 0,
		397, 399, 5, 103, 0, 0, 398, 387, 1, 0, 0, 0, 398, 392, 1, 0, 0, 0, 398,
		395, 1, 0, 0, 0, 399, 20, 1, 0, 0, 0, 400, 401, 5, 97, 0, 0, 401, 402,
		5, 112, 0, 0, 402, 410, 5, 105, 0, 0, 403, 404, 5, 115, 0, 0, 404, 405,
		5, 111, 0, 0, 405, 406, 5, 117, 0, 0, 406, 407, 5, 114, 0, 0, 407, 408,
		5, 99, 0, 0, 408, 410, 5, 101, 0, 0, 409, 400, 1, 0, 0, 0, 409, 403, 1,
		0, 0, 0, 410, 22, 1, 0, 0, 0, 411, 412, 5, 116, 0, 0, 412, 413, 5, 114,
		0, 0, 413, 414, 5, 117, 0, 0, 414, 421, 5, 101, 0, 0, 415, 416, 5, 102,
		0, 0, 416, 417, 5, 97, 0, 0, 417, 418, 5, 108, 0, 0, 418, 419, 5, 115,
		0, 0, 419, 421, 5, 101, 0, 0, 420, 411, 1, 0, 0, 0, 420, 415, 1, 0, 0,
		0, 421, 24, 1, 0, 0, 0, 422, 431, 5, 48, 0, 0, 423, 427, 7, 2, 0, 0, 424,
		426, 7, 0, 0, 0, 425, 424, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 425,
		1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 431, 1, 0, 0, 0, 429, 427, 1, 0,
		0, 0, 430, 422, 1, 0, 0, 0, 430, 423, 1, 0, 0, 0, 431, 26, 1, 0, 0, 0,
		432, 436, 5, 39, 0, 0, 433, 435, 8, 3, 0, 0, 434, 433, 1, 0, 0, 0, 435,
		438, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 439,
		1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 439, 440, 5, 39, 0, 0, 440, 28, 1, 0,
		0, 0, 441, 442, 5, 48, 0, 0, 442, 443, 5, 46, 0, 0, 443, 447, 1, 0, 0,
		0, 444, 446, 7, 0, 0, 0, 445, 444, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447,
		445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 463, 1, 0, 0, 0, 449, 447,
		1, 0, 0, 0, 450, 452, 7, 2, 0, 0, 451, 450, 1, 0, 0, 0, 452, 453, 1, 0,
		0, 0, 453, 451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0,
		455, 459, 5, 46, 0, 0, 456, 458, 7, 0, 0, 0, 457, 456, 1, 0, 0, 0, 458,
		461, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 463,
		1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 462, 441, 1, 0, 0, 0, 462, 451, 1, 0,
		0, 0, 463, 30, 1, 0, 0, 0, 464, 469, 5, 91, 0, 0, 465, 470, 3, 23, 11,
		0, 466, 470, 3, 25, 12, 0, 467, 470, 3, 27, 13, 0, 468, 470, 3, 29, 14,
		0, 469, 465, 1, 0, 0, 0, 469, 466, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469,
		468, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 480, 1, 0, 0, 0, 471, 476,
		5, 44, 0, 0, 472, 477, 3, 23, 11, 0, 473, 477, 3, 25, 12, 0, 474, 477,
		3, 27, 13, 0, 475, 477, 3, 29, 14, 0, 476, 472, 1, 0, 0, 0, 476, 473, 1,
		0, 0, 0, 476, 474, 1, 0, 0, 0, 476, 475, 1, 0, 0, 0, 477, 479, 1, 0, 0,
		0, 478, 471, 1, 0, 0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480,
		481, 1, 0, 0, 0, 481, 483, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 483, 484,
		5, 93, 0, 0, 484, 32, 1, 0, 0, 0, 485, 486, 7, 4, 0, 0, 486, 34, 1, 0,
		0, 0, 487, 488, 7, 5, 0, 0, 488, 36, 1, 0, 0, 0, 489, 490, 5, 61, 0, 0,
		490, 494, 5, 61, 0, 0, 491, 492, 5, 33, 0, 0, 492, 494, 5, 61, 0, 0, 493,
		489, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 494, 38, 1, 0, 0, 0, 495, 496, 5,
		10, 0, 0, 496, 40, 1, 0, 0, 0, 497, 498, 7, 6, 0, 0, 498, 499, 1, 0, 0,
		0, 499, 500, 6, 20, 0, 0, 500, 42, 1, 0, 0, 0, 501, 502, 5, 47, 0, 0, 502,
		503, 5, 47, 0, 0, 503, 507, 1, 0, 0, 0, 504, 506, 8, 7, 0, 0, 505, 504,
		1, 0, 0, 0, 506, 509, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 507, 508, 1, 0,
		0, 0, 508, 510, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 510, 511, 6, 21, 0, 0,
		511, 44, 1, 0, 0, 0, 512, 513, 5, 123, 0, 0, 513, 46, 1, 0, 0, 0, 514,
		515, 5, 125, 0, 0, 515, 48, 1, 0, 0, 0, 516, 517, 5, 40, 0, 0, 517, 50,
		1, 0, 0, 0, 518, 519, 5, 41, 0, 0, 519, 52, 1, 0, 0, 0, 520, 521, 5, 61,
		0, 0, 521, 54, 1, 0, 0, 0, 522, 524, 7, 1, 0, 0, 523, 522, 1, 0, 0, 0,
		524, 525, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 525, 526, 1, 0, 0, 0, 526,
		56, 1, 0, 0, 0, 527, 529, 5, 36, 0, 0, 528, 530, 7, 1, 0, 0, 529, 528,
		1, 0, 0, 0, 530, 531, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0,
		0, 0, 532, 536, 1, 0, 0, 0, 533, 535, 7, 8, 0, 0, 534, 533, 1, 0, 0, 0,
		535, 538, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537,
		58, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0, 26, 0, 107, 113, 115, 138, 144, 149,
		383, 398, 409, 420, 427, 430, 436, 447, 453, 459, 462, 469, 476, 480, 493,
		507, 525, 531, 536, 1, 6, 0, 0,
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

// SourceLexerInit initializes any static state used to implement SourceLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSourceLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SourceLexerInit() {
	staticData := &SourceLexerLexerStaticData
	staticData.once.Do(sourcelexerLexerInit)
}

// NewSourceLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSourceLexer(input antlr.CharStream) *SourceLexer {
	SourceLexerInit()
	l := new(SourceLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SourceLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Source.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SourceLexer tokens.
const (
	SourceLexerT__0             = 1
	SourceLexerT__1             = 2
	SourceLexerT__2             = 3
	SourceLexerT__3             = 4
	SourceLexerT__4             = 5
	SourceLexerVersion          = 6
	SourceLexerMethod           = 7
	SourceLexerRoute            = 8
	SourceLexerPreflightHeader  = 9
	SourceLexerPlugin           = 10
	SourceLexerBool             = 11
	SourceLexerInt              = 12
	SourceLexerString_          = 13
	SourceLexerFloat            = 14
	SourceLexerArray            = 15
	SourceLexerMuOperator       = 16
	SourceLexerArOperator       = 17
	SourceLexerJudgeOperator    = 18
	SourceLexerNewLine          = 19
	SourceLexerWS               = 20
	SourceLexerLINE_COMMENT     = 21
	SourceLexerLeftBrace        = 22
	SourceLexerRightBrace       = 23
	SourceLexerLeftParenthesis  = 24
	SourceLexerRightParenthesis = 25
	SourceLexerEqual            = 26
	SourceLexerID               = 27
	SourceLexerVar              = 28
)