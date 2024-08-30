# Using TinyGo on a Raspberry Pi Pico

## 1. **Download and Install TinyGo**
- **Download TinyGo:** Visit the [official TinyGo website](https://tinygo.org/getting-started/). Choose the installation method suitable for your operating system (Windows, macOS, or Linux). 
- **Install TinyGo:** Follow the installation instructions provided on the website for your platform. Typically, you can install TinyGo via package managers like `brew` for macOS, `apt` for Linux, or directly from the release binaries.

For Linux (e.g., Ubuntu/Debian):
```bash
wget https://github.com/tinygo-org/tinygo/releases/download/v0.28.1/tinygo_0.28.1_amd64.deb
sudo dpkg -i tinygo_0.28.1_amd64.deb
```

- **Set Environment Variables:** Add the TinyGo installation directory to your system's `PATH` and set the `GOROOT` environment variable. For example:
  ```bash
  export PATH=$PATH:/usr/local/tinygo/bin
  export GOROOT=/usr/local/go
  ```

## 2. **Compiling a Program**
- **Write a Program:** Create a simple `main.go` file:
  ```go
  package main

  import "machine"
  import "time"

  func main() {
      led := machine.LED
      led.Configure(machine.PinConfig{Mode: machine.PinOutput})

      for {
          led.High()
          time.Sleep(time.Millisecond * 500)
          led.Low()
          time.Sleep(time.Millisecond * 500)
      }
  }
  ```

- **Compile with TinyGo:** To compile the Go program for the Raspberry Pi Pico, run:
  ```bash
  tinygo build -o firmware.uf2 -target=pico main.go
  ```
  This command generates a `firmware.uf2` file compatible with the Pico.

## 3. **Uploading to Raspberry Pi Pico**
- **Put Pico in Bootloader Mode:** Hold down the BOOTSEL button on the Pico, then plug it into your computer. Release the button once connected. The Pico will mount as a USB storage device.
- **Upload the Firmware:** Copy the generated `firmware.uf2` file to the mounted Pico drive. Once copied, the Pico will reboot and start running the program.

## 4. **Example Programs**
- **Blink an LED:** The provided code example makes the onboard LED blink every half a second.
- **Button Control:** Extend the example by reading a button input to control the LED.
  
  ```go
  button := machine.BUTTON
  button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

  for {
      if !button.Get() {
          led.High()
      } else {
          led.Low()
      }
      time.Sleep(time.Millisecond * 10)
  }
  ```

This setup lets you quickly develop and deploy Go programs on a Raspberry Pi Pico using TinyGo, making it a powerful tool for microcontroller projects.
