/*
Robot in a Grid: Imagine a robot sitting on the upper left corner of grid with r rows and c columns.
The robot can only move in two directions, right and down, but certain cells are "off limits" such that
the robot cannot step on them. Design an algorithm to find a path for the robot from the top left to
the bottom right. 
*/
ArrayList<Point> getPath(boolean[)[) maze) {
	if (maze == null || maze. length == 0) return null;
 		ArrayList<Point> path = new ArrayList<Point>();
 		HashSet<Point> failedPoints = new HashSet<Point>();
 	if (getPath(maze, maze.length - l, maze[0].length - 1, path, failedPoint)) {
 		return path;
 	}
 	return null;
 }

boolean getPath(boolean[][] maze, int row, int col, ArrayList<Point> path, HashSet<Point> failedPoints) {
/* If out of bounds or not available, return. */
	if (col < 0 || row < 0 || !maze[row][col]) {
		return false;
	}

	Point p = new Point(row, col);

/* If we've already visited this cell, return. */
	if (failedPoints.contains(p)) {
		return false;
	}

	boolean isAtOrigin = (row == 0) && (col == 0);

/* If there's a path from start to my current location, add my 10cation. */
 	if (isAtOrigin || getpath(maze, row, col - 1, path, failedPoints) || getPath(maze, row - 1, col, path, failedPoints)) {
 		path.add(p);
 		return true;
 	}

	failedPoints.add(p); // Cache result
	return false;
 }