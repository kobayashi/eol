# eol

`eol` is command line tool for [endoflife.date](https://endoflife.date/). This is mostly inspired from [norwegianblue](https://github.com/hugovk/norwegianblue) built in python.

## usage

- get list of projects

```
$ eol all
alpine
amazon-eks
amazon-linux
android
...
```

- get project matrix

Check EOL status for a project in the list.

```
$ eol project go
+-------+-----+------------+---------+
| CYCLE | EOL | RELEASE    | LATEST  |
+-------+-----+------------+---------+
| 1.17  | no  | 2021-08-16 | 1.17.6  |
| 1.16  | no  | 2021-02-16 | 1.16.13 |
| 1.15  | yes | 2020-08-11 | 1.15.15 |
| 1.14  | yes | 2020-02-25 | 1.14.15 |
| 1.13  | yes | 2019-09-03 | 1.13.15 |
| 1.12  | yes | 2019-02-25 | 1.12.17 |
| 1.11  | yes | 2018-08-04 | 1.11.13 |
| 1.10  | yes | 2018-02-16 | 1.10.8  |
+-------+-----+------------+---------+

```

- get a matrix for a any version of a project

```
$ eol project go 1.17
+------------+-----+--------+----------------+
| RELEASE    | EOL | LATEST | CYCLESHORTHAND |
+------------+-----+--------+----------------+
| 2021-08-16 | no  | 1.17.6 |            117 |
+------------+-----+--------+----------------+

```

## Custom output format

- markdown

```markdown
$ eol project go 1.17 -f markdown
| Release | EOL | Latest | CycleShortHand |
| --- | --- | --- | ---:|
| 2021-08-16 | no | 1.17.6 | 117 |
```

- csv

```csv
eol project go 1.17 -f csv
Release,EOL,Latest,CycleShortHand
2021-08-16,no,1.17.6,117
```

- html

```html
eol project go 1.17 -f html
<table class="go-pretty-table">
  <thead>
  <tr>
    <th>Release</th>
    <th>EOL</th>
    <th>Latest</th>
    <th align="right">CycleShortHand</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td>2021-08-16</td>
    <td>no</td>
    <td>1.17.6</td>
    <td align="right">117</td>
  </tr>
  </tbody>
</table>

```
