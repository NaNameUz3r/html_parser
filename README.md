# sitemap_gen

Sitemap Generator is a simple and lightweight MVP project that generates a site-map.xml file for a given website URL? implemented using Go and standard libraries.

### Features
- Crawls the given website URL and extracts all the links from the HTML pages
- Generates a sitemap.xml file that includes all the extracted links
- Lightweight and easy to use
- Gives you time to brew a cup of coffee thanks to unoptimized n*3 BFS

### Usage

To generate a sitemap for a website, simply run the main.go with the URL of the website as an argument:

run or build from source code

```go
git clone https://github.com/NaNameUz3r/sitemap_gen
cd sitemap_get && go run sitemap_gen.go --url https://example.com
```

The script will output a site-map.xml file in the current directory, by default. You can change this behavior with --out key.
Default depth search set to 10 links, you can change this with --depth key.

### Limitations

Sitemap Generator is a minimum viable product (MVP) proof-of-concept project, designed to demonstrate the basics of web crawling and sitemap generation. However, there are some limitations to be aware of:

- The program does not support websites with authentication or login requirements.
- The program does not handle JavaScript-based navigation or content loading.
- The program may encounter issues with websites that have complex URL structures or dynamically generated pages.
- The program may encounter rate limiting or other restrictions imposed by the website being crawled.

### Contributions

Contributions to this project are welcome. If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

### License

This project is licensed under the BSD 2-Clause License. See the LICENSE file for details.