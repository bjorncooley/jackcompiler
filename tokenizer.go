package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "text/scanner"
    "strconv"
    "strings"
    // "unicode/utf8"
)

type TokenType string

const (
    Keyword TokenType = "keyword"
    Symbol TokenType = "symbol"
    IntegerConstant TokenType = "integerConstant"
    StringConstant TokenType = "stringConstant"
    Identifier TokenType = "identifier"
)

var keywordMap = map[string]bool {
    "boolean": true,
    "char": true,
    "class": true, 
    "constructor": true, 
    "do": true,
    "else": true,
    "false": true,
    "field": true, 
    "function": true, 
    "if": true,
    "int": true, 
    "let": true,
    "method": true, 
    "null": true,
    "return": true,
    "static": true, 
    "this": true,
    "true": true,
    "var": true, 
    "void": true,
    "while": true,
}

var symbolMap = map[string]bool {
    "{": true,
    "}": true,
    "(": true,
    ")": true,
    "[": true,
    "]": true,
    ".": true,
    ",": true,
    ";": true,
    "+": true,
    "_": true,
    "*": true,
    "/": true,
    "&": true,
    "|": true,
    "<": true,
    ">": true,
    "=": true,
    "-": true,
}



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
    for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
        text := s.TokenText()
        fmt.Printf("%s: %s\n", getTokenType(text), text)
    }
    fmt.Printf("\n")
}

func getTokenType(token string) TokenType {
    if keywordMap[token] {
        return Keyword;
    } else if symbolMap[token] {
        return Symbol;
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
