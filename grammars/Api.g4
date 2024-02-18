grammar Api;
import Common;
@parser::members{

}

@lexer::members{

}

//全部规则
whole:
version
name
endpoint
preflight*
plugin*
validate?
stage*
return EOF
;


//版本
version : 'Version' Version NewLine;
//名称
name : 'Name' ID NewLine;
//方法、Endpoint
endpoint : Method Route NewLine;
//preflight
preflight: 'Options'  PreflightHeader  String NewLine;

//参数验证
validate: 'Validate' LeftBrace NewLine* fieldValidates RightBrace NewLine;
fieldValidates: fieldValidate*;
fieldValidate: Var String NewLine+;

//中间件
plugin: 'Middleware' Plugin pluginParams? NewLine;
pluginParams: pluginParam+;
pluginParam: '--'ID'=' String;

//阶段
stage: 'Stage' ID LeftBrace NewLine handles  RightBrace NewLine;

//逻辑
handles: (handle NewLine)+;
handle : ID Equal (query
                | express
                | value
                );

//query
query: 'query' Var sql value*;
//sql
sql: String| Var;


//value
value: Bool| Float |Int | String| Array| Var;

//表达式
express: '"' calculate '"';

//参数
param: value | func;
//三元表达式
metaExpress: param (JudgeOperator param)? '?' calculate ':' calculate;
//计算
calculate:
         | param
         | LeftParenthesis calculate RightParenthesis
         | func
         | metaExpress
         | param MuOperator calculate
         | calculate MuOperator param
         | calculate ArOperator calculate
         ;
//函数
func: ID LeftParenthesis params? RightParenthesis;
//函数参数，支持嵌套
params: param (',' param)*;


//返回
return: 'Return' LeftBrace NewLine+ responseHeaders?  responseCode? content RightBrace NewLine* ;

//响应头
responseHeaders: responseHeader NewLine+ (responseHeader NewLine+)*;
responseHeader: 'Response-Header' String String;

//响应码
responseCode: 'Response-Code' (Int|express) NewLine+;

//内容
content: 'Content' LeftBrace NewLine+ body NewLine+ RightBrace NewLine+;
body: express
    | object
    ;

//对象
object: LeftBrace NewLine+ objectPair* RightBrace;
objectPair: ID ':' objectValue NewLine+;
objectValue: value
           | object
           | express
           ;

