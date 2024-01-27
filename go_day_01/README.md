## Different formats.

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/go_day_01/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/go_day_01/internal/app/run)
- [Internal Package Read_DB](http://localhost:6060/pkg/go_day_01/internal/app/read_db)
- [Internal Package Compare_DB](http://localhost:6060/pkg/go_day_01/internal/app/compare_db)
- [Internal Package Compare_FS](http://localhost:6060/pkg/go_day_01/internal/app/compare_fs)
- [Internal Package My_Errors](http://localhost:6060/pkg/go_day_01/internal/app/my_errors)

### Reading
There are two implementations of the same interface, `DBReader`, one for reading JSON and one for reading XML. 
Both implementations return the same object types as a result. The JSON version of the database is printed 
when reading from XML, and vice versa.

**Usage:**
```bash
$ ./app -f original_database.xml
$ ./app -f stolen_database.json
```

### Comparison
A comparison is made between two databases. 
The comparison works with both formats (JSON and XML), reusing code from the read_db package. 
It is evident that one of the databases is a modified version of another, 
resulting in multiple possible scenarios:

1) The addition or removal of a cake.
2) The cooking time for the same cake has been modified.
3) There is a modification made to the recipe by either adding or subtracting an ingredient.
4) There has been a change in the count of units for the same ingredient.
5) There has been a change in the unit of measurement for the same ingredient.
6) There is a missing or an additional unit.

**Usage:**
```bash
$ ./app --old original_database.xml --new stolen_database.json
```

### Dumps-comparison
Dumps-comparison compares two files in .txt format and outputs the differences. 
To compare two files, only one of them is read.

**Usage:**
```bash
$ ./app --old snapshot1.txt --new snapshot2.txt
```
