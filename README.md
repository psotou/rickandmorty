# Rick and Morty API

### Generals 

This project works both as a kind of SDK to consume the Rick And Morty API and as a standalone executable for thouse who aren't really into Go (although you totally should, the language is awesome). So if you aren't a gopher, go clone this repo

```bash
git clone https://github.com/psotou/rickandmortyapi.git
```

and then run

```bash
bin/rmc
```

This should produce a JSON file (results.json) in the same directory where you run the executable.

### Examples

To get a certain resource, you have to pass an array (slice) of strings with the `ids` that you want to retrieve. For example, if you want to get the resource Character for the characters with ids 1, 2 and 3, you can do:

```go
func main () {
    idsRange := []string{"1", "2", "3"}
    characters := getCharacters(idsRange)
    fmt.Println(characters)
}
```

If you want to get all the elements of a resource, do:

```go
func main () {
    maxCharacterId := getInfo(Character).Count
    idsRange := makeRange(1, maxCharacterId)
    characters := getCharacters(idsRange)
    fmt.Println(characters)
}
```

Additionally, there exists a `counChar` method for counting the number of ocurrences of a certain letter (case sensitive) within the `name` element of the resource:

```go
func main () {
    idsRange := []string{"1", "2", "3"}
    characters := getCharacters(idsRange)
    fmt.Println(characters.countChar("o"))
}
```