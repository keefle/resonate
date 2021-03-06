# resonate
A bidirectional folder synchronisation daemon.

### Installation
```
$ go get -u git.nightcrickets.space/keefleoflimon/resonate
```

### How to run
```
$ mkdir test1
$ resonate -dir test1 -port 1234 -peer localhost:4321
```

In another terminal:
```
$ mkdir test2
$ resonate -dir test2 -port 4321 -peer localhost:1234
```

Now if you list files you will see two new folders:
```
$ ls
...
test1-resonate
test2-resonate
...
```

Make any changes to any of the two new folders (test2-resonate, test1-resonate)
create files in one of them, write to them etc.  Those changes will be replicated
bit by bit on the other folder.

### Support
Currently only works on Linux distributions and FreeBSD.

### Limitations
- Currently both devices that need to be in sync need to be on the same network
- Only two devices can sync currently
- No Windows support


## IDEAS:
- Binary Change log (mongodb)
- Transactional file systems
- Cap theorem
- Only Big group writes in case of partition "corum? (strict majority)
floor(N/2)" Downside? if the system is cut into three parts,
you may not be aable to know if you hold the majority
- Virtual Filesystem FUSE

---

Note: The text below was taken from another similar project (this was originally
an Interview project). I've taken the text of the problem for inspiration.

# Index

