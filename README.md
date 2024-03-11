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

```bash
celvercel-cli --output result.txt

prompt: celvercel

# result.txt

 /¯¯¯¯\
/______\
   /\
  /  \
 /    \
/______\

```

### Todo

- [x] Create the project
- [x] Create the Oreo meme with Vercel logo
- [x] Add a LICENSE
- [x] Add Cobra to handle args and flags
- [x] Add tests
- [ ] Add CI/CD
- [ ] Publish the package to Homebrew and APT