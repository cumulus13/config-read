param(
    [string]$Version = "latest"
)

$ErrorActionPreference = "Stop"

$Repo = "cumulus13/config-read"
$InstallDir = "$env:LOCALAPPDATA\config-read"

# Colors for output
function Write-ColorOutput($ForegroundColor) {
    $fc = $host.UI.RawUI.ForegroundColor
    $host.UI.RawUI.ForegroundColor = $ForegroundColor
    if ($args) {
        Write-Output $args
    }
    $host.UI.RawUI.ForegroundColor = $fc
}

# Detect architecture
function Get-Architecture {
    $arch = $env:PROCESSOR_ARCHITECTURE
    switch ($arch) {
        "AMD64" { return "amd64" }
        "ARM64" { return "arm64" }
        default {
            Write-ColorOutput Red "Unsupported architecture: $arch"
            exit 1
        }
    }
}

# Download and install
function Install-ConfigRead {
    $arch = Get-Architecture
    
    if ($Version -eq "latest") {
        $url = "https://github.com/$Repo/releases/latest/download/config-read_windows_${arch}.zip"
    } else {
        $url = "https://github.com/$Repo/releases/download/$Version/config-read_windows_${arch}.zip"
    }
    
    Write-ColorOutput Yellow "Downloading config-read $Version..."
    
    # Create temp directory
    $tempDir = Join-Path $env:TEMP "config-read-install"
    New-Item -ItemType Directory -Force -Path $tempDir | Out-Null
    
    $zipFile = Join-Path $tempDir "config-read.zip"
    
    # Download
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
    Invoke-WebRequest -Uri $url -OutFile $zipFile
    
    # Extract
    Expand-Archive -Path $zipFile -DestinationPath $tempDir -Force
    
    # Install
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    $exePath = Join-Path $InstallDir "config-read.exe"
    Move-Item -Path (Join-Path $tempDir "config-read.exe") -Destination $exePath -Force
    
    # Add to PATH
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($userPath -notlike "*$InstallDir*") {
        [Environment]::SetEnvironmentVariable("Path", "$userPath;$InstallDir", "User")
        $env:Path = "$env:Path;$InstallDir"
    }
    
    # Cleanup
    Remove-Item -Recurse -Force $tempDir
    
    Write-ColorOutput Green "✓ config-read installed successfully!"
    
    # Verify
    if (Get-Command config-read -ErrorAction SilentlyContinue) {
        config-read version
    }
}

# Run installation
Install-ConfigRead
