# `csv-viewer`
CSV viewer offers different ways to display in your terminal a CSV content. From JSON, table to simple column view.

## Install
```bash
sudo make install
```

## Uninstall
```bash
sudo make uninstall
```

## How-to-Use
### Column view
Column view will display the content in a clean easy way to see the content:  
```bash
csv-viewer ./test/test.csv
```

Result: 
```
AAA      Description with a comma, and a quote "like this"     More text
BBB      Another description, with a newline like this         More text
```

> The spacing is similar to `column` gives you.

### JSON view
```bash
csv-viewer --json ./test/test.csv
```

Result:
```
[
    [
        "AAA",
        "Description with a comma, and a quote \"like this\"",
        "More text"
    ],
    [
        "AAA",
        "Another description, with a newline like this",
        "More text"
    ]
]
```

### Table view
```bash
csv-viewer --table ./test/test.csv
```

Result:
```

```

### Selecting columns to display
```bash
csv-viewer --columns="$1,$3" ./test/test.csv
```

Result:
```
AAA    More text
BBB    More text
```

> This does not affect the usage on view options.