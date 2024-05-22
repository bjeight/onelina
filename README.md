# Onelina

Linearize a fasta file (put every sequence on a single line)

##  Installation

First, [install go](https://go.dev/dl/),

then:

```
go install github.com/bjeight/onelina@latest
```

or

```
git clone https://github.com/bjeight/onelina.git
cd onelina
go build -o onelina
```

##  Usage

```
onelina input.fasta > output.fasta
```
```
onelina input.fa.bgz | bgzip > output.fa.bgz
```
```
onelina input.fa.gz | gzip > output.fa.gz
```