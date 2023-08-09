# scanner-proxy
This is the second chaper code.
Now it's only working on 'tcp', in future will imporove and become a package.
## book suggestions:
"
You could make a couple of improvements to this program. First, you’re 
sending on the results channel for every port scanned, and this isn’t necessary. The alternative requires code that is slightly more complex as it uses an 
additional channel not only to track the workers, but also to prevent a race 
condition by ensuring the completion of all gathered results. As this is an 
introductory chapter, we purposefully left this out; but don’t worry! We’ll 
introduce this pattern in Chapter 3. Second, you might want your scanner to 
be able to parse port-strings—for example, 80,443,8080,21-25, like those that 
can be passed to Nmap. If you want to see an implementation of this, see 
https://github.com/blackhat-go/bhg/blob/master/ch-2/scanner-port-format/. We’ll 
leave this as an exercise for you to explore.
"