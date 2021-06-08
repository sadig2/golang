from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


n1 = ListNode(2)
n2 = ListNode(4)
n3 = ListNode(3)

n1.next = n2
n2.next = n3


m1 = ListNode(5)
m2 = ListNode(6)
m3 = ListNode(4)
m4 = ListNode(8)

m1.next = m2
m2.next = m3
m3.next = m4


def addTwoNumbers(l1, l2):
    a, b = l1, l2
    l = ListNode()
    cur = l
    car = 0

    while a != None or b != None:

        if a is None:
            x = 0
        else:
            x = a.val
        if b is None:
            y = 0
        else:
            y = b.val

        sum = x + y + car
        car = sum // 10
        digit = sum % 10

        cur.next = ListNode(digit)
        cur = cur.next

        if a is not None:
            a = a.next

        if b is not None:
            b = b.next

    return l.next


result = addTwoNumbers(n1, m1)

while result is not None:
    print(result.val)
    result = result.next
