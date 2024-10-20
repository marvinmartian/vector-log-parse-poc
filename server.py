import http.server
import socketserver
import json

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

if __name__ == '__main__':
    PORT = 8000
    with socketserver.TCPServer(("", PORT), MyHandler) as httpd:
        print(f"Serving on port {PORT}")
        httpd.serve_forever()  # Start the server
