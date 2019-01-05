# twitter-friends
Pull tweets just from your friends

To run this you'll need credentials from Twitter themselves, you can get this from https://developer.twitter.com. Place them in the connect command line arguments.

Even better than using the command line arguments is to put the tokens in ~/.twitter-friends.yaml, formated something like this:

```yaml
----
access-token: secret
access-token-secret: secret
consumer-key: secret
consumer-secret: secret
```