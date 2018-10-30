# Contributing
All contributions are welcomed!

## Opening Issues
[Open an issue][create-issue-url] on the GitHub repository

## Developing
All that is needed to work with this repo is [Go][go-download-url] and your favorite editor or IDE, although we recommend [VS Code][vscode-url].

### Setup
To set up your local environment to work on this project:

Clone the repo, change into the directory where you cloned the repo, and then run the bootstrap script
```sh     
git clone https://github.com/swellaby/ci-detective.git
cd ci-detective 
go run ./scripts/dev_setup.go
```

### Testing
More details coming soon... Run:
```sh
task test
```

### Linting
More details coming soon... Run:
```sh
task lint
```

### Submitting changes
Swellaby members should create a branch within the repository, make changes there, and then submit a PR. 

Outside contributors should fork the repository, make changes in the fork, and then submit a PR.


[Back to Top][top]

[create-bug-url]: https://github.com/swellaby/captain-githook/issues/new?template=BUG_TEMPLATE.md&labels=bug&title=Bug:%20
[create-question-url]: https://github.com/swellaby/captain-githook/issues/new?template=QUESTION_TEMPLATE.md&labels=question&title=Q:%20
[create-enhancement-url]: https://github.com/swellaby/captain-githook/issues/new?template=ENHANCEMENT_TEMPLATE.md&labels=enhancement
[go-download-url]: https://golang.org/dl/
[vscode-url]: https://code.visualstudio.com/
[top]: CONTRIBUTING.md#contributing