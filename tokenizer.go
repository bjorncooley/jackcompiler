package main

import (
    "io/ioutil"
    "log"
    "text/scanner"
    "strconv"
    "strings"
)

type TokenType string
const (
    Keyword TokenType = "keyword"
    Symbol TokenType = "symbol"
    IntegerConstant TokenType = "integerConstant"
    StringConstant TokenType = "stringConstant"
    Identifier TokenType = "identifier"
)

type Token struct {
    tokenType TokenType
    lexeme string
}

var tokenList []Token

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func tokenize(filePath string) {
    sourceString := getSourceAsString(filePath)
    readTokens(sourceString)
}

func getSourceAsString(filePath string) string {
    data, err := ioutil.ReadFile(filePath)
    check(err)
    return string(data)
}

func readTokens(sourceString string) {
    var s scanner.Scanner;
    s.Init(strings.NewReader(sourceString))
    compilationEngine := new(CompilationEngine)

    for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
        text := s.TokenText()
        tokenType := getTokenType(text)
        t := Token{tokenType: tokenType, lexeme: text}
        tokenList = append(tokenList, t)
    }

    compilationEngine.Compile(tokenList)
}

func getTokenType(token string) TokenType {
    if KeywordMap[token] {
        return Keyword
    } else if SymbolMap[token] {
        return Symbol
    } else if isInteger(token) {
        return IntegerConstant
    } else if isString(token) {
        return StringConstant
    }
    return Identifier
}

func isInteger(token string) bool {
    if _, err := strconv.Atoi(token); err == nil {
        return true
    }
    return false
}

func isString(token string) bool {
    if len(token) <= 2 {
        return false
    }
    firstChar := token[0:1]
    lastChar := token[len(token)-1:]

    if (firstChar == `"` && lastChar == `"`) {
        return true
    }
    if (firstChar == "'" && lastChar == "'") {
        return true
    }
    return false
}
