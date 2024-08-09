# Torrent-Client

A BitTorrent client implemented in Golang. This project demonstrates a basic torrent downloader with core BitTorrent protocol functionalities, including tracker communication, peer-to-peer data transfer, and piece management.

## Project Structure

Torrent-Client  
│    
├─ bitfield  
│   ` ` ` ` ` `└─ `bitfield.go`  
├─ client  
│   ` ` ` ` ` `└─ `client.go`  
├─ handshake  
│  ` ` ` ` ` ` └─ `handshake.go`    
├─ message  
│   ` ` ` ` ` `└─ `message.go`  
├─ p2p  
│   ` ` ` ` ` `└─ `p2p.go`  
├─ peers  
│   ` ` ` ` ` `└─ `peers.go`  
├─ torrentFile  
│  ` ` ` ` ` ` ├─ `torrentFile.go`  
│  ` ` ` ` ` ` └─ `tracker.go`  
└─ `main.go`

## Installation

#### (Method-1)
1. Clone the repository:
```bash
git clone https://github.com/prabal199251/Torrent-Client.git
```

2. Navigate to the project directory:
```bash
cd torrent-client
```

3. Build the project:
```bash
go build -o torrent-client main.go
```

#### (Method-2)

1. Install the package using `go get`:
   ```bash
   go get github.com/prabal199251/Torrent-Client
   ```

2. Import the package in your Go code:
    ```bash
    import "github.com/prabal199251/Torrent-Client"
    ```

## Usage

Run the application with the path to your `.torrent` file and a destination directory:

```bash
./torrent-client path/to/your.torrent /path/to/destination
```

Try downloading [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist)

```bash
go run main.go debian-12.6.0-amd64-netinst.iso.torrent debian.iso
```

## Features
* Download torrent files from trackers.
* Connect to peers and exchange torrent pieces.
* Manage and verify downloaded pieces.


## Limitations
* Only supports `.torrent` files (no magnet links)
* Only supports HTTP trackers
* Does not support multi-file torrents
* Strictly leeches (does not support uploading pieces)

