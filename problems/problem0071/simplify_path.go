package problem0071

import (
	"fmt"
	"strings"
)

/*
Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system,
convert it to the simplified canonical path.
In a Unix-style file system, a period '.' refers to the current directory,
a double period '..' refers to the directory up a level, and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'.
For this problem, any other format of periods such as '...' are treated as file/directory names.
The canonical path should have the following format:
    The path starts with a single slash '/'.
    Any two directories are separated by a single slash '/'.
    The path does not end with a trailing '/'.
    The path only contains the directories on the path from the root directory to the target file or directory
    (i.e., no period '.' or double period '..')
Return the simplified canonical path.
*/

func simplifyPath(path string) string {
	var words = strings.Split(path, "/")
	var res = make([]string, len(words))
	var rCur, wCur int
	for rCur = 0; rCur < len(words); rCur++ {
		switch words[rCur] {
		case "", ".":
			continue
		case "..":
			wCur = max(0, wCur-1)
		default:
			res[wCur] = words[rCur]
			wCur++
		}
	}
	return fmt.Sprintf("/%s", strings.Join(res[:wCur], "/"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
