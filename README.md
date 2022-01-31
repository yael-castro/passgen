```
     ____                      _____   _____  _     __
    /    \ ____   ___ ___     //  //  / ___/ / \   / /
   /  ___//   /| //_ //_  == //  __  / __/  / / \ / /  Password Generator  
  /__/    \___\|__//__//    //___// /____/ /_/   \_/

```
# **`PassGEN`** - generate your own word dictionaries based on a character sets

Command Line Interface (CLI) that can be used to make a word dictionaries.

Made by [Yael](github.com/yael-castro)

###### How to use
```
  passgen [options]
```

###### Options
```
  -length uint
        use to set length for words (this flag will ignore if the flag 'state' is present)
  -out string
        use to set name to the output file (default "{unix time}.txt")
  -state value
        password generation state, must be an array (the state is used to continue with a password generation previously canceled)
  -string string
        use to set characters used to generate values
  -v    verbose
```
