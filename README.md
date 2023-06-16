# PTree

WARNING: this is a proof-of-concept, it cannot be used in any real code.

## Why?
This radix tree uses a bit unusual way to match strings. It is actually much faster, than ordinary ones. However, it cannot handle multiple keys, if one key is a prefix of another. So, to say, you can't store keys Hell and Hello at the same time. But Hello and Hallo are just fine.

## How?
To achieve better efficiency, we're using a slice of another slices, where nested are fix-sized (256 entries) int slices. The point is that each fix-sized slice corresponds to some matching branch. Each index in such a slice corresponds to an ASCII-letter, and every value - to a next "frame" (its index, to be precious). In case the value is 0 - the character doesn't match. In case value is negative - make it positive, subtract 1, and use as the index to the slice with values.

## Why not ... ?
### Why can't I use keys `Hell` and `Hello` at the same time?
As said before, each frame contains only an index to a next frame (or index to a value).
### Is it possible to fix it?
Yes. We can use int64, but treat as a pointer to a next frame only the first part of it. The second part can be left to point at the index of the value (if it is presented at the current point). This won't much affect either performance or capabilities, as bitmask is pretty cheap to apply, and 1<<31-1 is still a huge number of values and frames can be stored.
### Why the advice above isn't implemented yet?
Unfortunately, this project was just a fun experiment, I spent half an hour at. No more plans about it. In other words, it's just a demo of some curious way to implement a radix tree.
### Just in case, can I implement it by my own?
Everybody is just welcome to contribute, or ask further questions!
