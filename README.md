# Agate ( *A*PI *Ga*teway *Te*mplate)

This is a [boilr](https://github.com/tmrts/boilr) template for scaffolding Go api server project.

## Stacks

- echo

### Optionals

- mysql
- mongo
- redis
- memcached

# Installation

## Install boilr on your machine

```
go get github.com/tmrts/boilr
```
 
## Register this template via boilr command
  
```
boilr template download github.com/katsew/agate agate
```

### Tips!

If you cannot download template,  
just `git clone github.com/katsew/agate`, `cd` to `agate`,      
then run `boilr template save . agate` .

## Usage

Just run

```
boilr template use agate <paht/to/your/repo>
```

Answer to the questions like, PkgName, Version, Author...  
and that's it!


## What's next?

After scaffolding your project, you may run `dep ensure` to install go dependencies via [dep](https://github.com/golang/dep).  
If you don't have dep installed, simply run `go get -u github.com/golang/dep/cmd/dep` to install it.  

### For docker user

This template includes Dockerfile for pkg, and docker-compose.yaml for local-machine development.  
This will make your dev life better.  

## License

MIT

Hope you enjoy :)
