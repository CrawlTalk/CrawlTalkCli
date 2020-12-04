# CrawlTalkCli

Golang command-line client compatible with [MoreliaTalk] server and protocol.

Русская версия доступна здесь - [Описание на русском]

## Release

## Command-line options

  -email string
  
        Default email. For registration only. If specified client will not request it interactive.
        
  -login string
  
        Default login. If specified client will not request it interactive.
        
  -no-color
  
        Disable color output
        
  -password string
  
        Default password. If specified client will not request it interactive.
        
  -port int
  
        Default port. If specified client will not request it interactive.
        
  -register
  
        Register new account
        
  -schema string
  
        Default schema (ws or wss). If specified client will not request it interactive.

  -server string
  
        Default server. If specified client will not request it interactive.

  -sign-in
  
        Sign in to existed account
        
  -username string
  
        Default username. For registration only. If specified client will not request it interactive.

## Available commands in interactive mode:
* /help - this help page
## Flow list mode:
* /exit - exit from program to command prompt
* /create - create new flow
## In flow mode:
* /exit - exit to flow list
* /inception - show all flow messages from beginning


[MoreliaTalk]: https://github.com/MoreliaTalk
[Описание на русском]: README_RU.md