- [Problem Statement](#problem-statement)
- [Challenges](#challenges)
- [Original Problem](#original-problem)

## Problem Statement

**We need to create a solution to synchronize two folders in two different devices, each folder have a file called `data`, which needs to be synchronized and it should be bi-directional**. To finally decide upon our solution, we need to make sure the following characteristics are met.

- **Reliability:** Solution should be reliable, and should take care of all the possible cases it can go wrong.

- **Automatic:** Everything should happen in the background with minimum human intervene.

- **Cheap:** It should do everything in minimum time while using very few resources, with least latency.

> Find the original problem statement at the end of the file.

## Challenges

Following are the challenges we need to tackle, mentioned with their possible solutions.

**Challenge: How should we communicate between two computers?**

When communicating over the internet, following possible scenarios:

No |IP-1 | IP-2 |  Details |
---|-----|------|----------|
 1 | Static | Static |`Server to Server` <br> Example: Backup servers connected to production servers.
 2 | Dynamic | Static | `Server to remote devices` <br> Example: Dropbox, Google drive connected to remote devices.
 3 | Dynamic | Dynamic | `Two remote devices` <br> Example: Two remote devices connected.

For our solution, we will only cover the 1st case. The same solution can be applied for other cases, which little changes. We have multiple methods to communicate over the internet, some are mentioned below.

No | Method | Secure | Reliable | Speed | Automatic |
---|--------|--------|----------|-------|-----------|
1 | TCP (Transmission Control Protocol) | Medium | High | Fast | Yes
2 | SSH (Secure Shell) | High | High | Medium | Not by default
3 | UDP (User Datagram Protocol) | Less | Less | Fastest | Yes

*Result*

After comparing all the methods, I have decided to use **`SSH`**.

***

**Challenge: When files should be synchronized?**

Following are the options on when files should be synchronized:

1. Immediately `inotify-tools`
2. After a time gap `crontab`

*Result*

After comparing all the options, I have decided to synchronize **`immediately`** a change is detected. *It will make it real time, will prevent merge conflict.*

> **Note:** if we update the device 1, updates will be sent to device 2, that will trigger the script, which will try to update the device 1, and this might go in a loop. We need to handle the **butterfly effect**.

***

**Challenge: How should we measure the difference?**

To measure the difference, we can use any of method mentioned below.

No | Method | Reliable | Speed
---|--------|----------|------
1 | Time Modification (Metadata) | `Less Reliable` <br>software can and does manipulate the modification time. Also, the user might change system time and confuse the sync program. | `Fast` <br> O(1)
2 | Checksum (Hash the file) | `More Reliable` <br> It's an (almost) certain way measure difference, hash collisions do happen, but It is rare. | `Slow` <br> O(n)

*Result*

After comparing all the methods, I have decided to use **`Checksum`**.

***

**Challenge: How should we tackle the differences?**

Following are the 3 cases that we need to handle:

- The file exists on device 1, not on device 2
- The file exists on both devices and is identical
- The file exists on both devices and is different

Following is the action table:

File 1 | File 2 | Action
-------|--------|-------
Deleted | No Deleted | Delete
Deleted | Deleted | Nothing
No Change | No change | Nothing
Modification | No change | Use A
Modification | Modification | Merge

> **Note:** vice-versa is also true in this action table.

Following is the action table based on the time:

Time x | Time x+1 | Action
-------|----------|-------
Does not exist | Exist | Created
Existed | Does not exist |  Deleted
Exist | Modification | Modification

*Result*

For this solution, we are just tracking one file `data`, and we are doing synchronization immediately, so most of the cases won't apply to us.

***

**Challenge: How will you handle a merge conflict?**

Following are the methods we have to prevent merge conflict:

No | Method | Details | Merge Quality | Automatic
---|--------|---------|---------------|-----------
1 | Ask the user | Ask the user how to merge them or which one to pick. | Best| No
2 | Lock other user files | Lock a file if it is owned by the other user. | No-Merge | Yes
3 | overwrite with latest changes | We can overwrite the file with the latest changes. | Medium | Yes

> **Note:** Resolving to merge conflict is technically impossible without a human to intervene

*Result*

Again, in our case, this won't be a constant problem, we have just one file. So I have decided to use `overwrite with latest changes` method.

***

## Final Solution

## Original Problem

```
Clone Wars 2.0
--------------

Prime Minister Lama Su,

I hope this letter finds you in the best of health.

The last batch of clones you built for us was faulty
and did not perform as expected (https://www.youtube.com/watch?v=b0DuUnhGBK4)

We unearthed some secrets about how the droid army was trained and hope that
you can use this information to make a better army this time around. With the
galaxy on the brink of another war, I cannot help but emphasize how much a
large discount will help the Republic in its efforts.

One of our allies came across these schematics in an abandoned base that shed some
light on the droid training exercises, master Yoda concluded that a pair of droids
undergo various kinds of battle simulations during which each droid records its
progress and learning in a force, currently unfamiliar to us, called "Data".
This force from both droids is then combined in a ritual called the
"Sync" resulting in both droids having an increased data force.

Please have a look at this schematic, your engineers may have better luck
decoding its mysteries.

            +----------------+                +----------------+
            |                |                |                |
            |   +--------+   |      Sync      |   +--------+   |
            |   |-|Data|-|   | +------------> |   |-|Data|-|   |
            |   +--------+   | <------------+ |   +--------+   |
            |                |                |                |
            |    Driod  A    |                |    Driod  B    |
            |                |                |                |
            +----------------+                +----------------+

May the force be with you.

- Sifo-Dyas


[....2 months later....]


Prime Minister Lama Su!,

I hope the army is coming along nicely. The force has given us more clarity in
the last few months. As it turns out, this "Data" that we were so worried about,
is just a method by which the droids store information about their experiences and
orders. Most importantly, the "Sync" ritual was just an exchange of files
from one droid to another in both directions. This is how their data force
increased after the ritual.

Master Windoo has been doing extensive research and has come up with a simplified
experiment to test if this training method can be implemented. He says that you
should start by figuring out how to synchronize data between a folder on one
device (say device A) and a folder on another device (say device B).
In addition to that, a change made to the data on one device should also be made
available to the other device as well. If we have a way to do this then we could
potentially improve the quality of the new clone army. I hope your engineers
are able to make sense of all of this information. Do write back to me if you
need more information.

Please share your method and implementation in great detail with us so
that it can be added to our records in the Jedi Temple. I wish you luck.

May the force be with you.

- Sifo-Dyas


                                 +---------------------+
                                 | What's going on here?|
                                 +------------------+--+
                                                    |
                                                    |
  _                                                 |
  \\                                                |
   \\_          _.-._                               |
    X:\        (_/ \_)     <------------------------+
    \::\       ( ==  )
     \::\       \== /
    /X:::\   .-./`-'\.--.
    \\/\::\ / /     (    l
     ~\ \::\ /      `.   L.
       \/:::|         `.'  `
       /:/\:|          (    `.
       \/`-'`.          >    )
              \       //  .-'
               |     /(  .'
               `-..-'_ \  \
               __||/_ \ `-'
              / _ \ #  |
             |  #  |#  |   B-SD3 Security Droid
          LS |  #  |#  |      - Front View -

(http://www.ascii-art.de/ascii/s/starwars.txt)
```

## Step 0:

## Step 1:
## Step 2:
## Step 3:
## Step 4:


## Goal
Live Sync of two folders on two different machines

Q: How do you sync?
A: Each server keeps track of the files it has (tree)
A: On change, the server sends the files which changed
A: The second server overwrites the old changed files
A: with the new ones


