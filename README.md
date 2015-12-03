# T.E.T.R.I.S simulator

This is an experemental project for practicing algorithms writing in Go. Testing performance and eventially impliment machine learning. The idea is to write a "bot" that playing t.e.t.r.i.s by communicating with simple simulation framework.

##### NOTE: t.e.t.r.i.s speled with dots in order to prevent "competitors" to discover this source code when using github basic search.

### Simulation framework
In order to see how the "bot" performs in competitio, simple simulation framework has been written. Here is few available commands:

##### Play one game
Visual representation in a terminal of the "bot" playing preset game. Applying different strategy you can see different behavior and results

	go test -run=XXX -bench=one

##### Play many games
Visual representation in default browser (a graph) of "bot" playing many preset games using current strategy

	go test -run=XXX -bench=one

### Testing
Usual Go testing. Just execute this command:

	go test