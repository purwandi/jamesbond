# James Bond

Sample application wrapper go using obfuscate binary.

## Purpose

When you need to run application that must consume super secret but don't want use
secret provider like vault.

- The `cmeder` folder consist simple go exec command logic that consist the secret, when 
  build pipeline run we obfuscate this program.
- The `rest` folder consist app that you want run and consume secret from `cmder`

## Dependencies

We are using garble to obfuscate wrapper binary

```
go install mvdan.cc/garble@latest
```

## Build

To build this sample apps, just using make command

```
make builder
```

## Run

Just run command below:

```
make run
```
