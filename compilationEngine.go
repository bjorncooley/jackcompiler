package main 

import (
    "fmt"
    "log"
)

type CompilerState string
const (
    Initial CompilerState = "initial"
    CompilingClass CompilerState = "compilingClass"
    CompilingFunction CompilerState = "compilingFunction"
    CompilingMethod CompilerState = "compilingMethod"
)

var compilerState = []CompilerState{Initial}

type CompilationEngine struct {
}

var tokenIndex int = 0

func (engine *CompilationEngine) Compile(tokenList []Token) {
    for tokenIndex < len(tokenList) {
        parseToken(tokenList[tokenIndex])
        tokenIndex += 1
    }
}

func parseToken(token Token) {
    switch(token.tokenType) {
    case Keyword:
        compileKeyword(token)
    default:
        compileStatement(token)
    }
}

func compileKeyword(token Token) {
    lexeme := token.lexeme
    if (lexeme == "class") {
        checkAllowedState(CompilingClass, lexeme)
        compilerState = append(compilerState, CompilingClass)
        compileClass(token)
    } else if (lexeme == "function") {
        checkAllowedState(CompilingFunction, lexeme)
        compilerState = append(compilerState, CompilingFunction)
        compileFunction(token)
    } else if (lexeme == "method") {
        checkAllowedState(CompilingMethod, lexeme)
        compilerState = append(compilerState, CompilingMethod)
        compileMethod(token)
    }
}

func compileClass(token Token) {
    checkValidToken(token, Keyword)
    output("<class>")

    token = advanceToNextToken()
    checkValidToken(token, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", token.lexeme))

    token = advanceToNextToken()
    checkValidToken(token, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", token.lexeme))

    token = advanceToNextToken()
    parseToken(token)

    output("</class>")
}

func checkValidToken(token Token, expected TokenType) {
    if token.tokenType != expected {
        log.Fatal(fmt.Sprintf("Syntax error: unexpected %s, expected token of type %s", token.lexeme, expected))
    }
}

func advanceToNextToken() Token {
    tokenIndex += 1
    return tokenList[tokenIndex]
}

func compileFunction(token Token) {
    fmt.Printf("compiling function\n")
}

func compileMethod(token Token) {
    checkValidToken(token, Keyword)
    output("<subroutineDec>")

    token = advanceToNextToken()
    checkValidToken(token, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", token.lexeme))

    token = advanceToNextToken()
    checkValidToken(token, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", token.lexeme))

    token = advanceToNextToken()
    checkValidToken(token, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", token.lexeme))
    output("<parameterList>")

    token = advanceToNextToken()
    compileParameterList(token)
    output("</parameterList>")

    token = advanceToNextToken()
    checkValidToken(token, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", token.lexeme))

    advanceToNextToken()
    for tokenList[tokenIndex].lexeme != "}" {
        compileStatement(tokenList[tokenIndex])
    }

    output("</subroutineDec>")
}

func compileStatement(token Token) {
    if token.lexeme == "var" {
        output("<varDec>")
        compileVarStatement()
        output("</varDec>")
    }
    fmt.Printf("compiling statement with %s\n", token.lexeme)
    advanceToNextToken()
}

func compileVarStatement() {
    output("<keyword>var</keyword>")

    expectedType := Keyword
    outputTag := "keyword"
    token := advanceToNextToken()

    for token.lexeme != ";" {

        checkValidToken(token, expectedType)
        if token.lexeme == "," {
            output(fmt.Sprintf("<%s>%s</%s>", outputTag, token.lexeme, outputTag))
        }

        if expectedType == Identifier {
            checkValidVarDeclarationSyntax()
        }

        if expectedType == Keyword {
            expectedType = Identifier
            outputTag = "identifier"
        } else if expectedType == Identifier {
            expectedType = Symbol
            outputTag = "symbol"
        } else {
            expectedType = Identifier
            outputTag = "identifier"
        }

        token = advanceToNextToken()
    }
}

func checkValidVarDeclarationSyntax() {
    nextToken := tokenList[tokenIndex+1]

    if nextToken.lexeme != ";" {
        return
    }

    if nextToken.lexeme == "," && tokenList[tokenIndex+2].lexeme == ";" {
        log.Fatal("Syntax error: unexpected token ,")
    }

    if nextToken.tokenType == Identifier {
        log.Fatal(fmt.Sprintf("Syntax error: unexpected %s, expected ,", nextToken.lexeme))
    }
}

func compileParameterList(token Token) {
    expectedType := Keyword
    outputTag := "keyword"
    for token.lexeme != ")" {

        checkValidToken(token, expectedType)
        output(fmt.Sprintf("<%s>%s</%s>", outputTag, token.lexeme, outputTag))

        if expectedType == Identifier {
            checkValidParameterListSyntax()
        }

        if expectedType == Keyword {
            expectedType = Identifier
            outputTag = "identifier"
        } else {
            expectedType = Keyword
            outputTag = "keyword"
        }

        token = advanceToNextToken()
    }
}

func checkValidParameterListSyntax() {
    nextToken := tokenList[tokenIndex+1].lexeme
    if nextToken == ")" {
        return
    }

    if tokenList[tokenIndex+2].tokenType == Identifier && nextToken != "," {
        log.Fatal(fmt.Sprintf("Syntax error: unexpected %s, expected ,", nextToken))
    }

    if tokenList[tokenIndex+2].tokenType != Keyword && nextToken == "," {
        log.Fatal("Syntax error: unexpected ,")
    }

    advanceToNextToken()
}

// Utils

func checkAllowedState(state CompilerState, token string) {
    if compilerState[len(compilerState) - 1] == state {
        log.Fatal("Invalid syntax: %s", token)
    }
}

func output(str string) {
    fmt.Printf("%s\n", str)
}
