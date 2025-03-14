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
	Name         string
	Description  string
	Example      string
	DefaultValue *string
	Forced       bool
	ArgumentType ArgumentType
}

var ArgumentsArray []Argument

func AddArgument(com *Argument) {
	ArgumentsArray = append(ArgumentsArray, Argument{
		Name:         com.Name,
		Description:  com.Description,
		Example:      com.Example,
		DefaultValue: com.DefaultValue,
		Forced:       com.Forced,
		ArgumentType: com.ArgumentType,
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
				if args[i][2:] == Argument.Name {
					if Argument.ArgumentType == CommandType {
						if i+1 < len(args) {
							if args[i+1][:1] == "-" {
								if Argument.DefaultValue != nil {
									parsedArgs[Argument.Name] = *Argument.DefaultValue
								} else {
									Argument.Help()
									return nil, nil
								}
							} else {
								parsedArgs[Argument.Name] = args[i+1]
							}
							i++
						} else {
							if Argument.DefaultValue != nil {
								parsedArgs[Argument.Name] = Argument.DefaultValue
							} else {
								Argument.Help()
								return nil, nil

							}
						}
					} else {
						parsedArgs[Argument.Name] = true
					}
					wrongCommand = false
					break
				} else {
					wrongCommand = true
				}
			}
		}
		if wrongCommand && Argument.Forced {
			return nil, errors.New("komut zorunlu olarak kullanılmalı")
		}
	}
	return parsedArgs, nil
}
func (arg *Argument) Help() {
	fmt.Println(arg.Name, "\n", arg.Description, "\n", arg.Example)
}
