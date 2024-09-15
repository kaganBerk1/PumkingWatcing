# PumkingWatcing

PumkingWatcing is a Windows system tray application that provides real-time system usage information. It monitors and displays CPU, RAM, disk, and network usage statistics, along with system uptime. The application also plays a sound notification when it exits.

## Features

- Displays CPU, RAM, disk, and network usage in the system tray.
- Shows system uptime.
- Supports localization: displays information in English or Turkish based on the system language.
- Plays a sound notification upon exiting.

## Requirements

- Go 1.18 or later
- Windows operating system
- Go packages: `faiface/beep`, `getlantern/systray`, `shirou/gopsutil`

## Installation

1. **Clone the repository:**

    ```
    git clone https://github.com/kaganBerk1/PumkingWatcing.git
    cd PumkingWatcing
    ```

2. **Download the required Go modules:**

    ```
    go mod download
    ```

3. **Build the application:**

    ```b
    go build -o PumkingWatcing.exe
    ```

## Usage

1. **Run the application:**

    Double-click `PumkingWatcing.exe` or run it from the command line:

    ```
    ./PumkingWatcing.exe
    ```

2. **System Tray Interaction:**

    The application will appear in the system tray (bottom-right corner of the screen). You can view real-time system usage statistics by hovering over or clicking the tray icon.

3. **Exit the Application:**

    To close the application, right-click the system tray icon and select "Exit."

## Customization

- **Icon:** Replace `pumpkin.ico` with your own icon file to customize the system tray icon.
- **Exit Sound:** Update the `song.wav` file to change the sound played when the application exits.

## Troubleshooting

- **Application Does Not Appear in System Tray:**
  - Ensure that the application is running and not minimized to the taskbar.
  - Verify that the `pumpkin.ico` file is present and correctly referenced in the code.

- **No Sound on Exit:**
  - Check that `song.wav` exists in the same directory as the executable.
  - Make sure your system's sound settings are correctly configured.

## Contributing

We welcome contributions to improve this project. If you'd like to contribute, please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please refer to the [CONTACT.md](CONTACT.md) file.

## Acknowledgements

- Special thanks to the contributors of the Go packages used in this project: `faiface/beep`, `getlantern/systray`, and `shirou/gopsutil`.
