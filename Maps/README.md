# Maps

In Go language, a map is a built-in data structure that associates keys with values. It is somewhat similar to a dictionary in Python or an object in JavaScript, but with some key differences. Here are the key points about maps in Go:

1. Declaration :

   ```go
   var m map[string]int
   ```

2. Initialization : There are two ways to initialize a map

   - Using make function :

   ```go
   m = make(map[string]int)
   ```

   - using map literals :

   ```go
   m = map[string]int{"apple" : 1,"banana" : 2}
   ```

3. Adding and updating elements : you can add and update the map using key notation

   ```go
   m["orange"] = 7
   ```

4. Deleting element from map : you can delete element from map using delete keyword

   ```go
   delete(m,"banana")
   ```

5. checking if key exists

   ```go
   _,exists := m["strawberry"]
   if exists {
       fmt.Println("key exists")
   } else{
       fmt.Println("key does not exists")
   }
   ```

6. Iterating over an map

   ```go
   for key,value := range m{
       fmt.Println(key, "->", value)
   }
   ```
