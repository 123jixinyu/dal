grammar Common;


//版本
version : 'Version' Version NewLine;
//名称
name : 'Name' ID NewLine;


//版本
Version: Kind'/'[0-9]+(.[0-9]+)?;

//Methods
Method: 'GET'|'POST'|'PUT'|'DELETE'|'PATCH';

//Route
Route: ('/'[a-zA-Z]*)+;

//AccessControlHeader
PreflightHeader:  '\'' ('access-control-allow-credentials'
               | 'access-control-allow-headers'
               | 'access-control-allow-methods'
               | 'access-control-allow-origin'
               | 'access-control-expose-headers'
               | 'access-control-max-age'
               | 'access-control-request-headers'
               | 'access-control-request-method'
               | 'origin'
               )'\'';


//插件
Plugin: 'basic'| 'jwt' | 'log';

//DSL 类型
fragment Kind : 'api' | 'source';

//类型
Bool: 'true' | 'false';
Int: '0'|[1-9][0-9]*;
String:  '\''(~'\'')*'\'';
Float: ('0.' [0-9]* ) | ([1-9]+'.' [0-9]*);
Array: '[' (Bool|Int|String|Float)? (',' (Bool|Int|String|Float))* ']';

//操作
MuOperator: '*'
        | '/'
        ;
ArOperator: '+'
        | '-'
        ;
JudgeOperator:'==' | '!=';

//以下是通用符号
NewLine: '\n' ; // 以分号为终止符
WS: [ \r\t] -> skip; // 忽略TAB、换行、空白
LINE_COMMENT : '//' ~[\r\n]* -> skip; // 注释
LeftBrace: '{';
RightBrace: '}';
LeftParenthesis: '(';
RightParenthesis: ')';

Equal: '=';
//标识符、变量名
ID : [a-zA-Z]+ ;
Var : '$'[a-zA-Z]+([.a-zA-Z])* ;
