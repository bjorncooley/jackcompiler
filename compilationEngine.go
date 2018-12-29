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

func (engine *CompilationEngine) Compile(tokenList []TokenObject) {
    for tokenIndex < len(tokenList) {
        parseTokenObject(tokenList[tokenIndex])
        tokenIndex += 1
    }
}

func parseTokenObject(tokenObject TokenObject) {
    switch(tokenObject.tokenType) {
    case Keyword:
        compileKeyword(tokenObject)
    default:
        compileStatement(tokenObject)
    }
}

func compileKeyword(tokenObject TokenObject) {
    token := tokenObject.token
    if (token == "class") {
        checkAllowedState(CompilingClass, token)
        compilerState = append(compilerState, CompilingClass)
        compileClass(tokenObject)
    } else if (token == "function") {
        checkAllowedState(CompilingFunction, token)
        compilerState = append(compilerState, CompilingFunction)
        compileFunction(tokenObject)
    } else if (token == "method") {
        checkAllowedState(CompilingMethod, token)
        compilerState = append(compilerState, CompilingMethod)
        compileMethod(tokenObject)
    }
}

func compileClass(tokenObject TokenObject) {
    checkValidToken(tokenObject, Keyword)
    output("<class>")

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", tokenObject.token))

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", tokenObject.token))

    tokenObject = advanceToNextToken()
    parseTokenObject(tokenObject)

    output("</class>")
}

func checkValidToken(tokenObject TokenObject, expected TokenType) {
    if tokenObject.tokenType != expected {
        log.Fatal(fmt.Sprintf("Syntax error: unexpected %s, expected symbol of type %s", tokenObject.token, expected))
    }
}

func advanceToNextToken() TokenObject {
    tokenIndex += 1
    return tokenList[tokenIndex]
}

func compileFunction(tokenObject TokenObject) {
    fmt.Printf("compiling function\n")
}

func compileMethod(tokenObject TokenObject) {
    checkValidToken(tokenObject, Keyword)
    output("<subroutineDec>")

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", tokenObject.token))

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Identifier)
    output(fmt.Sprintf("<identifier>%s</identifier>", tokenObject.token))

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", tokenObject.token))
    output("<parameterList>")

    tokenObject = advanceToNextToken()
    compileParameterList(tokenObject)
    output("</parameterList>")

    tokenObject = advanceToNextToken()
    checkValidToken(tokenObject, Symbol)
    output(fmt.Sprintf("<symbol>%s</symbol>", tokenObject.token))

    advanceToNextToken()
    for tokenList[tokenIndex].token != "}" {
        compileStatement(tokenList[tokenIndex])
    }

    output("</subroutineDec>")
}

func compileStatement(tokenObject TokenObject) {
    if tokenObject.token == "var" {
        output("<varDec>")
        compileVarStatement()
        output("</varDec>")
    }
    fmt.Printf("compiling statement with %s\n", tokenObject.token)
    advanceToNextToken()
}

func compileVarStatement() {
    output("<keyword>var</keyword>")

    expectedType := Keyword
    outputTag := "keyword"
    tokenObject := advanceToNextToken()

    for tokenObject.token != ";" {

        checkValidToken(tokenObject, expectedType)
        if tokenObject.token != "," {
            output(fmt.Sprintf("<%s>%s</%s>", outputTag, tokenObject.token, outputTag))
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

        tokenObject = advanceToNextToken()
    }
}

func checkValidVarDeclarationSyntax() {
    nextToken := tokenList[tokenIndex+1]

    if nextToken.token == ";" {
        return
    }

    if nextToken.token == "," && tokenList[tokenIndex+2].token == ";" {
        log.Fatal("Syntax error: unexpected token ,")
    }

    if nextToken.tokenType == Identifier {
        log.Fatal(fmt.Sprintf("Syntax error: unexpected %s, expected ,", nextToken.token))
    }
}

func compileParameterList(tokenObject TokenObject) {
    expectedType := Keyword
    outputTag := "keyword"
    for tokenObject.token != ")" {

        checkValidToken(tokenObject, expectedType)
        output(fmt.Sprintf("<%s>%s</%s>", outputTag, tokenObject.token, outputTag))

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

        tokenObject = advanceToNextToken()
    }
}

func checkValidParameterListSyntax() {
    nextToken := tokenList[tokenIndex+1].token
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
