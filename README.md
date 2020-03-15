# Installer CLI
This is an installer CLI for the https://github.com/smartassets/backend-server.

# Installation steps:
1. Download CLI for your operation system:

Mac OS X 64 bit | Windows 32 bit | Windows 64 bit | Linux 32 bit | Linux 64 bit
--- | --- | --- | --- | ---
[assets.osx](https://github.com/smartassets/installer/releases/download/1.0.1/assets.osx) | [assets.win32](https://github.com/smartassets/installer/releases/download/1.0.1/assets.win32) | [assets.win64](https://github.com/smartassets/installer/releases/download/1.0.1/assets.win64) | [assets.linux32](https://github.com/smartassets/installer/releases/download/1.0.1/assets.linux32) | [assets.linux64](https://github.com/smartassets/installer/releases/download/1.0.1/assets.linux64)

2. Create a folder and add the downloaded content there
- For Mac/Unix:
```
mkdir -p <prefered-directory-name>;
mv <path-to-downloaded-content> <prefered-directory-name>/assets
```

3. Add the newly created directory to PATH variable
- For Mac/Unix - edit the ${HOME}/.bash_profile file with the following content:
```
export PATH=$PATH:<full-path-to-prefered-directory-name>
```
