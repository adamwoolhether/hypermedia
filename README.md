https://hypermedia.systems/

Building hypermedia apps with HTMX, HXML, Go and Templ.

# Run the web example
Navigate to http://localhost:42069/
```shell
make up
```

## Dev
### Install Templ and WGO command line tools.
```shell
make dev.setup
```
### Run wgo with templ file watching to auto-reload changes
```shell
make dev
```

Will update this readme in the near future.

## TODO enhancements/issue:
- Better flash handling: currently just stuffing into context?
- Smooth progress bar transition, needed to use JS.
- Web cannot delete single contact with checkbox (alpine / sweet confirm)