# cameron_interview

![Execution](https://github.com/compassmarketing/cameron_interview/blob/master/Images/example.png)

## Thoughts, Observations, Difficulties, Etc.
---
* Go is cool. What a cool language. I'll definitely be digging into this more when I get an opportunity.
* This took me longer than an hour. Mostly due to my stubbornness. I thought I could get around the parsing `ps aux` route by using the gosigar library. I was also unfamiliar with the process of calculating cpu usage. Getting the memory usage was easy; that took maybe 10 minutes. I experimented a long time trying to get the CPU usage, and I _thought_ I had it down by taking a sample, taking another sample 500 ms later, calculating all the delta values and summing them up (User, Idle, Nice, System), then deducting the idle from this and dividing by the sum. Indeed this appeared to give me an accurate snapshot, but it wasn't updating over time for some reason. After too long I gave up and went the parsing `ps aux` route. 
* I spent too long trying to pick apart the code in the sigar library. This was certainly beneficial, but I think a better approach may have been to go through the Go tour and language specification first. Live and learn.
* Things that were initially confusing:
  1. The fact there are no objects, yet we can bind methods to receivers
  2. The unnamed types, and the := "shortcut" vs = for assignment
  3. Channels
  4. No while loops
  5. goroutines
  

Overall, I had fun with the project and it left me with a desire to pick up a book and/or go through some Go tutorials. I certainly find this area to be more stimulating than webdev, so I think I'll guide myself more in this direction.

Thanks for a fun experience.

-Cameron
