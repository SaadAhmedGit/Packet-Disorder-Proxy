# Packet Disorder Proxy Server

This is made specifically for distrupting the order of packets sent as rows of an image.
This server assumes it receives image dimensions as two 4-byte little-endian integers in the first two requests by the client.
It also assumes that each packet (or image row) has the first 4-bytes in it, dedicated to the packet id.

Packet disordering happens on real networks too and this was made as an exercise for DSA students I am TA'ing so they can use the heap data structure to efficiently recover the original packet ordering on the server that they wrote.
