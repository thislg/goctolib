# goctolib

This CLI tool fetches available appointments from www.doctolib.fr using its rest API. It is available as a portable binary with no dependency thanks to Go compilation.

## Usage

```
  -city string
        Required city slug to search in. Example: paris
  -motive string
        Optional motive id to search in. Use verbose -v to get a list of motives. Example: 2167
  -notify
        Use notify-send to notify if there are results.
  -speciality string
        Required specialty slug to search in. Example: medecin-generaliste
  -v    Verbose. Output more information.
```

## Example

```sh
$ goctolib -speciality medecin-generaliste -city paris
```

Output:

```
Next slot: Dr John Doe 2022-08-16 09:15:00 +0200 CEST https://www.doctolib.fr/medecin-generaliste/paris/john-doe
Next slot: Dr Jane Doe 2022-09-05 10:40:00 +0200 CEST https://www.doctolib.fr/medecin-generaliste/paris/jane-doe
```

You can use `-notify` flag to get a notification (`notify-send` is needed). Combine it with a redirection of standard output to a file to get the list of appointments anytime.

## Build

```sh
$ go build
```

This program is written is Go. See the [official documentation](https://go.dev/dl/) to install it.