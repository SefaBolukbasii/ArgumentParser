package argParser

import (
	"errors"
	"fmt"
	"os"
)

type ArgumentType string

const (
	CommandType ArgumentType = "command"
	OptionType  ArgumentType = "option"
)

type Argument struct {
	name         string
	description  string
	example      string
	defaultValue *string
	forced       bool
	argumentType ArgumentType
}

var ArgumentsArray []Argument

func AddArgument(com *Argument) {
	ArgumentsArray = append(ArgumentsArray, Argument{
		name:         com.name,
		description:  com.description,
		example:      com.example,
		defaultValue: com.defaultValue,
		forced:       com.forced,
		argumentType: com.argumentType,
	})
}

func Parse() (map[string]any, error) {
	args := os.Args[1:]
	parsedArgs := make(map[string]any)
	wrongCommand := false
	for _, Argument := range ArgumentsArray {
		for i := 0; i < len(args); i++ {
			if args[i] == "--help" || args[i] == "-h" {
				for _, ArgumentHelp := range ArgumentsArray {
					ArgumentHelp.Help()
				}
				return nil, nil
			} else {
				if args[i][2:] == Argument.name {
					if Argument.argumentType == CommandType {
						if i+1 < len(args) {
							if args[i+1][:1] == "-" {
								if Argument.defaultValue != nil {
									parsedArgs[Argument.name] = Argument.defaultValue
								} else {
									return nil, errors.New("yanlış kullanılmış")
								}
							} else {
								parsedArgs[Argument.name] = args[i+1]
							}
							i++
						} else {
							if Argument.defaultValue != nil {
								parsedArgs[Argument.name] = Argument.defaultValue
							} else {
								return nil, errors.New("yanlış kullanılmış")
							}
						}
					} else {
						parsedArgs[Argument.name] = true
					}
					wrongCommand = false
					break
				} else {
					wrongCommand = true
				}
			}
		}
		if wrongCommand && Argument.forced {
			return nil, errors.New("komut zorunlu olarak kullanılmalı")
		}
	}
	return parsedArgs, nil
}
func (arg *Argument) Help() {
	fmt.Println(arg.name, "\n", arg.description, "\n", arg.example)
}
