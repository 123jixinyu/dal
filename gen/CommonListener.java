// Generated from /Users/jixinyu/Documents/Code/Go/dal/grammars/Common.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link CommonParser}.
 */
public interface CommonListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link CommonParser#version}.
	 * @param ctx the parse tree
	 */
	void enterVersion(CommonParser.VersionContext ctx);
	/**
	 * Exit a parse tree produced by {@link CommonParser#version}.
	 * @param ctx the parse tree
	 */
	void exitVersion(CommonParser.VersionContext ctx);
	/**
	 * Enter a parse tree produced by {@link CommonParser#name}.
	 * @param ctx the parse tree
	 */
	void enterName(CommonParser.NameContext ctx);
	/**
	 * Exit a parse tree produced by {@link CommonParser#name}.
	 * @param ctx the parse tree
	 */
	void exitName(CommonParser.NameContext ctx);
}