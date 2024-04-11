import socket
import threading

SERVER_ADDRESS = ('localhost', 5000)
NUM_CLIENTS = 10000

def client_request():
    try:
        client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client_socket.connect(SERVER_ADDRESS)
        # You can customize the request message as per your requirement
        request_message = "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"
        client_socket.sendall(request_message.encode())
        response = client_socket.recv(4096)
        print(f"Received response: {response.decode()}")
    except Exception as e:
        print(f"Error: {e}")
    finally:
        client_socket.close()

def main():
    threads = []
    for i in range(NUM_CLIENTS):
        print(f"Client no: {i}")
        thread = threading.Thread(target=client_request)
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()

if __name__ == "__main__":
    main()