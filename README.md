# Uptime Kuma Push Agent

This repository contains scripts and instructions for setting up an Uptime Kuma Push Agent. Once you have added a push device in Uptime Kuma, you can use this repository to monitor network uptime and send results to your Uptime Kuma server. The agent is designed to run on both non-embedded systems (such as desktops and servers) and embedded systems (such as IoT devices and routers).

## Python implementation - Project Setup and Usage

Below are instructions on how to use and deploy program.py on non-embedded systems (desktops and servers).

### Prerequisites

- Python 3.x
- pip (Python package installer)

### Setting Up the Virtual Environment

1. **Create a Virtual Environment**:

    ```bash
    python3 -m venv venv
    ```

2. **Activate the Virtual Environment**:

    On macOS and Linux:

    ```bash
    source venv/bin/activate
    ```

    On Windows:

    ```bash
    .\venv\Scripts\activate
    ```

### Get Source Script

Use git clone of this repository or copy program.py to your machine (non-embedded system).

### Installing Dependencies

Install the required Python packages:

```bash
pip install requests pingparsing pyinstaller
```

Or install from requirements.txt

```bash
pip install -r requirements.txt
```

### Testing the Script

To test the script, run the following command:

```bash
python3 program.py --isp "1.1.1.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

This example will ping the ISP server at `8.8.8.8` and send the results to the specified URL every 60 seconds.

```bash
python3 program.py --isp "8.8.8.8" --base_url "https://monitoring.example.com/api/push/XXXXXX" --interval 60
```

### Building the Executable

To build the project into a standalone executable, use PyInstaller:

```bash
pyinstaller --onedir program.py
```

To build for Windows please run:

```bash
pyinstaller --onedir .\program.py --name uptime-kuma-agent_win_x64 --distpath .\release\win_x64\dist --workpath .\release\win_x64\build --specpath .\release\win_x64\
```

To build for Ubuntu please run:

```bash
pyinstaller --onedir program.py --name uptime-kuma-agent_linux_x64 --distpath release/linux_x64/dist --workpath release/linux_x64/build --specpath release/linux_x64
```

To build for Raspberry Pi please run:

```bash
pyinstaller --onedir program.py --name uptime-kuma-agent_linux_aarch64 --distpath release/linux_aarch64/dist --workpath release/linux_aarch64/build --specpath release/linux_aarch64
```

### Deployment

Schedule executable using Windows Task Manager:

....

Schedule executable using crontab:

Use crontab to schedule the executable:

```bash
crontab -e
```

Enter the following line to run the script every minute:

```bash
* * * * * /root/program --isp "1.1.1.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

Schedule executable using Task Scheduler on Synology NAS:

...

### Tested Systems

* Windows 11 (Build ..)
* Ubuntu Desktop 22.04
* Raspberry Pi 4 (version ..)
* Synology NAS (DSM7.4)



## Go implementation - Project Setup and Usage

Below are instructions on how to use and deploy program.go on non-embedded systems (desktops and servers).

### Prerequisites

- Golang installed on you development computer (https://go.dev/dl/)

### Get Source Script

Use git clone of this repository or copy program.go to your machine (non-embedded system).

### Building the Executable

To build the project into a standalone executable:

```bash
go build -o program program.go
```

### Testing the Script

To test the script, run the following command:

```bash
./program --isp "1.1.1.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

This example will ping the ISP server at `8.8.8.8` and send the results to the specified URL every 60 seconds.

```bash
./program --isp "8.8.8.8" --base_url "https://monitoring.example.com/api/push/XXXXXX" --interval 60
```

### Deployment

Schedule executable using crontab:

Use crontab to schedule the executable:

```bash
crontab -e
```

Enter the following line to run the script every minute:

```bash
* * * * * /root/program --isp "1.1.1.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

### Tested Systems

* Windows 11 (Build ..)
* Ubuntu Desktop 22.04
* Raspberry Pi 4 (version ..)
* Synology NAS (DSM7.4)


## Project Setup and Usage on embedded systems

Below are instructions on how to use and deploy program.sh on a embedded system eg routers or IoT devices.

### Prerequisites

- SSH with root access

### Get Source Script

Use git clone of this repository or copy program.sh to your machine (embedded system).

### Script Privileges

Run the following command to make the script executable privileges:

```bash
chmod +x program.sh
```

### Verify Path

Open up program.sh in your editor of choice and verify that the log file path is correct. By default its /root/uptime-kuma-agent/log.txt

### Testing the Script

To test the script, run the following command:

```bash
./program.sh --isp "100.0.0.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

### Deployment

Use crontab to schedule the executable:

```bash
crontab -e
```

Enter the following line to run the script every minute:

```bash
* * * * * /root/uptime-kuma-agent/program.sh --isp "1.1.1.1" --base_url "https://monitoring.example.com/api/push/XXXXXX"
```

### Tested Systems

* Teltonika RUT955


## Future Plans

* Be able to send get request to multiple Uptime Kuma Push Devices. Eg if this software runs on a customers network, this software can be installed on a central machine which will be able to ping PCs on the network. Then one does not have to install this software on every PC. The Uptime Kuma Server can then be located on a completely different network, preferably a VPS.
* Test and add support for more embedded systems eg Unifi USG, Advancech, Moxa, Pfsense, Codesys, Siemens WinCC, IX Developer
