// Generated from /Users/jixinyu/Documents/Code/Go/dal/grammars/Common.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link CommonParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface CommonVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link CommonParser#version}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVersion(CommonParser.VersionContext ctx);
	/**
	 * Visit a parse tree produced by {@link CommonParser#name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitName(CommonParser.NameContext ctx);
}