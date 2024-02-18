// Generated from /Users/jixinyu/Documents/Code/Antlr/data-api/grammars/Common.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class CommonLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, Version=3, Method=4, Route=5, PreflightHeader=6, Plugin=7, 
		Bool=8, Int=9, String=10, Float=11, Array=12, MuOperator=13, ArOperator=14, 
		JudgeOperator=15, NewLine=16, WS=17, LINE_COMMENT=18, LeftBrace=19, RightBrace=20, 
		LeftParenthesis=21, RightParenthesis=22, Equal=23, ID=24, Var=25;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "Version", "Method", "Route", "PreflightHeader", "Plugin", 
			"Kind", "Bool", "Int", "String", "Float", "Array", "MuOperator", "ArOperator", 
			"JudgeOperator", "NewLine", "WS", "LINE_COMMENT", "LeftBrace", "RightBrace", 
			"LeftParenthesis", "RightParenthesis", "Equal", "ID", "Var"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Version'", "'Name'", null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "'\\n'", null, null, "'{'", "'}'", 
			"'('", "')'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "Version", "Method", "Route", "PreflightHeader", "Plugin", 
			"Bool", "Int", "String", "Float", "Array", "MuOperator", "ArOperator", 
			"JudgeOperator", "NewLine", "WS", "LINE_COMMENT", "LeftBrace", "RightBrace", 
			"LeftParenthesis", "RightParenthesis", "Equal", "ID", "Var"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public CommonLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Common.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0019\u01f7\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0004\u0002F\b\u0002\u000b\u0002\f\u0002G\u0001"+
		"\u0002\u0001\u0002\u0004\u0002L\b\u0002\u000b\u0002\f\u0002M\u0003\u0002"+
		"P\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003g\b\u0003"+
		"\u0001\u0004\u0001\u0004\u0005\u0004k\b\u0004\n\u0004\f\u0004n\t\u0004"+
		"\u0004\u0004p\b\u0004\u000b\u0004\f\u0004q\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0003\u0005\u015c\b\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u016b"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u0176\b\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003"+
		"\b\u0181\b\b\u0001\t\u0001\t\u0001\t\u0005\t\u0186\b\t\n\t\f\t\u0189\t"+
		"\t\u0003\t\u018b\b\t\u0001\n\u0001\n\u0005\n\u018f\b\n\n\n\f\n\u0192\t"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005"+
		"\u000b\u019a\b\u000b\n\u000b\f\u000b\u019d\t\u000b\u0001\u000b\u0004\u000b"+
		"\u01a0\b\u000b\u000b\u000b\f\u000b\u01a1\u0001\u000b\u0001\u000b\u0005"+
		"\u000b\u01a6\b\u000b\n\u000b\f\u000b\u01a9\t\u000b\u0003\u000b\u01ab\b"+
		"\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u01b2\b\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u01b9\b\f\u0005\f\u01bb\b\f"+
		"\n\f\f\f\u01be\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u01ca\b\u000f"+
		"\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u01d6\b\u0012"+
		"\n\u0012\f\u0012\u01d9\t\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0016\u0001"+
		"\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0004\u0018\u01e8\b\u0018\u000b"+
		"\u0018\f\u0018\u01e9\u0001\u0019\u0001\u0019\u0004\u0019\u01ee\b\u0019"+
		"\u000b\u0019\f\u0019\u01ef\u0001\u0019\u0005\u0019\u01f3\b\u0019\n\u0019"+
		"\f\u0019\u01f6\t\u0019\u0000\u0000\u001a\u0001\u0001\u0003\u0002\u0005"+
		"\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\u0000\u0011\b\u0013"+
		"\t\u0015\n\u0017\u000b\u0019\f\u001b\r\u001d\u000e\u001f\u000f!\u0010"+
		"#\u0011%\u0012\'\u0013)\u0014+\u0015-\u0016/\u00171\u00183\u0019\u0001"+
		"\u0000\t\u0001\u000009\u0002\u0000AZaz\u0001\u000019\u0001\u0000\'\'\u0002"+
		"\u0000**//\u0002\u0000++--\u0003\u0000\t\t\r\r  \u0002\u0000\n\n\r\r\u0003"+
		"\u0000..AZaz\u021e\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001"+
		"\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001"+
		"\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000"+
		"\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000"+
		"\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000"+
		"\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000"+
		"\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000"+
		"\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'"+
		"\u0001\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000"+
		"\u0000\u0000\u0000-\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000"+
		"\u00001\u0001\u0000\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00015"+
		"\u0001\u0000\u0000\u0000\u0003=\u0001\u0000\u0000\u0000\u0005B\u0001\u0000"+
		"\u0000\u0000\u0007f\u0001\u0000\u0000\u0000\to\u0001\u0000\u0000\u0000"+
		"\u000bs\u0001\u0000\u0000\u0000\r\u016a\u0001\u0000\u0000\u0000\u000f"+
		"\u0175\u0001\u0000\u0000\u0000\u0011\u0180\u0001\u0000\u0000\u0000\u0013"+
		"\u018a\u0001\u0000\u0000\u0000\u0015\u018c\u0001\u0000\u0000\u0000\u0017"+
		"\u01aa\u0001\u0000\u0000\u0000\u0019\u01ac\u0001\u0000\u0000\u0000\u001b"+
		"\u01c1\u0001\u0000\u0000\u0000\u001d\u01c3\u0001\u0000\u0000\u0000\u001f"+
		"\u01c9\u0001\u0000\u0000\u0000!\u01cb\u0001\u0000\u0000\u0000#\u01cd\u0001"+
		"\u0000\u0000\u0000%\u01d1\u0001\u0000\u0000\u0000\'\u01dc\u0001\u0000"+
		"\u0000\u0000)\u01de\u0001\u0000\u0000\u0000+\u01e0\u0001\u0000\u0000\u0000"+
		"-\u01e2\u0001\u0000\u0000\u0000/\u01e4\u0001\u0000\u0000\u00001\u01e7"+
		"\u0001\u0000\u0000\u00003\u01eb\u0001\u0000\u0000\u000056\u0005V\u0000"+
		"\u000067\u0005e\u0000\u000078\u0005r\u0000\u000089\u0005s\u0000\u0000"+
		"9:\u0005i\u0000\u0000:;\u0005o\u0000\u0000;<\u0005n\u0000\u0000<\u0002"+
		"\u0001\u0000\u0000\u0000=>\u0005N\u0000\u0000>?\u0005a\u0000\u0000?@\u0005"+
		"m\u0000\u0000@A\u0005e\u0000\u0000A\u0004\u0001\u0000\u0000\u0000BC\u0003"+
		"\u000f\u0007\u0000CE\u0005/\u0000\u0000DF\u0007\u0000\u0000\u0000ED\u0001"+
		"\u0000\u0000\u0000FG\u0001\u0000\u0000\u0000GE\u0001\u0000\u0000\u0000"+
		"GH\u0001\u0000\u0000\u0000HO\u0001\u0000\u0000\u0000IK\t\u0000\u0000\u0000"+
		"JL\u0007\u0000\u0000\u0000KJ\u0001\u0000\u0000\u0000LM\u0001\u0000\u0000"+
		"\u0000MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000NP\u0001\u0000"+
		"\u0000\u0000OI\u0001\u0000\u0000\u0000OP\u0001\u0000\u0000\u0000P\u0006"+
		"\u0001\u0000\u0000\u0000QR\u0005G\u0000\u0000RS\u0005E\u0000\u0000Sg\u0005"+
		"T\u0000\u0000TU\u0005P\u0000\u0000UV\u0005O\u0000\u0000VW\u0005S\u0000"+
		"\u0000Wg\u0005T\u0000\u0000XY\u0005P\u0000\u0000YZ\u0005U\u0000\u0000"+
		"Zg\u0005T\u0000\u0000[\\\u0005D\u0000\u0000\\]\u0005E\u0000\u0000]^\u0005"+
		"L\u0000\u0000^_\u0005E\u0000\u0000_`\u0005T\u0000\u0000`g\u0005E\u0000"+
		"\u0000ab\u0005P\u0000\u0000bc\u0005A\u0000\u0000cd\u0005T\u0000\u0000"+
		"de\u0005C\u0000\u0000eg\u0005H\u0000\u0000fQ\u0001\u0000\u0000\u0000f"+
		"T\u0001\u0000\u0000\u0000fX\u0001\u0000\u0000\u0000f[\u0001\u0000\u0000"+
		"\u0000fa\u0001\u0000\u0000\u0000g\b\u0001\u0000\u0000\u0000hl\u0005/\u0000"+
		"\u0000ik\u0007\u0001\u0000\u0000ji\u0001\u0000\u0000\u0000kn\u0001\u0000"+
		"\u0000\u0000lj\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000mp\u0001"+
		"\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000oh\u0001\u0000\u0000\u0000"+
		"pq\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000"+
		"\u0000r\n\u0001\u0000\u0000\u0000s\u015b\u0005\'\u0000\u0000tu\u0005a"+
		"\u0000\u0000uv\u0005c\u0000\u0000vw\u0005c\u0000\u0000wx\u0005e\u0000"+
		"\u0000xy\u0005s\u0000\u0000yz\u0005s\u0000\u0000z{\u0005-\u0000\u0000"+
		"{|\u0005c\u0000\u0000|}\u0005o\u0000\u0000}~\u0005n\u0000\u0000~\u007f"+
		"\u0005t\u0000\u0000\u007f\u0080\u0005r\u0000\u0000\u0080\u0081\u0005o"+
		"\u0000\u0000\u0081\u0082\u0005l\u0000\u0000\u0082\u0083\u0005-\u0000\u0000"+
		"\u0083\u0084\u0005a\u0000\u0000\u0084\u0085\u0005l\u0000\u0000\u0085\u0086"+
		"\u0005l\u0000\u0000\u0086\u0087\u0005o\u0000\u0000\u0087\u0088\u0005w"+
		"\u0000\u0000\u0088\u0089\u0005-\u0000\u0000\u0089\u008a\u0005c\u0000\u0000"+
		"\u008a\u008b\u0005r\u0000\u0000\u008b\u008c\u0005e\u0000\u0000\u008c\u008d"+
		"\u0005d\u0000\u0000\u008d\u008e\u0005e\u0000\u0000\u008e\u008f\u0005n"+
		"\u0000\u0000\u008f\u0090\u0005t\u0000\u0000\u0090\u0091\u0005i\u0000\u0000"+
		"\u0091\u0092\u0005a\u0000\u0000\u0092\u0093\u0005l\u0000\u0000\u0093\u015c"+
		"\u0005s\u0000\u0000\u0094\u0095\u0005a\u0000\u0000\u0095\u0096\u0005c"+
		"\u0000\u0000\u0096\u0097\u0005c\u0000\u0000\u0097\u0098\u0005e\u0000\u0000"+
		"\u0098\u0099\u0005s\u0000\u0000\u0099\u009a\u0005s\u0000\u0000\u009a\u009b"+
		"\u0005-\u0000\u0000\u009b\u009c\u0005c\u0000\u0000\u009c\u009d\u0005o"+
		"\u0000\u0000\u009d\u009e\u0005n\u0000\u0000\u009e\u009f\u0005t\u0000\u0000"+
		"\u009f\u00a0\u0005r\u0000\u0000\u00a0\u00a1\u0005o\u0000\u0000\u00a1\u00a2"+
		"\u0005l\u0000\u0000\u00a2\u00a3\u0005-\u0000\u0000\u00a3\u00a4\u0005a"+
		"\u0000\u0000\u00a4\u00a5\u0005l\u0000\u0000\u00a5\u00a6\u0005l\u0000\u0000"+
		"\u00a6\u00a7\u0005o\u0000\u0000\u00a7\u00a8\u0005w\u0000\u0000\u00a8\u00a9"+
		"\u0005-\u0000\u0000\u00a9\u00aa\u0005h\u0000\u0000\u00aa\u00ab\u0005e"+
		"\u0000\u0000\u00ab\u00ac\u0005a\u0000\u0000\u00ac\u00ad\u0005d\u0000\u0000"+
		"\u00ad\u00ae\u0005e\u0000\u0000\u00ae\u00af\u0005r\u0000\u0000\u00af\u015c"+
		"\u0005s\u0000\u0000\u00b0\u00b1\u0005a\u0000\u0000\u00b1\u00b2\u0005c"+
		"\u0000\u0000\u00b2\u00b3\u0005c\u0000\u0000\u00b3\u00b4\u0005e\u0000\u0000"+
		"\u00b4\u00b5\u0005s\u0000\u0000\u00b5\u00b6\u0005s\u0000\u0000\u00b6\u00b7"+
		"\u0005-\u0000\u0000\u00b7\u00b8\u0005c\u0000\u0000\u00b8\u00b9\u0005o"+
		"\u0000\u0000\u00b9\u00ba\u0005n\u0000\u0000\u00ba\u00bb\u0005t\u0000\u0000"+
		"\u00bb\u00bc\u0005r\u0000\u0000\u00bc\u00bd\u0005o\u0000\u0000\u00bd\u00be"+
		"\u0005l\u0000\u0000\u00be\u00bf\u0005-\u0000\u0000\u00bf\u00c0\u0005a"+
		"\u0000\u0000\u00c0\u00c1\u0005l\u0000\u0000\u00c1\u00c2\u0005l\u0000\u0000"+
		"\u00c2\u00c3\u0005o\u0000\u0000\u00c3\u00c4\u0005w\u0000\u0000\u00c4\u00c5"+
		"\u0005-\u0000\u0000\u00c5\u00c6\u0005m\u0000\u0000\u00c6\u00c7\u0005e"+
		"\u0000\u0000\u00c7\u00c8\u0005t\u0000\u0000\u00c8\u00c9\u0005h\u0000\u0000"+
		"\u00c9\u00ca\u0005o\u0000\u0000\u00ca\u00cb\u0005d\u0000\u0000\u00cb\u015c"+
		"\u0005s\u0000\u0000\u00cc\u00cd\u0005a\u0000\u0000\u00cd\u00ce\u0005c"+
		"\u0000\u0000\u00ce\u00cf\u0005c\u0000\u0000\u00cf\u00d0\u0005e\u0000\u0000"+
		"\u00d0\u00d1\u0005s\u0000\u0000\u00d1\u00d2\u0005s\u0000\u0000\u00d2\u00d3"+
		"\u0005-\u0000\u0000\u00d3\u00d4\u0005c\u0000\u0000\u00d4\u00d5\u0005o"+
		"\u0000\u0000\u00d5\u00d6\u0005n\u0000\u0000\u00d6\u00d7\u0005t\u0000\u0000"+
		"\u00d7\u00d8\u0005r\u0000\u0000\u00d8\u00d9\u0005o\u0000\u0000\u00d9\u00da"+
		"\u0005l\u0000\u0000\u00da\u00db\u0005-\u0000\u0000\u00db\u00dc\u0005a"+
		"\u0000\u0000\u00dc\u00dd\u0005l\u0000\u0000\u00dd\u00de\u0005l\u0000\u0000"+
		"\u00de\u00df\u0005o\u0000\u0000\u00df\u00e0\u0005w\u0000\u0000\u00e0\u00e1"+
		"\u0005-\u0000\u0000\u00e1\u00e2\u0005o\u0000\u0000\u00e2\u00e3\u0005r"+
		"\u0000\u0000\u00e3\u00e4\u0005i\u0000\u0000\u00e4\u00e5\u0005g\u0000\u0000"+
		"\u00e5\u00e6\u0005i\u0000\u0000\u00e6\u015c\u0005n\u0000\u0000\u00e7\u00e8"+
		"\u0005a\u0000\u0000\u00e8\u00e9\u0005c\u0000\u0000\u00e9\u00ea\u0005c"+
		"\u0000\u0000\u00ea\u00eb\u0005e\u0000\u0000\u00eb\u00ec\u0005s\u0000\u0000"+
		"\u00ec\u00ed\u0005s\u0000\u0000\u00ed\u00ee\u0005-\u0000\u0000\u00ee\u00ef"+
		"\u0005c\u0000\u0000\u00ef\u00f0\u0005o\u0000\u0000\u00f0\u00f1\u0005n"+
		"\u0000\u0000\u00f1\u00f2\u0005t\u0000\u0000\u00f2\u00f3\u0005r\u0000\u0000"+
		"\u00f3\u00f4\u0005o\u0000\u0000\u00f4\u00f5\u0005l\u0000\u0000\u00f5\u00f6"+
		"\u0005-\u0000\u0000\u00f6\u00f7\u0005e\u0000\u0000\u00f7\u00f8\u0005x"+
		"\u0000\u0000\u00f8\u00f9\u0005p\u0000\u0000\u00f9\u00fa\u0005o\u0000\u0000"+
		"\u00fa\u00fb\u0005s\u0000\u0000\u00fb\u00fc\u0005e\u0000\u0000\u00fc\u00fd"+
		"\u0005-\u0000\u0000\u00fd\u00fe\u0005h\u0000\u0000\u00fe\u00ff\u0005e"+
		"\u0000\u0000\u00ff\u0100\u0005a\u0000\u0000\u0100\u0101\u0005d\u0000\u0000"+
		"\u0101\u0102\u0005e\u0000\u0000\u0102\u0103\u0005r\u0000\u0000\u0103\u015c"+
		"\u0005s\u0000\u0000\u0104\u0105\u0005a\u0000\u0000\u0105\u0106\u0005c"+
		"\u0000\u0000\u0106\u0107\u0005c\u0000\u0000\u0107\u0108\u0005e\u0000\u0000"+
		"\u0108\u0109\u0005s\u0000\u0000\u0109\u010a\u0005s\u0000\u0000\u010a\u010b"+
		"\u0005-\u0000\u0000\u010b\u010c\u0005c\u0000\u0000\u010c\u010d\u0005o"+
		"\u0000\u0000\u010d\u010e\u0005n\u0000\u0000\u010e\u010f\u0005t\u0000\u0000"+
		"\u010f\u0110\u0005r\u0000\u0000\u0110\u0111\u0005o\u0000\u0000\u0111\u0112"+
		"\u0005l\u0000\u0000\u0112\u0113\u0005-\u0000\u0000\u0113\u0114\u0005m"+
		"\u0000\u0000\u0114\u0115\u0005a\u0000\u0000\u0115\u0116\u0005x\u0000\u0000"+
		"\u0116\u0117\u0005-\u0000\u0000\u0117\u0118\u0005a\u0000\u0000\u0118\u0119"+
		"\u0005g\u0000\u0000\u0119\u015c\u0005e\u0000\u0000\u011a\u011b\u0005a"+
		"\u0000\u0000\u011b\u011c\u0005c\u0000\u0000\u011c\u011d\u0005c\u0000\u0000"+
		"\u011d\u011e\u0005e\u0000\u0000\u011e\u011f\u0005s\u0000\u0000\u011f\u0120"+
		"\u0005s\u0000\u0000\u0120\u0121\u0005-\u0000\u0000\u0121\u0122\u0005c"+
		"\u0000\u0000\u0122\u0123\u0005o\u0000\u0000\u0123\u0124\u0005n\u0000\u0000"+
		"\u0124\u0125\u0005t\u0000\u0000\u0125\u0126\u0005r\u0000\u0000\u0126\u0127"+
		"\u0005o\u0000\u0000\u0127\u0128\u0005l\u0000\u0000\u0128\u0129\u0005-"+
		"\u0000\u0000\u0129\u012a\u0005r\u0000\u0000\u012a\u012b\u0005e\u0000\u0000"+
		"\u012b\u012c\u0005q\u0000\u0000\u012c\u012d\u0005u\u0000\u0000\u012d\u012e"+
		"\u0005e\u0000\u0000\u012e\u012f\u0005s\u0000\u0000\u012f\u0130\u0005t"+
		"\u0000\u0000\u0130\u0131\u0005-\u0000\u0000\u0131\u0132\u0005h\u0000\u0000"+
		"\u0132\u0133\u0005e\u0000\u0000\u0133\u0134\u0005a\u0000\u0000\u0134\u0135"+
		"\u0005d\u0000\u0000\u0135\u0136\u0005e\u0000\u0000\u0136\u0137\u0005r"+
		"\u0000\u0000\u0137\u015c\u0005s\u0000\u0000\u0138\u0139\u0005a\u0000\u0000"+
		"\u0139\u013a\u0005c\u0000\u0000\u013a\u013b\u0005c\u0000\u0000\u013b\u013c"+
		"\u0005e\u0000\u0000\u013c\u013d\u0005s\u0000\u0000\u013d\u013e\u0005s"+
		"\u0000\u0000\u013e\u013f\u0005-\u0000\u0000\u013f\u0140\u0005c\u0000\u0000"+
		"\u0140\u0141\u0005o\u0000\u0000\u0141\u0142\u0005n\u0000\u0000\u0142\u0143"+
		"\u0005t\u0000\u0000\u0143\u0144\u0005r\u0000\u0000\u0144\u0145\u0005o"+
		"\u0000\u0000\u0145\u0146\u0005l\u0000\u0000\u0146\u0147\u0005-\u0000\u0000"+
		"\u0147\u0148\u0005r\u0000\u0000\u0148\u0149\u0005e\u0000\u0000\u0149\u014a"+
		"\u0005q\u0000\u0000\u014a\u014b\u0005u\u0000\u0000\u014b\u014c\u0005e"+
		"\u0000\u0000\u014c\u014d\u0005s\u0000\u0000\u014d\u014e\u0005t\u0000\u0000"+
		"\u014e\u014f\u0005-\u0000\u0000\u014f\u0150\u0005m\u0000\u0000\u0150\u0151"+
		"\u0005e\u0000\u0000\u0151\u0152\u0005t\u0000\u0000\u0152\u0153\u0005h"+
		"\u0000\u0000\u0153\u0154\u0005o\u0000\u0000\u0154\u015c\u0005d\u0000\u0000"+
		"\u0155\u0156\u0005o\u0000\u0000\u0156\u0157\u0005r\u0000\u0000\u0157\u0158"+
		"\u0005i\u0000\u0000\u0158\u0159\u0005g\u0000\u0000\u0159\u015a\u0005i"+
		"\u0000\u0000\u015a\u015c\u0005n\u0000\u0000\u015bt\u0001\u0000\u0000\u0000"+
		"\u015b\u0094\u0001\u0000\u0000\u0000\u015b\u00b0\u0001\u0000\u0000\u0000"+
		"\u015b\u00cc\u0001\u0000\u0000\u0000\u015b\u00e7\u0001\u0000\u0000\u0000"+
		"\u015b\u0104\u0001\u0000\u0000\u0000\u015b\u011a\u0001\u0000\u0000\u0000"+
		"\u015b\u0138\u0001\u0000\u0000\u0000\u015b\u0155\u0001\u0000\u0000\u0000"+
		"\u015c\u015d\u0001\u0000\u0000\u0000\u015d\u015e\u0005\'\u0000\u0000\u015e"+
		"\f\u0001\u0000\u0000\u0000\u015f\u0160\u0005b\u0000\u0000\u0160\u0161"+
		"\u0005a\u0000\u0000\u0161\u0162\u0005s\u0000\u0000\u0162\u0163\u0005i"+
		"\u0000\u0000\u0163\u016b\u0005c\u0000\u0000\u0164\u0165\u0005j\u0000\u0000"+
		"\u0165\u0166\u0005w\u0000\u0000\u0166\u016b\u0005t\u0000\u0000\u0167\u0168"+
		"\u0005l\u0000\u0000\u0168\u0169\u0005o\u0000\u0000\u0169\u016b\u0005g"+
		"\u0000\u0000\u016a\u015f\u0001\u0000\u0000\u0000\u016a\u0164\u0001\u0000"+
		"\u0000\u0000\u016a\u0167\u0001\u0000\u0000\u0000\u016b\u000e\u0001\u0000"+
		"\u0000\u0000\u016c\u016d\u0005a\u0000\u0000\u016d\u016e\u0005p\u0000\u0000"+
		"\u016e\u0176\u0005i\u0000\u0000\u016f\u0170\u0005s\u0000\u0000\u0170\u0171"+
		"\u0005o\u0000\u0000\u0171\u0172\u0005u\u0000\u0000\u0172\u0173\u0005r"+
		"\u0000\u0000\u0173\u0174\u0005c\u0000\u0000\u0174\u0176\u0005e\u0000\u0000"+
		"\u0175\u016c\u0001\u0000\u0000\u0000\u0175\u016f\u0001\u0000\u0000\u0000"+
		"\u0176\u0010\u0001\u0000\u0000\u0000\u0177\u0178\u0005t\u0000\u0000\u0178"+
		"\u0179\u0005r\u0000\u0000\u0179\u017a\u0005u\u0000\u0000\u017a\u0181\u0005"+
		"e\u0000\u0000\u017b\u017c\u0005f\u0000\u0000\u017c\u017d\u0005a\u0000"+
		"\u0000\u017d\u017e\u0005l\u0000\u0000\u017e\u017f\u0005s\u0000\u0000\u017f"+
		"\u0181\u0005e\u0000\u0000\u0180\u0177\u0001\u0000\u0000\u0000\u0180\u017b"+
		"\u0001\u0000\u0000\u0000\u0181\u0012\u0001\u0000\u0000\u0000\u0182\u018b"+
		"\u00050\u0000\u0000\u0183\u0187\u0007\u0002\u0000\u0000\u0184\u0186\u0007"+
		"\u0000\u0000\u0000\u0185\u0184\u0001\u0000\u0000\u0000\u0186\u0189\u0001"+
		"\u0000\u0000\u0000\u0187\u0185\u0001\u0000\u0000\u0000\u0187\u0188\u0001"+
		"\u0000\u0000\u0000\u0188\u018b\u0001\u0000\u0000\u0000\u0189\u0187\u0001"+
		"\u0000\u0000\u0000\u018a\u0182\u0001\u0000\u0000\u0000\u018a\u0183\u0001"+
		"\u0000\u0000\u0000\u018b\u0014\u0001\u0000\u0000\u0000\u018c\u0190\u0005"+
		"\'\u0000\u0000\u018d\u018f\b\u0003\u0000\u0000\u018e\u018d\u0001\u0000"+
		"\u0000\u0000\u018f\u0192\u0001\u0000\u0000\u0000\u0190\u018e\u0001\u0000"+
		"\u0000\u0000\u0190\u0191\u0001\u0000\u0000\u0000\u0191\u0193\u0001\u0000"+
		"\u0000\u0000\u0192\u0190\u0001\u0000\u0000\u0000\u0193\u0194\u0005\'\u0000"+
		"\u0000\u0194\u0016\u0001\u0000\u0000\u0000\u0195\u0196\u00050\u0000\u0000"+
		"\u0196\u0197\u0005.\u0000\u0000\u0197\u019b\u0001\u0000\u0000\u0000\u0198"+
		"\u019a\u0007\u0000\u0000\u0000\u0199\u0198\u0001\u0000\u0000\u0000\u019a"+
		"\u019d\u0001\u0000\u0000\u0000\u019b\u0199\u0001\u0000\u0000\u0000\u019b"+
		"\u019c\u0001\u0000\u0000\u0000\u019c\u01ab\u0001\u0000\u0000\u0000\u019d"+
		"\u019b\u0001\u0000\u0000\u0000\u019e\u01a0\u0007\u0002\u0000\u0000\u019f"+
		"\u019e\u0001\u0000\u0000\u0000\u01a0\u01a1\u0001\u0000\u0000\u0000\u01a1"+
		"\u019f\u0001\u0000\u0000\u0000\u01a1\u01a2\u0001\u0000\u0000\u0000\u01a2"+
		"\u01a3\u0001\u0000\u0000\u0000\u01a3\u01a7\u0005.\u0000\u0000\u01a4\u01a6"+
		"\u0007\u0000\u0000\u0000\u01a5\u01a4\u0001\u0000\u0000\u0000\u01a6\u01a9"+
		"\u0001\u0000\u0000\u0000\u01a7\u01a5\u0001\u0000\u0000\u0000\u01a7\u01a8"+
		"\u0001\u0000\u0000\u0000\u01a8\u01ab\u0001\u0000\u0000\u0000\u01a9\u01a7"+
		"\u0001\u0000\u0000\u0000\u01aa\u0195\u0001\u0000\u0000\u0000\u01aa\u019f"+
		"\u0001\u0000\u0000\u0000\u01ab\u0018\u0001\u0000\u0000\u0000\u01ac\u01b1"+
		"\u0005[\u0000\u0000\u01ad\u01b2\u0003\u0011\b\u0000\u01ae\u01b2\u0003"+
		"\u0013\t\u0000\u01af\u01b2\u0003\u0015\n\u0000\u01b0\u01b2\u0003\u0017"+
		"\u000b\u0000\u01b1\u01ad\u0001\u0000\u0000\u0000\u01b1\u01ae\u0001\u0000"+
		"\u0000\u0000\u01b1\u01af\u0001\u0000\u0000\u0000\u01b1\u01b0\u0001\u0000"+
		"\u0000\u0000\u01b1\u01b2\u0001\u0000\u0000\u0000\u01b2\u01bc\u0001\u0000"+
		"\u0000\u0000\u01b3\u01b8\u0005,\u0000\u0000\u01b4\u01b9\u0003\u0011\b"+
		"\u0000\u01b5\u01b9\u0003\u0013\t\u0000\u01b6\u01b9\u0003\u0015\n\u0000"+
		"\u01b7\u01b9\u0003\u0017\u000b\u0000\u01b8\u01b4\u0001\u0000\u0000\u0000"+
		"\u01b8\u01b5\u0001\u0000\u0000\u0000\u01b8\u01b6\u0001\u0000\u0000\u0000"+
		"\u01b8\u01b7\u0001\u0000\u0000\u0000\u01b9\u01bb\u0001\u0000\u0000\u0000"+
		"\u01ba\u01b3\u0001\u0000\u0000\u0000\u01bb\u01be\u0001\u0000\u0000\u0000"+
		"\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bc\u01bd\u0001\u0000\u0000\u0000"+
		"\u01bd\u01bf\u0001\u0000\u0000\u0000\u01be\u01bc\u0001\u0000\u0000\u0000"+
		"\u01bf\u01c0\u0005]\u0000\u0000\u01c0\u001a\u0001\u0000\u0000\u0000\u01c1"+
		"\u01c2\u0007\u0004\u0000\u0000\u01c2\u001c\u0001\u0000\u0000\u0000\u01c3"+
		"\u01c4\u0007\u0005\u0000\u0000\u01c4\u001e\u0001\u0000\u0000\u0000\u01c5"+
		"\u01c6\u0005=\u0000\u0000\u01c6\u01ca\u0005=\u0000\u0000\u01c7\u01c8\u0005"+
		"!\u0000\u0000\u01c8\u01ca\u0005=\u0000\u0000\u01c9\u01c5\u0001\u0000\u0000"+
		"\u0000\u01c9\u01c7\u0001\u0000\u0000\u0000\u01ca \u0001\u0000\u0000\u0000"+
		"\u01cb\u01cc\u0005\n\u0000\u0000\u01cc\"\u0001\u0000\u0000\u0000\u01cd"+
		"\u01ce\u0007\u0006\u0000\u0000\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cf"+
		"\u01d0\u0006\u0011\u0000\u0000\u01d0$\u0001\u0000\u0000\u0000\u01d1\u01d2"+
		"\u0005/\u0000\u0000\u01d2\u01d3\u0005/\u0000\u0000\u01d3\u01d7\u0001\u0000"+
		"\u0000\u0000\u01d4\u01d6\b\u0007\u0000\u0000\u01d5\u01d4\u0001\u0000\u0000"+
		"\u0000\u01d6\u01d9\u0001\u0000\u0000\u0000\u01d7\u01d5\u0001\u0000\u0000"+
		"\u0000\u01d7\u01d8\u0001\u0000\u0000\u0000\u01d8\u01da\u0001\u0000\u0000"+
		"\u0000\u01d9\u01d7\u0001\u0000\u0000\u0000\u01da\u01db\u0006\u0012\u0000"+
		"\u0000\u01db&\u0001\u0000\u0000\u0000\u01dc\u01dd\u0005{\u0000\u0000\u01dd"+
		"(\u0001\u0000\u0000\u0000\u01de\u01df\u0005}\u0000\u0000\u01df*\u0001"+
		"\u0000\u0000\u0000\u01e0\u01e1\u0005(\u0000\u0000\u01e1,\u0001\u0000\u0000"+
		"\u0000\u01e2\u01e3\u0005)\u0000\u0000\u01e3.\u0001\u0000\u0000\u0000\u01e4"+
		"\u01e5\u0005=\u0000\u0000\u01e50\u0001\u0000\u0000\u0000\u01e6\u01e8\u0007"+
		"\u0001\u0000\u0000\u01e7\u01e6\u0001\u0000\u0000\u0000\u01e8\u01e9\u0001"+
		"\u0000\u0000\u0000\u01e9\u01e7\u0001\u0000\u0000\u0000\u01e9\u01ea\u0001"+
		"\u0000\u0000\u0000\u01ea2\u0001\u0000\u0000\u0000\u01eb\u01ed\u0005$\u0000"+
		"\u0000\u01ec\u01ee\u0007\u0001\u0000\u0000\u01ed\u01ec\u0001\u0000\u0000"+
		"\u0000\u01ee\u01ef\u0001\u0000\u0000\u0000\u01ef\u01ed\u0001\u0000\u0000"+
		"\u0000\u01ef\u01f0\u0001\u0000\u0000\u0000\u01f0\u01f4\u0001\u0000\u0000"+
		"\u0000\u01f1\u01f3\u0007\b\u0000\u0000\u01f2\u01f1\u0001\u0000\u0000\u0000"+
		"\u01f3\u01f6\u0001\u0000\u0000\u0000\u01f4\u01f2\u0001\u0000\u0000\u0000"+
		"\u01f4\u01f5\u0001\u0000\u0000\u0000\u01f54\u0001\u0000\u0000\u0000\u01f6"+
		"\u01f4\u0001\u0000\u0000\u0000\u001a\u0000GMOflq\u015b\u016a\u0175\u0180"+
		"\u0187\u018a\u0190\u019b\u01a1\u01a7\u01aa\u01b1\u01b8\u01bc\u01c9\u01d7"+
		"\u01e9\u01ef\u01f4\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}