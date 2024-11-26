# ascii-art-web

## Description

`ascii-art-web` is a web GUI version of our [`ascii-art`](https://01.gritlab.ax/git/isaavuor/ascii-art.git) program. Our webpage now allows the use of three different banners: shadow, standard, and thinkertoy.

## Authors

- Inka Säävuori 👑
- [Jedi Reston](https://github.com/jeeeeedi)
- [Aung Khant Min (Richard)](https://github.com/Richard-AungKhantMin)

## Usage: how to run

To run our `ascii-art-web` program, write this in your terminal:

```bash
go run main.go
```

Then on your browser, go to `http://localhost:8080/` or `http://localhost:8080/ascii-art`. This is where you can input your text and select your banner style.

## Implementation details: algorithm

In the `main` function, `http.HandleFunc` handles both GET and POST requests only via the `/` and `/ascii-art` URLs. Finally, the server is started on `http://localhost:8080`.

The `indexHandlerFunc` function loads the HTML template (`index.html`) using `tmpl.ParseFiles()`. Then it handles both GET and POST requests only. The GET request only loads and renders the input form using the HTML template. The POST request reads the user's input text (`textInputName`) and the selected style (`bannerName`). Then it calls `AsciiArt()` to convert the input text to ASCII art.

### Error handling

- 400: Text input is empty.
- 404: URL not found.
- 405: Method not allowed.
- 500: Internal server errors during ASCII art processing.
