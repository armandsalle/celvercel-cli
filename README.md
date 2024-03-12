# celvercel-cli

A simple CLI tool to create Oreo meme with the Vercel logo.

## What is Oreo meme?

Oreo meme is a meme where someone separates the 3 parts of an Oreo `O` (the top cookie), `Re` (the creme) and `O` (the bottom cookie) and stack them in a differents orders to make it `Re`, `O`, `O` (ReOreo) or `O`, `O`, `Re` (OORe) or whatever you want.

### Examples

`vercel`
```
   /\
  /  \
 /    \
/______\
```

`lecrev`
```
\¯¯¯¯¯¯/
 \    /
  \  /
   \/
```

`vercellecver`
```
   /\
  /  \
 /    \
/      \
\      /
 \____/
   /\
  /__\
```

## How to use it?

One command to rule them all. Just run the following command and follow the instructions.

```bash
npx celvercel
```

if you want to save the result in a file, just add the `--output` flag and the name of the file. The file will be saved in the current directory.
```bash
npx celvercel --output my-oreo-meme.txt
```

### How to build and release

I'm using [goreleaser](https://goreleaser.com/) to release the package. Documentation is [here](https://goreleaser.com/quick-start/).

```bash
git tag -a v0.0.1 -m "Message"
git push origin v0.0.1
goreleaser release --clean
```

### Todo

- [x] Create the project
- [x] Create the Oreo meme with Vercel logo
- [x] Add a LICENSE
- [x] Add Cobra to handle args and flags
- [x] Add tests
- [x] Add CI/CD
- [x] Publish the package to NPM