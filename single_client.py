import subprocess

# Define the range for the loop
no_of_client = 1000
for i in range(1, no_of_client + 1):
    # Define your Bash script, where you want to use the value of i
    bash_script = f'''
    #!/bin/bash
    echo "Client no: {i}"
    curl http://localhost:5000
    '''

    # Execute the Bash script using subprocess
    subprocess.run(['bash', '-c', bash_script], check=True)