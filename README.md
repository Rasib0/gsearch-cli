# gsearch-cli

`gsearch-cli` is a command-line tool for searching the web using various search engines directly from your terminal.

## Usage

- To perform a search using the default search engine (Google):

   ```
   gs <query>
   ```

- To specify a search engine using a flag:

   ```
   gs <flag> <query>
   ```

Available search engine flags:

- `!g` - Google
- `!y` - YouTube
- `!d` - DuckDuckGo
- `!p` - Phind

To display the help message with available search engine flags, run:

```
gs -h
```

To check the version of `gsearch-cli`, use:

```
gs --version
```

## Installation

1. **Clone the repository:**

   ```
   git clone https://github.com/Rasib0/gsearch-cli.git
   ```

2. **Build the executable:**

   ```
   go build -o gs
   ```

3. **Move the executable to a directory in your PATH (optional):**

   ```
   mv gs /usr/local/bin/
   ```

## Contributing

Contributions to `gsearch-cli` are welcome! If you have any improvements, bug fixes, or new features to add, please open an issue or submit a pull request on [GitHub](https://github.com/Rasib0/gsearch-cli).

## License

This project is licensed under the Apache License 2.0
