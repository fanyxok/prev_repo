/*
 *  The scanner definition for COOL.
 */

/*
 *  Stuff enclosed in %{ %} in the first section is copied verbatim to the
 *  output, so headers and global definitions are placed here to be visible
 * to the code in the file.  Don't remove anything that was here initially
 */
%{
#include <cool-parse.h>
#include <stringtab.h>
#include <utilities.h>
#include <ast-parse.h>

/* The compiler assumes these identifiers. */
#define yylval cool_yylval
#define yylex  cool_yylex

/* Max size of string constants */
#define MAX_STR_CONST 1025
#define YY_NO_UNPUT   /* keep g++ happy */

extern FILE *fin; /* we read from this file */

/* define YY_INPUT so we read from the FILE fin:
 * This change makes it possible to use this scanner in
 * the Cool compiler.
 */
#undef YY_INPUT
#define YY_INPUT(buf,result,max_size) \
  if ( (result = fread( (char*)buf, sizeof(char), max_size, fin)) < 0) \
    YY_FATAL_ERROR( "read() in flex scanner failed");

char string_buf[MAX_STR_CONST]; /* to assemble string constants */
char *string_buf_ptr;

extern int curr_lineno;

extern YYSTYPE cool_yylval;

unsigned int comment_nested = 0;
unsigned int string_buf_left;
bool string_error;
char *error_msg;
char *backslash_comstr()
{
  char *c = &yytext[1];
  if (*c == '\n')
    curr_lineno++;
  return c;
} 

int str_copy(char *c, int l)
{
  if (l < string_buf_left)
  {
    strncpy(string_buf_ptr, c, l);
    string_buf_ptr += l;
    string_buf_left -= l;
    return 1;
  }else
  {
    string_buf_left -= l;
    error_msg = "String constant too long";
    return 0;
  }

}

/*
 *  Add Your own definitions here
 */

%}

%option noyywrap

/*
 * Define names for regular expressions here.
 */
CLASS       [cC][lL][aA][sS][sS]   
INHERITS    [iI][nN][hH][eE][rR][iI][tT][sS]
INT_T       Int
STRING_T    String
BOOL_T      Bool
LIST_T      List
IO_T        IO



OBJECT_ID   [a-z][a-zA-Z0-9_]*
TYPE_ID     [A-Z][a-zA-Z0-9_]*

SELF_TYPE   SELF_TYPE 
SELF        self

IF          [iI][fF]
THEN        [tT][hH][eE][nN]
ELSE        [eE][lL][sS][eE]
FI          [fF][iI]

WHILE       [wW][hH][iI][lL][eE]
LOOP        [lL][oO][oO][pP]
POOL        [pP][oO][oO][lL]

LET         [lL][eE][tT]
IN          [iI][nN]

CASE        [cC][aA][sS][eE]
OF          [oO][fF]
ESAC        [eE][sS][aA][cC]


NEW         [nB][eE][wW]
ISVOID      [iI][sS][vV][oO][iI][dD]
NOT         [nN][oO][tT]

LE          <=
DARROW      =>
ASSIG       <-

TRUE        t[rR][uU][eE]
FALSE       f[aA][lL][sS][eE]

DELIM       [ \t\f\r\v]
WS          {DELIM}+
DIGIT       [0-9]
LETTER      [A-Za-z]
ID          {letter}({letter}|{digit}|_)*

QUOTES      \"
STR_CONST   QUOTES{{LETTER}|{DIGIT}}+QUOTES 
INT_CONST   {DIGIT}+

STAR        [*]
NOTSTAR     [^*]
LP          [(]
NOTLP       [^(]
RP          [)]
NOTRP       [^)]

