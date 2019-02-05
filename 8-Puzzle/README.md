`8 Puzzle Solver`

The puzzle can be solved by moving the tiles one by one in the single empty space and thus achieving the Goal configuration

Eg. For puzzle configuration -
```go
startState := [][]int {
		{3,0,7},
		{2,8,1},
		{6,4,5},
	}
```
We can achieve the goal state -
```go
goalState 	:= [][]int {
		{3,8,7},
		{2,0,1},
		{6,4,5},
	}
```
the following ways  -

            		3 0 7
            		2 8 1
            		6 4 5
            		/ | \
            	       /  |  \
            	      /   |   \
            	 0 3 7  3 7 0  3 8 7
                     2 8 1  2 8 1  2 0 1    <= Level 1
            	 6 4 5  6 4 5  6 4 5
            	/  |    ...  \    ...
                   /   |          \ 
                  /    |           \
               3 0 7  2 3 7       3 0 7
               2 8 1  0 8 1       2 8 1     <= Level 2
               6 4 5  6 4 5       6 4 5
               / ...   ...        ... \
             3 8 7                  3 8 7
             2 0 1                  2 0 1   <= Level 3
             6 4 5                  6 4 5
             

*By specifying the depth of the algorithm to 3, 3 solutions can be found for the given initial configuration*
