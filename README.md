# `csv-viewer` 📑💻
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
Help: 
```bash
csv-viewer --help
-h, --help
-v, --view (column, table, json)
```

### Column view
Column view will display the content in a clean easy way to see the content:  
```bash
csv-viewer ./test/test.csv
```
or
```bash
csv-viewer --view column ./test/test.csv
```

Result: 
```
AAA      Description with a comma, and a quote "like this"     More text
BBB      Another description, with a newline like this         More text
```

> The spacing is similar to `column` gives you.

### Table view
```bash
csv-viewer --view table ./test/test.csv
```

Result:
```

```

### JSON view
```bash
csv-viewer --view json ./test/test.csv
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