NULLCHAR    [\0]
ELEM_COMNT  [^\n(*\\]
ELEM_STRING [^\0\\\"\n]

NEWLINE     [\n]
NOTNEWLINE  [^\n]

BACKSLASH   [\\]
COMNT_LINE  "--"
COMNT_START "(*"
COMNT_END   "*)"


%x COMMENT
%x STRING

/*
 *  + - * / < > = ~
 */

%%

 /*
  * Define regular expressions for the tokens of COOL here. Make sure, you
  * handle correctly special cases, like:
  *   - Nested comments
  *   - String constants: They use C like systax and can contain escape
  *     sequences. Escape sequence \c is accepted for all characters c. Except
  *     for \n \t \b \f, the result is c.
  *   - Keywords: They are case-insensitive except for the values true and
  *     false, which must begin with a lower-case letter.
  *   - Multiple-character operators (like <-): The scanner should produce a
  *     single token for every such operator.
  *   - Line counting: You should keep the global variable curr_lineno updated
  *     with the correct line number
  */

 /*  
  *  Deal with block comment and line comment 
  */
<INITIAL,COMMENT>{NEWLINE}  {
  curr_lineno++;
}

{COMNT_START} {
  comment_nested++;
  BEGIN(COMMENT);
}

<COMMENT><<EOF>>  {
  yylval.error_msg="EOF in comment";
  BEGIN(0);
  return(ERROR);
}

<COMMENT>{STAR}/{NOTRP}  ;

<COMMENT>{LP}/{NOTSTAR}  ;

<COMMENT>{ELEM_COMNT}*    ;

<COMMENT>{COMNT_START}  {
  comment_nested++;
}

<COMMENT>{BACKSLASH}(.|NEWLINE) {
  backslash_comstr();
}
<COMMENT>{BACKSLASH}  {}

<COMMENT>{COMNT_END}  {
  comment_nested--;
  if (comment_nested == 0)
  {
    BEGIN(INITIAL);
  }
}

<INITIAL>{COMNT_END}  {
  yylval.error_msg="Unmatched *)";
  return(ERROR);
}



<INITIAL>{COMNT_LINE}{NOTNEWLINE}*   {}


 /*
  *----------------------------------------
  * Deal with string
  */

<INITIAL>{QUOTES} {
  BEGIN(STRING);
  string_buf_ptr = string_buf;
  string_buf_left = MAX_STR_CONST;
  string_error = false;
  error_msg = "";
}

<STRING><<EOF>> {
  yylval.error_msg = "EOF in string constant";
  BEGIN(0);
  return ERROR;
}

<STRING>\0 {
  if (string_error == false)
  {
    error_msg = "String contains null character.";
  }
  string_error = true;
}

<STRING>{BACKSLASH}(\0)  {
  string_buf_ptr++;
  string_buf_left--;
  if (strcmp(error_msg,"")==0)
  {
    error_msg = "String contains null character.";
  }
  string_error = true;
}


<STRING>{BACKSLASH}{NOTNEWLINE}  {

  char *c = backslash_comstr();
  int escape;

  if       (strcmp(c,"b")==0)
  {
    escape = str_copy("\b", 1);
  }else if (strcmp(c,"t")==0)
  {
    escape = str_copy("\t", 1);
  }else if (strcmp(c,"n")==0)
  {
    escape = str_copy("\n", 1);
  }else if (strcmp(c,"f")==0)
  {
    escape = str_copy("\f", 1);
  }else
  {
    escape = str_copy(c, 1);
  }
  if (escape == 0){
    string_error = true;
  }
}

<STRING>{NEWLINE} {
  curr_lineno++;
  yylval.error_msg = "Unterminated string constant";
  BEGIN(0);
  return ERROR;
}

<STRING>{BACKSLASH}{NEWLINE}  {
  curr_lineno++;
  int copy_right = str_copy("\n", 1);
  if (copy_right == 0)
  {
    string_error = true;
  }
}


<STRING>{ELEM_STRING}+ {
  int rt = str_copy(yytext, strlen(yytext));
  if (rt == 0){
    string_error = true;
  }
}

<STRING>{QUOTES}  {
  if (string_error == true)
  {
    yylval.error_msg = error_msg;
    BEGIN(0);
    strncpy(string_buf, "", MAX_STR_CONST);
    return(ERROR);
  }else
  {
    *string_buf_ptr = '\0';
    yylval.symbol = stringtable.add_string(string_buf, string_buf_ptr - string_buf + 1);
    BEGIN(0);
    strncpy(string_buf, "", MAX_STR_CONST);
    return(275);
  }
}
 

<INITIAL>{WS} {}

 /*----------------------------------------------------------------
  * Deal with Keywords
  */


<INITIAL>{TRUE}   { yylval.boolean = true; return (277); }
<INITIAL>{FALSE}    { yylval.boolean = false; return (277); }

<INITIAL>{CLASS}  { return (258); }
<INITIAL>{INT_CONST}  { yylval.symbol = stringtable.add_string(yytext); return (276); }

<INITIAL>{IF} { return (261); }
<INITIAL>{ELSE} { return (259); }
<INITIAL>{THEN} { return (267); }
<INITIAL>{FI} { return (260); }

<INITIAL>{LOOP} { return (265); }
<INITIAL>{POOL} { return (266); }
<INITIAL>{WHILE}  { return (268); }

<INITIAL>{LET}  { return (264); }
<INITIAL>{IN} { return (262); }

<INITIAL>{CASE} { return (269); }
<INITIAL>{ESAC} { return (270); }
<INITIAL>{OF} { return (271); }

<INITIAL>{DARROW} { return (272); }
<INITIAL>{ASSIG} { return (280); }
<INITIAL>{LE} { return (282); }

<INITIAL>{INHERITS} { return (263); }
<INITIAL>{ISVOID} { return (274); }

<INITIAL>{NEW}  { return (273); }
<INITIAL>{NOT}  { return (281); }

<INITIAL>{TYPE_ID}  { yylval.symbol = stringtable.add_string(yytext); return (278); }
<INITIAL>{OBJECT_ID}  { yylval.symbol = stringtable.add_string(yytext); return (279); }


<INITIAL>";"                     { return int(';'); }
<INITIAL>","                     { return int(','); }
<INITIAL>":"                     { return int(':'); }
<INITIAL>"{"                     { return int('{'); }
<INITIAL>"}"                     { return int('}'); }
<INITIAL>"+"                     { return int('+'); }
<INITIAL>"-"                     { return int('-'); }
<INITIAL>"*"                     { return int('*'); }
<INITIAL>"/"                     { return int('/'); }
<INITIAL>"<"                     { return int('<'); }
<INITIAL>"="                     { return int('='); }
<INITIAL>"~"                     { return int('~'); }
<INITIAL>"."                     { return int('.'); }
<INITIAL>"@"                     { return int('@'); }
<INITIAL>"("                     { return int('('); }
<INITIAL>")"                     { return int(')'); }

<INITIAL>.                       { yylval.error_msg = yytext; return (ERROR); }




%%
