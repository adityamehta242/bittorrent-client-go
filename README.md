# BitTorrent Client

A simple BitTorrent client implementation written in Go that can download files from .torrent files using the BitTorrent protocol.

## Features

- Parse .torrent files
- Connect to BitTorrent trackers
- Handshake with peers
- Download files using the BitTorrent protocol
- Multi-peer concurrent downloading
- Piece integrity verification
- Progress tracking

## Prerequisites

- Go 1.23.4 or higher
- Internet connection
- Valid .torrent files

## Installation

1. **Clone or download the project:**
   ```bash
   git clone <your-repo-url>
   cd bittorrent-client
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Build the application:**
   ```bash
   # Windows
   go build -o bittorrent-client.exe
   
   # macOS/Linux
   go build -o bittorrent-client
   ```

## Usage

### Basic Usage

```bash
# Windows
./bittorrent-client.exe <torrent-file> <output-file>

# macOS/Linux
./bittorrent-client <torrent-file> <output-file>
```

### Parameters

- `<torrent-file>`: Path to the .torrent file you want to download
- `<output-file>`: Path where you want to save the downloaded file

### Examples

**Download Ubuntu ISO:**
```bash
# Windows
bittorrent-client.exe ubuntu-22.04.torrent ubuntu-22.04.iso

# macOS/Linux
./bittorrent-client ubuntu-22.04.torrent ubuntu-22.04.iso
```

**Download to specific directory:**
```bash
# Windows
bittorrent-client.exe movie.torrent "C:\Downloads\movie.mp4"

# macOS/Linux
./bittorrent-client movie.torrent "/home/user/Downloads/movie.mp4"
```

## How It Works

1. **Parse Torrent File**: Reads and parses the .torrent file to extract metadata
2. **Contact Tracker**: Connects to the tracker to get a list of peers
3. **Peer Handshake**: Establishes connections with available peers
4. **Download Pieces**: Downloads file pieces from multiple peers simultaneously
5. **Verify Integrity**: Checks each piece against its SHA-1 hash
6. **Assemble File**: Combines all pieces into the final file

## Project Structure

```
bittorrent-client/
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
├── bitfield/
│   └── bitfield.go         # Bitfield operations
├── client/
│   └── client.go           # Peer client implementation
├── handshake/
│   └── handshake.go        # BitTorrent handshake protocol
├── message/
│   └── message.go          # BitTorrent message types
├── peers/
│   └── peers.go            # Peer management
└── torrentfile/
    └── torrentfile.go      # Torrent file parsing and downloading
```

## Sample Output

```
Completed handshake with 192.168.1.100
Completed handshake with 10.0.0.50
Completed handshake with 172.16.0.20
(5.25%) Downloaded piece #12 from 3 peers
(10.50%) Downloaded piece #3 from 3 peers
(15.75%) Downloaded piece #7 from 3 peers
...
(100.00%) Downloaded piece #95 from 3 peers
```

## Configuration

The client uses these default settings:
- **Port**: 6881
- **Max Block Size**: 16,384 bytes (16 KB)
- **Max Backlog**: 5 concurrent requests per peer
- **Connection Timeout**: 3 seconds
- **Download Timeout**: 30 seconds per piece

## Supported Torrent Files

- Single-file torrents
- Torrents with HTTP/HTTPS trackers
- Standard BitTorrent protocol (BEP 3)

## Limitations

- Does not support multi-file torrents
- No support for DHT (Distributed Hash Table)
- No support for UDP trackers
- No seeding capability (download-only)
- No resume functionality
- No bandwidth throttling

## Troubleshooting

### Common Issues

**1. "Could not handshake with peer" errors:**
- Some peers may not be available
- Firewall blocking connections
- Try different .torrent files

**2. Slow download speeds:**
- Limited number of available peers
- Network connectivity issues
- Try popular torrents with more seeders

**3. "Failed integrity check" errors:**
- Corrupted pieces from peers
- The client will retry downloading bad pieces

**4. Build errors:**
- Ensure Go 1.23.4+ is installed
- Run `go mod tidy` to fix dependencies

### Getting Help

1. Check that your .torrent file is valid
2. Ensure you have internet connectivity
3. Try with different .torrent files
4. Check firewall settings

## Legal Notice

This BitTorrent client is for educational purposes. Users are responsible for ensuring they have the right to download and distribute any content. Always respect copyright laws and terms of service.

## License

This project is open source. Please check the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## Acknowledgments

- Built following the BitTorrent protocol specification
- Uses the bencode-go library for parsing .torrent files

---

**Note**: This is a basic implementation intended for learning purposes. For production use, consider using established BitTorrent clients.