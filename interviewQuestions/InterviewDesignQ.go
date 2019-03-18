/*
You can use any language, and I do not expect the solution to run or even compile in our 45 minute conversation.

Your task is to implement a PriorityExpiryCache cache with a max capacity.  Specifically please fill out the data structures on the PriorityExpiryCache object and implement the entry eviction method.

You do NOT need to implement the get or set methods.

It should support these operations:
  Get: Get the value of the key if the key exists in the cache and is not expired.
  Set: Update or insert the value of the key with a priority value and expiretime.
    Set should never ever allow more items than maxItems to be in the cache.
    When evicting we need to evict the lowest priority item(s) which are least recently used.

Example:
*/
p5 => priority 5
e10 => expires at 10 seconds since epoch

c = New(5)
c.Set("A", 1, priority=5,  expireTime=100)
c.Set("B", 2, priority=15, expireTime=3)
c.Set("C", 3, priority=5,  expireTime=10)
c.Set("D", 4, priority=1,  expireTime=15)
c.Set("E", 5, priority=5,  expireTime=150)
c.Get("C")

// 10+5+5 = 20
// Current time = 0
c.SetMaxItems(5)
c.Keys() = ["A", "B", "C", "D", "E"]
// space for 5 keys, all 5 items are included

time.Sleep(5)

// Current time = 5
c.SetMaxItems(4)
c.Keys() = ["A", "C", "D", "E"]
// "B" is removed because it is expired.  e3 < e5

c.SetMaxItems(3)
c.Keys() = ["A", "C", "E"]
// "D" is removed because it the lowest priority
// D's expire time is irellevant.

c.SetMaxItems(2)
c.Keys() = ["C", "E"]
// "A" is removed because it is least recently used."
// A's expire time is irellevant.

c.SetMaxItems(1)
c.Keys() = ["C"]
// "E" is removed because C is more recently used (due to the Get("C") event).
