#!/usr/bin/python
import base64
import cv2, numpy as np

def application(environ, start_response):

    request_body=None
    retcode = '200 OK'
    resp = "dummy\r\n"
    ct  ="text/html"
    try:
       request_body_size = int(environ.get('CONTENT_LENGTH', 0))
       request_body = environ['wsgi.input'].read(request_body_size)
    except (ValueError):
       resp = "no response"
    url = environ['PATH_INFO'];
    if url == "/":
        f = open("up.html","r")
        resp = f.read()
        f.close()
    elif url == "/dn":
        ct = 'image/png'
        f = open("my.png","rb")
        resp = f.read()
        f.close()
    elif url == "/up" and request_body:
        ct = 'image/png'
        resp = request_body.replace('data:' + ct + ';base64,', "")
        data = base64.b64decode(resp)
        buf = np.frombuffer(data, dtype=np.uint8)
        img = cv2.imdecode(buf, 1)
        img = cv2.flip(img,1)
        cv2.imwrite("my.png", img)
        ok, enc = cv2.imencode(".png", img)
        resp = base64.b64encode(enc.tostring())
        resp = 'data:' + ct + ';base64,' + resp
    start_response(retcode, [('Content-Type', ct), ('Content-Length', str(len(resp)))])
    return [resp]

if __name__ == '__main__':
    print cv2.__version__
    from wsgiref.simple_server import make_server
    httpd = make_server('localhost', 9000, application)
    while True: httpd.handle_request()
