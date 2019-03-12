/*
Eight houses, represented as cells, are arranged in a straight line. Each day every cell competes with its adjacent cells. An integer value of 1 represents an active cell and a value of 0 represents an inactive cell. If the neighbors on both the sides of a cell are either active or inactive cell. If the neighbors are on both the sides of a cell are either active or inactive, the cell becomes inactive on the next day; otherwise the cell becomes active. The two cells on each end have a single adjacent cell, so assume that the unoccupied space on the opposite side is an inactive cell. Even after updating the cell state, consider its previous state when updating the state of other cell. The state information of all cells should be updated simultaneously.
write an algorithm to output the state of other cells. The state information of all cells should be updated simultaneously.
Wrote an algorithm to output the state of the cells after the given number of days .
input
The input to the function consists of two arguments;
states, a list of integers representing the current state of cells,
days, an integer representing the number of days.

Output 
Return a list of integers representing the state of the cells after the given number of days
*/

// IMPORT LIBRARY PACKAGES NEEDED BY YOUR PROGRAM
import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;
// SOME CLASSES WITHIN A PACKAGE MAY BE RESTRICTED
// DEFINE ANY CLASS AND METHOD NEEDED
// CLASS BEGINS, THIS CLASS IS REQUIRED
// public class Solution
// {        
//   // METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
//     public List<Integer> cellCompete(int[] states, int days)
//     {
//     // WRITE YOUR CODE HERE
//     }
//   // METHOD SIGNATURE ENDS
// }

public class Solution
{        
  // METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
    public List<Integer> cellCompete(int[] states, int days)
    {
    // WRITE YOUR CODE HERE
        List<Integer> tmpStates = new ArrayList<Integer>();
        for (int state : states)
        {
            tmpStates.add(state);
        }
        int j = 0;
        while(j<days) {
            List<Integer> tmpList = new ArrayList<Integer>();
            for(int i = 0; i < tmpStates.size(); i++) {
                if (i == 0) {
                    if (tmpStates.get(i+1) == 0) {
                        tmpList.add(0);
                    } else {
                        tmpList.add(1);
                    }
                } else if (i == states.length-1) {
                    if (tmpStates.get(i-1) == 0) {
                        tmpList.add(0);
                    } else {
                        tmpList.add(1);
                    }
                } else {
                    if (tmpStates.get(i-1) == tmpStates.get(i+1)) {
                        tmpList.add(0);
                    } else {
                        tmpList.add(1);
                    }
                }
            }
            tmpStates = new ArrayList(tmpList);
            j++;
        }
        return tmpStates;
    }
  // METHOD SIGNATURE ENDS
}