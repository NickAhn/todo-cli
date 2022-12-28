# TODO-CLI
A simple and effective TODO list CLI for personal use written with Go and Cobra. (WIP)

## Motivation
I have been looking for the perfect TODO app for a long time, but they all seem to fall short of what I want or need. So I decided to code myself a CLI one that is customized to my preferences (that's the beauty of coding!). This also gives me the opportunity to improve my skills in Go and test their renowed CLI capabilities.

## Features
* The todo list is displayed in a ascending, zero-indexed list, displaying the number of tasks todo.
* Each item has a description followed by its priority (low number, higher priority)
```
TODO (2):
  0. Write a README (p1)
  1. Buy Groceries (p4)
```


### Commands
Add one or multiple items with different descriptions, priorities, and due dates
```
> todo add "Write a README.md"
> todo add "Sleep" "Repeat" -p2     // priority 2
```

Delete one or multiple items
```
> todo del 0      
> todo del 3 5 2
```

Edit tasks
```
> todo edit 1 "Do something else"
> todo editp 1 2      // change priority of item 1 to p2
```

## Built With
* [Go](https://go.dev/) - main programming language for the CLI <br> I chose Go to improve my skills on the language
* [Cobra](https://github.com/spf13/cobra) - CLI Command Framework <br>
I chose to learn this framework because it is used by big projects such as Kubernetes, Docker, Dropbox, etc...
 
