# pip install pyperclip
# apt-get install xclip

import pyperclip as clip
import time

while True:
    input = clip.waitForNewPaste()
    print("Before:", input)
    input = input.replace("\n", "").replace("  "," ")
    print("After :",input)
    time.sleep(0.5)
    clip.copy(input)
