# thecodelessctl
## What is this?
A CLI reader for [The Codeless Code :: Fables and Koans for the Software Engineer](http://thecodelesscode.com) by Qi.

This project is created under a [Creative Commons Attribution-NonCommercial 3.0 Unported License](Creative Commons Attribution-NonCommercial 3.0 Unported License). Original content is unchanged other than formatting and omitting images that won't render in a shell.

Give these a little read while you are waiting for your code to compile, your Terraform to plan, or your yaml to  configure something :)

## Installation
```
brew tap jasonblanchard/homebrew-thecodelessctl
brew install thecodelessctl
```

```
thecodelessctl read
```

Initialize a config file to store bookmark state:

```
thecodelessctl init
```

Then read through each story using your bookmark state by calling `next`:

```
thecodelessctl next
```

Pipe story output to `less` for a nicer reading experience:

```
thecodelessctl read 42 | less
```