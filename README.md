whoami
======

Simple HTTP docker service that prints it's container ID

    $ docker build -t whoami-json .

    # ---

    $ docker run -d -p 8000:8000 --name whoami -t whoami-json
    0dd8f366534907fbc2b734bd3f2a62affe4ac945402130590160ad4aa24421e1
    
    $ curl $(hostname --all-ip-addresses | awk '{print $2}'):8000
    {"hostname":"0dd8f3665349"}
