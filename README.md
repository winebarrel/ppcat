# ppcat

ppcat is a tool to format JSON in the middle of text.

## Usage

```
Usage: ppcat [OPTION] [FILE]
  -version
    	print version and exit
```

```
$ cat sample.txt
London Bridge is broken down,
{"foo":"bar","zoo":[100,200]}
Broken down, broken down.
[{"foo":"bar","zoo":[100,200]},{"zoo":{"baz":[300,400]}}]
London Bridge is broken down,
{"foo": // broken
My fair lady.

$ cat sample.txt | ppcat
London Bridge is broken down,
{
  "foo": "bar",
  "zoo": [
    100,
    200
  ]
}
Broken down, broken down.
[
  {
    "foo": "bar",
    "zoo": [
      100,
      200
    ]
  },
  {
    "zoo": {
      "baz": [
        300,
        400
      ]
    }
  }
]
London Bridge is broken down,
{"foo": // broken
My fair lady.
```
