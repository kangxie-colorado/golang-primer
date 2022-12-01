# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
# class NestedInteger:
#    def isInteger(self) -> bool:
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        """
#
#    def getInteger(self) -> int:
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        """
#
#    def getList(self) -> [NestedInteger]:
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        """

class NestedInteger:

    def isInteger(self) -> bool:
        """
        @return True if this NestedInteger holds a single integer, rather than a nested list.
        """
        ...

    def getInteger(self) -> int:
        """
        @return the single integer that this NestedInteger holds, if it holds a single integer
        Return None if this NestedInteger holds a nested list
        """
        ...

    def getList(self):
        """
        @return the nested list that this NestedInteger holds, if it holds a nested list
        Return None if this NestedInteger holds a single integer
        """
        ...


class NestedIterator:
    def __init__(self, nestedList):
        self.nestedList = nestedList  # a list of NestedInteger
        self.offset = 0
        self.nextLevelList = {}

    def next(self) -> int:
        iter = self.nestedList[self.offset]
        ret = None
        if iter.isInteger():
            ret = iter.getInteger()
            self.offset += 1
        else:
            if self.offset not in self.nextLevelList:
                self.nextLevelList[self.offset] = NestedIterator(
                    iter.getList())

            # if self.nextLevelList[self.offset].hasNext():
            ret = self.nextLevelList[self.offset].next()
            if not self.nextLevelList[self.offset].hasNext():
                self.offset += 1

        return ret

    def hasNext(self) -> bool:
        print(self.nestedList, self.offset)
        if self.offset < len(self.nestedList):
            iter = self.nestedList[self.offset]
            if iter.isInteger():
                print("Heer")
                return True
            else:
                if self.offset not in self.nextLevelList:
                    self.nextLevelList[self.offset] = NestedIterator(
                        iter.getList())
                if not self.nextLevelList[self.offset].hasNext():
                    self.offset += 1
                    return self.hasNext()

                return self.nextLevelList[self.offset].hasNext()
        else:
            return False


""""
                if not self.nextLevelList[self.offset].hasNext():
                    self.offset += 1
                    return self.hasNext()
                
                return self.nextLevelList[self.offset].hasNext()

            ^ see what the problem is
            ^ you call hasNext in two branches... and they will interleave calling themselves again
            ^ end up with a recursive call hell
            should change to 
                if not self.nextLevelList[self.offset].hasNext():
                    self.offset += 1
                    return self.hasNext()
                else:
                    self.nextLevelList[self.offset].hasNext()
            ^ painful lessons, but I will remember being careful around recursive calls
            in different branches... avoid calling them more than necessary

"""
