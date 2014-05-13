go-jaro-winkler-distance
=====

Native [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) in Go. Makes heavy use of strings package, but single query doesn't take longer than about 30us.

The script has some inaccuracies between different implementations, so before using in critical applications please check it against your previous libary. For more information see algo.go line 56.