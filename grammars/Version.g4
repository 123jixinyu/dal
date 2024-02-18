grammar Version;
import Common;
whole:
version
name;

//版本
version : 'Version' Version NewLine;
//名称
name : 'Name' ID NewLine;
