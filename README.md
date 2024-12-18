
# Persian Subtitle Fixer

This is a simple Go (Golang) script designed to fix issues with Unicode subtitles caused by incorrect encoding. It processes `.srt` subtitle files in the `subs` directory, converting their encoding from Windows-1256 to UTF-8.

## Problem Explanation
Subtitle files, especially those in certain languages, are often incorrectly encoded. This can result in garbled or unreadable text when viewed in media players or editors that expect proper UTF-8 encoding. This script solves the problem by detecting and converting the encoding to UTF-8, ensuring the subtitles display correctly.

## How to Use
1. **Prepare Your Environment**:
    - Ensure you have Go installed on your system.
    - Place all the `.srt` subtitle files you want to fix into a folder named `subs` (in the same directory as the script).

2. **Run the Script**:
   ```bash
   go run main.go
   ```

3. **Check the Results**:
    - The script will overwrite the original `.srt` files with their UTF-8 encoded versions.

## Example
### Input (Before Running the Script):
A subtitle file in `subs/example.srt` encoded in Windows-1256:

```
1
00:00:01,000 --> 00:00:03,000
\xE3\xC7\xD1\xD3
```

### Output (After Running the Script):
The same file now encoded in UTF-8:

```
1
00:00:01,000 --> 00:00:03,000
مارس
```

## Contributing
Contributions are welcome! If you find any bugs or have ideas for improvements:
1. Fork the repository.
2. Make your changes.
3. Open a pull request describing your updates.

## Notes
- The script processes all `.srt` files in the `subs` directory. Ensure you back up your files before running it.
- It uses the [golang.org/x/text/encoding](https://pkg.go.dev/golang.org/x/text/encoding) library for encoding conversions.
