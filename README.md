# GetGoing-introduction-to-GoLang
leacture exercises for GoLang introductory course

### running files that import stuff from other files
 - #### CLI-way
   
   `go run my_other_file_1.go my_other_file_2.go file_that_i_want_to_run.go`
 - #### GoLand way
   (the hacky / tacky way that i know of)
 
   - Put all such files in single directly
   - Run | Edit Configurations...
   - Under Configurations tab, set
     - Run kind: Directory
     - Directory: _select the directory_
