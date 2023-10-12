# Packet Disorder Proxy

This is made specifically for distrupting the order of packets sent as rows of an image.
Packet disordering happens in real networks too and this was made for an exercise for the DSA students I am TA'ing, to use the heap data structure to efficiently recover the original packet ordering on the server that they wrote.

## Server-side Assumptions:
- The server is running on `localhost` with port `9989` and is using the `tcp` protocol.

## Client-side Assumptions:
- This proxy receives image dimensions as two 4-byte little-endian integers in the first two messages by the client.
- The client processes the acknowledgment byte from the proxy and only after receiving it, sends another packet.
- Each packet (or image row) has the first 4-bytes in it, dedicated to the packet id.
- Each packet is smaller than 8kB.
