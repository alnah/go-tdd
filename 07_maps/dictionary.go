// Package dictionary provides a simple dictionary implementation with 
// functionalities include adding, searching, updating, and deleting words 
// along with their definitions. It also defines custom error types for 
// handling errors related to dictionary operations.
package dictionary

// ErrNotFound is returned when a word cannot be found in the dictionary.
const (
	ErrNotFound         = Err("could not find the word you were looking for")
	ErrWordExists       = Err("cannot add word because it already exists")
	ErrWordDoesNotExist = Err("cannot update word because it does not exist")
)

// Err represents an error related to the dictionary operations.
type Err string

// Error implements the error interface for DictionaryErr.
func (e Err) Error() string {
	return string(e)
}

// Dictionary is a type that maps words to their definitions.
type Dictionary map[string]string

// Search looks up a word in the dictionary and returns its definition or an error.
func (d Dictionary) Search(word string) (string, error) {
	// Short variable declaration; maps return two values when accessed.
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a new word and its definition into the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	// A switch statement provides extra safety in case Search returns
	// an error other than ErrNotFound.
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update modifies the definition of an existing word in the dictionary.
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

// Delete removes a word from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Sidenote: a map value is a pointer to a runtime.hmap structure.
// In other words, when you pass a map to a function/method, you are indeed
// copying it, but just the pointer part, not the underlying data structure
// that contains the data. A gotcha with maps is that they can be a nil value.
// A nil map behaves like an empty map when reading, but will cause a runtim
// panic error when  writing. Therefore, never initialize a map like this:
//		var m map[string]string
// But initialize it as an empty map or with the make keyword:
//		var m = map[string]string{}
//		var m = make(map[string]string)
// This ensures you will never get a runtime panic error.
