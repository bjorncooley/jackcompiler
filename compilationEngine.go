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
        tokenObject := tokenList[tokenIndex]
        switch(tokenObject.tokenType) {
        case Keyword:
            compileKeyword(tokenObject)
        default:
            compileStatement(tokenObject)
        }
        tokenIndex += 1
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
        compileFunction()
    } else if (token == "method") {
        checkAllowedState(CompilingMethod, token)
        compilerState = append(compilerState, CompilingMethod)
        compileMethod()
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

func compileFunction() {
    fmt.Printf("compiling function\n")
}

func compileMethod() {
    fmt.Printf("compiling method\n")
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
