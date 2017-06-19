import sys
import binascii
import argparse
import re
import time
import tabulate
import os

parser = argparse.ArgumentParser()
parser.add_argument("file", help = "access log file", nargs='?', default="access.log")
parser.add_argument("-v", "--verbose", action="store_true", help="print details")
args = parser.parse_args()


# t := fmt.Sprintf("%d", time.Now().Unix())
# tStr := fmt.Sprintf("[%s]", time.Now().Format(time.RFC3339))
# addr := r.RemoteAddr
# uri := r.RequestURI
# ua := r.UserAgent()
# referer := r.Referer()
# url := r.URL.String()

# download latest access log
cmd = "scp aws:/root/trackserver/access.log ./"
print cmd
print os.system(cmd)

pattern = re.compile("/([0-9abcdef]+)\.png")

data = []

for l in open(args.file):
    parts =  l.strip().split(',', 6)
    ts_str, epoch, src_ip, uri, url, referer, ua = parts

    m = re.match(pattern, uri)

    if m is None:
        continue
    decode = binascii.unhexlify(m.group(1))
    if ',' not in decode:
        continue

    email, sent_epoch = decode.split(',')

    data.append([
        email,
        time.ctime(float(sent_epoch)),
        time.ctime(float(epoch)),
        src_ip
    ])


headers = ["Email", "Sent At", "Open At", "Src IP"]

print tabulate.tabulate(data, headers=headers)