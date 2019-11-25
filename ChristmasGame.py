import sys
import base64
import random
from random import shuffle

#Encoding
def encode(key, clear):
    enc = []
    for i in range(len(clear)):
        key_c = key[i % len(key)]
        enc_c = chr((ord(clear[i]) + ord(key_c)) % 256)
        enc.append(enc_c)
    return base64.urlsafe_b64encode("".join(enc))

# Decoding
def decode(key, enc):
    dec = []
    enc = base64.urlsafe_b64decode(enc)
    for i in range(len(enc)):
        key_c = key[i % len(key)]
        dec_c = chr((256 + ord(enc[i]) - ord(key_c)) % 256)
        dec.append(dec_c)
    return "".join(dec)

rsg = ["David", "Mike", "Edvin", "Alexander", "Jens", "Per", "Max", "Jacob"]
pool = ["David", "Mike", "Edvin", "Alexander", "Jens", "Per", "Max", "Jacob"]

shuffle(rsg)
shuffle(pool)

for x in range(1,9):
	print ("TURN " + str(x))
	buyer = random.choice(rsg) # pick a random buyer
	reciever = random.choice(pool) # pick a random reciever
	while buyer == reciever: # if buyer == reciever, get a new reciever
		reciever = random.choice(pool)
	encrypt = encode("r$gchristmasgame",reciever)
	print (buyer + " buys for " + encrypt)
	rsg.pop(rsg.index(buyer)) # remove the buyer
	pool.pop(pool.index(reciever)) # remove the reciever
	print ("Remaining people in RSG")
	print (rsg) # print rsg-list
