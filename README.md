# TinyScanner

TinyScanner is a simple port scanner that checks if a port is open or closed on a specified host. It supports both TCP and UDP scans and can scan a single port or a range of ports.

# Usage
You can run TinyScanner from the command line with various options:

```
$> ./tinyscanner --help
Usage: tinyscanner [OPTIONS] [HOST] [PORTS]
Options:
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.

```

# Examples

UDP Scan for a specific port:
```  
$> tinyscanner -u 127.0.0.1 80
Port 80 is open
```

TCP Scan for a specific port:
```
$> tinyscanner -t 127.0.0.1 1604
Port 1604 is closed
```

TCP Scan for a range of ports:

```
$> tinyscanner -t 10.53.224.5 80-83
Port 80 is open
Port 81 is open
Port 82 is closed
Port 83 is open
```
