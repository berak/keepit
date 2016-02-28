"""
sdf.lonestar.org
gopher.quux.org
gopher.floodgap.com
gopher.metafilter.com
gopherspace.de
redhill.net.nz
http://gopher.floodgap.com/gopher/gw?gopher://gopher.floodgap.com:70/1/world
"""

import sys, socket

def readline(sock):
    s = ""
    while sock:
        c = sock.recv(1)
        if not c:      break
        elif c==b'\n': break
        elif c==b'\r': continue
        else:          s += c
    return s
    
ser = sys.argv[1]
pretty = False;
if len(sys.argv)>2: pretty = sys.argv[2]
while True:
   sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
   sock.connect((ser, 70))
   s = sys.stdin.readline().strip()
   if s == ".": break
   sock.send(s + "\r\n")
   while sock:
      s = readline(sock)
      if len(s)==0 : break
      if pretty and s.find("\t")>=0: s = " ".join(s[1:].split("\t")[:-2]).replace("fake","").replace("TITLE","")
      print s
   sock.close()
