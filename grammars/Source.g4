grammar Source;

import Common;

@parser::members{

}

@lexer::members{

}

//全部规则
whole:
version
name
dsn
timeout?
try?
EOF
;
// mysql://[user[:password]@][net]@[loc:port][,...][/schema/dbname][?param1=value1&...]
dsn: 'DSN' String  NewLine*;
timeout: 'Connect-Timeout' Int NewLine*;
try: 'Try_Times' Int NewLine*;
