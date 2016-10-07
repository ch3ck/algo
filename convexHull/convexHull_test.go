package convexHull

import "testing"

func TestConvexHull(t *testing.T) {
	var points = Points{{0, 3}, {2, 2}, {1, 1}, {2, 1},
		{3, 0}, {0, 0}, {3, 3}}
	var expected = Points{{0, 0}, {3, 0}, {3, 3}, {0, 3}, {0, 0}}
	var result = findConvexHull(points)

	if result == nil {
		t.Error("returned convex hull is nil")
		return
	}

	if len(result) != len(expected) {
		t.Error("returned convex hull has a different size")
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Error("returned convex hull is wrong")
			return
		}
	}
}
