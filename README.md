<div id="header">
    <p align="center">
      <b>Gods 🏗️</b><br>
  	  <span font-size="16px">Generic Data structures for Go </span><br>
      <span font-size="12px">Made by <a href="http://epmor.app">Ethan</a> with love.</span><br><br>
      <span><a href="https://www.twitter.com/glassofethanol">Twitter</a> | <a href="https://github.com/sponsors/gweithio">Support</a></span><br><br>
    </p>
</div>

### About

Basic single function to split strings based off digits and case

### Usage

#### Sets
```go
	stringSet := gods.MakeSet[string]()
	stringSet.Add("Hello")
	stringSet.Add("World")
	stringSet.Add("Again")

	if !stringSet.Contains("Hello") {
		fmt.Println("Set does not contain Hello")
	}

	stringSet.Remove("World")

	if stringSet.Contains("World") {
		fmt.Println("World sstill exists when it should be removed")
	}
```


#### Queue
```go
	floatQueue := gods.MakeQueue[float32](5)

	floatQueue.Enqueue(1.1)
	floatQueue.Enqueue(2.2)
	floatQueue.Enqueue(3.4)

	value, ok := floatQueue.Dequeue()

	fmt.Println("Last item was %q", value) => 3.4

	if !ok {
		fmt.Println("The queue is already empty")
	}

```

None yet

