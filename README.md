# Dependency Checker

Given an input file in the format:

```
a: b c d
c: e f
d: f
```

The program will work out all the dependencies related to a specific node, for example:

- all the dependencies for `a` are: `b, c, d, f` (`a` depends on `d` which depends on `f`)
- all the dependencies for `c` are: `e, f`
- all the dependencies for `d` are: `f`

To run simply:

```
go run main.go input searchNode
```

With the provided example:

```
go run main.go input.txt a.java
[a.class a.jar]
```
