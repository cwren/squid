https://plus.google.com/+ChrisWren/posts/RLh4ZjD75Qe

I enjoyed reading +Yonatan Zunger's
[post](https://plus.google.com/+YonatanZunger/posts/4urrnW3sZsi) about
the
[jellyfish](http://wavegrower.tumblr.com/post/126854522925/currents-if-i-had-the-time-i-would-check-if-one-of)
(they look like squid to me) and yet I was left with a desire to know
how long it would actually take for the squid to all return back to
the place they started.

If we number the squid 0-255 then we can describe the state of the
image with a 256 dimensional vector.  The squid start out with number
0 at the top and number 255 at the bottom.  Then the moves can be
described by 256x256 permutation matrices.

The identity matrix is the permutation matrix that leaves everything
right where it was.  In the images below that is represented by a
solid black line from upper left to lower right (element 0 goes to 0,
element 1 goes to 1, etc).

The animation is made up of three moves repeated over and over.  In
these permutations squid either move left and right within a row (near
the identity diagonal, but just off it) or up and down a row (which
appear as the diagonal lines away from the center line).

As the moves continue the squid move farther from home. In the
animation below the intermediate frames have squid scattered all over
the field, much as it feels when you try to watch them.

Repeating the moves is the same as multiplying these permutation
matrices. Asking "how long will it take for the squid to return home"
is the same as asking "how many times do I need to multiply the matrix
by itself before I get the identity matrix back".

If the squid were only moving in one fixed pattern, it's clear that it
would take 256 steps: each squid visits the location of each other
squid exactly once before they all end up back at their starting
locations in unison, like a big, convoluted, squid conga line. This is
exactly what we find, if we raise the matrix corresponding to one of
the patterns to the 256th power, we get the identity matrix back. That
matrix is the 256th root if the Identity matrix.

The really tedious part is extracting the permutation matrices to
begin with. For that, I manually extracted the frames from the GIF
from the points in time where the squid are just arriving at their
next location, while their tentacles are still visible.  Then a very
simple vision algorithm picks out those tentacles and writes down the
permutation matrix inferred by the direction of motion.

Once we have the three matrices for the three patterns, we multiply
them together to get the overall permutation matrix for one loop of
the animation.

Then we just multiply that by itself until we find the identity matrix
again, and we finally get there after 2064 steps.

Te animation has 75 frames played with a 40ms delay, so if you want to
see the quid return home, you'll need to watch for a little over 100
minutes.

[Go code for the curious](https://github.com/cwren/squid)

I have an animation of the whole 2064 step cycle but g+ won't accept
it for some reason, so it's [here](http://imgur.com/4CzDl2j) instead.
