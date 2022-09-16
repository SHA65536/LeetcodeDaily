class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
# https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
# Given two binary trees original and cloned and given a reference to a node target in the original tree.
#
# The cloned tree is a copy of the original tree.
#
# Return a reference to the same node in the cloned tree.
#
# Note that you are not allowed to change any of the two trees or the target node and the answer must be a reference to a node in the cloned tree.

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        if cloned.val == target.val:
            return cloned
        if cloned.left:
            val = self.getTargetCopy(original, cloned.left, target)
            if val : return val
        if cloned.right:
            val = self.getTargetCopy(original, cloned.right, target)
            if val : return val
        return None