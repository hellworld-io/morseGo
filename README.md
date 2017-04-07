# MorseGo

### `MorseGo` is written  by golang. It is a simple command line for converting Morse Code or text.

## What Morse Code?

### [Morse Code](https://en.wikipedia.org/wiki/Morse_code)

## How to use it

### This is able to convert from words to morse code, or convert from morse code to words.
```
./morse
Usage of morse:
    convert -morse string
        morse code to text ex) morse convert -morse '. .-  . .-'
    convert -text string
        Text to morse code ex) morse convert -text 'a b'
```

### There are results about running morse.
```
go run morse.go convert -text 'abc d' -morse '.- -... -.-.  -..'


converted text : .- -... -.-.  -..  
converted morse code : abc d
```


