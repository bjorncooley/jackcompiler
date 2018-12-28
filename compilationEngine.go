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
        log.Fatal(fmt.Sprintf("Invalid syntax: %s. Expected symbol of type %s", tokenObject.token, expected))
    }
}

func advanceToNextToken() TokenObject {
    tokenIndex += 1
    return tokenList[tokenIndex]
}

func compileStatement(tokenObject TokenObject) {
    fmt.Printf("compiling statement with %s\n", tokenObject.token)
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

    output("</subroutineDec>")
}

func compileParameterList(tokenObject TokenObject) {
    expectedType := Keyword
    outputTag := "keyword"
    for tokenObject.token != ")" {
        if tokenObject.token == "," {
            tokenObject = advanceToNextToken()
            continue
        }
        checkValidToken(tokenObject, expectedType)
        output(fmt.Sprintf("<%s>%s</%s>", outputTag, tokenObject.token, outputTag))
        tokenObject = advanceToNextToken()

        if expectedType == Keyword {
            expectedType = Identifier
            outputTag = "identifier"
        } else {
            expectedType = Keyword
            outputTag = "keyword"
        }
    }
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
