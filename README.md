# TODO-CLI
A simple and effective TODO list CLI for personal use written with Go and Cobra. (WIP)

## Motivation
I have been looking for the TODO app for me, but they all seem to fall short of what I want or need. So I decided to code myself one that is customized to my preferences (that's the beauty of coding!). This also gives me the opportunity to improve my skills in Go and test their renowed CLI capabilities.

## Features
The todo list is displayed in a ascending, zero-indexed list, displaying the number of tasks todo.
```
TODO (2):
  0. Write a README
  1. Buy Groceries
```


### Commands
Add one or multiple items with different descriptions, priorities, and due dates
```
> todo add "Write a README.md"
> todo add "Eat" "Sleep" "Repeat"
```

Delete items by index (flags can be used for different indexing)
```
> todo del 0
> todo del 0 1 3
```

Edit tasks
```
```

## Built With
* [Go](https://go.dev/) - General Purpose Language. I chose Go because I've been enjoying learning it a lot, and wanted to improve my skills on the language.
* [Cobra](https://github.com/spf13/cobra) - CLI Command Framework. I chose to learn this framework because it is used by big projects such as Kubernetes, Docker, Dropbox, etc...
 