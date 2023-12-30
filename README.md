# Butterbird-Go Discord Bot

## Description

Butterbird-Go is a Discord bot written in Go, designed to provide a range of functionalities to enhance your Discord server experience.

## Libraries Used

This project makes use of the following libraries:

1. **[go-openai](https://github.com/sashabaranov/go-openai)**: A Go client library for accessing the OpenAI API. This library is licensed under the Apache License 2.0. More details can be found in their [license file](https://github.com/sashabaranov/go-openai/blob/master/LICENSE).

2. **[discordgo](https://github.com/bwmarrin/discordgo)**: A Go package for creating Discord bots. Licensed under the BSD 3-Clause License. The full license text can be found in their [license file](https://github.com/bwmarrin/discordgo/blob/master/LICENSE).

## License

This project uses the [MIT License](<./LICENSE>)



## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of [Go](https://golang.org/dl/).
- You have a basic understanding of Go programming.
- You have a Discord account and have created a bot on the Discord Developer Portal.

## Installation

To install Butterbird-Go, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/chad-collins/butterbird-go.git
   cd butterbird-go
   ```

2. Install the required dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your `config.json` file based on the provided `config.sample.json`. You need to replace placeholder values with your actual Discord bot token and other necessary configuration values.


## Running the Bot

To run Butterbird-Go, use the following command from the root of the project:

```bash
go run main.go
```

This command will start the Discord bot and connect it to your server using the credentials provided in your configuration.

## Contributing to Butterbird-Go

To contribute to Butterbird-Go, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`.
4. Push to the original branch: `git push origin <project_name>/<location>`.
5. Create the pull request.

Alternatively, see the GitHub documentation on [creating a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## Contributors

Thanks to the following people who have contributed to this project:

- [@chad-collins](https://github.com/chad-collins)

## Contact

If you want to contact me, you can reach me at `<chad@chadcollins.net>`.



---

