# Packet Disorder Proxy Server

This is made specifically for distrupting the order of packets sent as rows of an image.
Packet disordering happens on real networks too and this was made for an exercise for the DSA students I am TA'ing, to use the heap data structure to efficiently recover the original packet ordering on the server that they wrote.

## Client-side Assumptions:
- The server is running on `localhost` with port `9989` and used the `tcp` protocol.
- This server assumes that it receives image dimensions as two 4-byte little-endian integers in the first two requests by the client.
- It also assumes that each packet (or image row) has the first 4-bytes in it, dedicated to the packet id.
