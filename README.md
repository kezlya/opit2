# T.e.t.r.i.s simulator

This is an experemental project for practicing algorithms writing in Go. Testing performance and eventially impliment machine learning. The idea is to write a "bot" that playing t.e.t.r.i.s by communication with simulation framework.

## Simulation framework
In order to see how the "bot" performs in competitio, simple simulation framework has been written. Here is few available commands:

### Play one game
Visual representation in a terminal of the "bot" playing preset game. Applying different strategy you can see different behavior and results
	go test -run=one

### Play many games
Visual representation in default browser (a graph) of "bot" playing many preset games using current strategy.
	go test -run=many

## Testing
Usual Go testing. Just execute this command:
	go test