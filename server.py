import http.server
import socketserver
import json
import signal
import sys

class MyHandler(http.server.SimpleHTTPRequestHandler):
    def do_POST(self):
        content_length = int(self.headers['Content-Length'])  # Get the length of the data
        post_data = self.rfile.read(content_length)  # Read the data
        try:
            # Attempt to decode the JSON data
            data = json.loads(post_data)
            print("Received POST data:")
            print(json.dumps(data, indent=2))  # Pretty-print the JSON data
        except json.JSONDecodeError:
            print("Received non-JSON data:")
            print(post_data.decode('utf-8'))  # Print the raw data if not JSON

        self.send_response(204)  # No Content response
        self.end_headers()  # End headers

def signal_handler(signal_received, frame, httpd):
    print('SIGTERM received, shutting down gracefully...')
    httpd.shutdown()
    sys.exit(0)

if __name__ == '__main__':
    PORT = 8000
    with socketserver.TCPServer(("", PORT), MyHandler) as httpd:
        print(f"Serving on port {PORT}")

        # Register the signal handler for SIGTERM
        signal.signal(signal.SIGTERM, lambda signal_received, frame: signal_handler(signal_received, frame, httpd))

        # Start the server
        httpd.serve_forever()
