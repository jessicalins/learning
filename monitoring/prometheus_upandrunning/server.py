import random
import http.server
from prometheus_client import start_http_server, Counter

REQUESTS = Counter('hello_words_total', '# of hello Words requested.')
EXCEPTIONS = Counter('hello_world_exceptions_total', '# exceptions serving Hello World.')

class MyHandler(http.server.BaseHTTPRequestHandler):
    @EXCEPTIONS.count_exceptions()
    def do_GET(self):
        REQUESTS.inc()
        if random.random() < 0.2:
            raise Exception
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b"Hello World")
        
if __name__ == "__main__":
    start_http_server(8000)
    server = http.server.HTTPServer(('localhost', 8001), MyHandler)
    server.serve_forever()
        